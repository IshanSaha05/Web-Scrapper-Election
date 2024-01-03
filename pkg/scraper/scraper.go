package scraper

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/IshanSaha05/IndiaVotes/pkg/config"
	"github.com/IshanSaha05/IndiaVotes/pkg/models"
	"github.com/IshanSaha05/IndiaVotes/pkg/mongodb"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func isNumeric(s string) bool {
	// Regular expression to match only digits
	numericRegex := regexp.MustCompile(`^\d+$`)
	return numericRegex.MatchString(s)
}

func getHTMLContent(url string) (string, error) {
	var cancelFuncs []context.CancelFunc

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), chromedp.Flag("headless", true), chromedp.WindowSize(1920, 1080))
	cancelFuncs = append(cancelFuncs, cancel)

	ctx, cancel = context.WithTimeout(ctx, time.Second*50)
	cancelFuncs = append(cancelFuncs, cancel)

	ctx, cancel = chromedp.NewContext(ctx)
	cancelFuncs = append(cancelFuncs, cancel)

	cancelAll := func() {
		for _, cancel := range cancelFuncs {
			cancel()
		}
	}
	defer cancelAll()

	var htmlContent string

	err := chromedp.Run(
		ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(time.Second*10),
		chromedp.Evaluate(`document.documentElement.outerHTML`, &htmlContent),
	)

	if err != nil {
		log.Fatal(err)
	}

	return htmlContent, nil
}

func getACData(htmlContent string) ([]models.ACData, error) {
	fmt.Println("Message: Getting Data.")
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))

	if err != nil {
		return nil, err
	}

	var datas []models.ACData

	// Taking out the common ac data which will be present in every document.
	var commonData models.CommonACData
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		text := s.Text()

		// Getting the data of State, Year, Parliament Constituency Name, Winner of that PC, Winning party, Total Assembly Consitiuency present.
		if strings.Contains(text, "State:") {
			commonData.State = strings.TrimSpace(strings.TrimPrefix(text, "State:"))
		} else if strings.Contains(text, "Year:") {
			number := strings.TrimSpace(strings.TrimPrefix(text, "Year:"))

			if isNumeric(number) {
				commonData.Year, err = strconv.Atoi(number)
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
			}
		} else if strings.Contains(text, "PC Name:") {
			commonData.ParliamentConstituency = strings.TrimSpace(strings.TrimPrefix(text, "PC Name:"))
		} else if strings.Contains(text, "Winner:") {
			commonData.WinningCandidate = strings.TrimSpace(strings.TrimPrefix(text, "Winner:"))
		} else if strings.Contains(text, "Winning Party:") {
			commonData.WinningCandidateParty = strings.TrimSpace(strings.TrimPrefix(text, "Winning Party:"))
		} else if strings.Contains(text, "Total ACs:") {
			number := strings.TrimSpace(strings.TrimPrefix(text, "Total ACs:"))

			if isNumeric(number) {
				commonData.TotalACs, err = strconv.Atoi(number)
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
			}
		}
	})

	// Taking the main data.
	mainDiv := doc.Find("div#m1.mapTabData")

	mainDiv.Contents().Each(func(i int, s *goquery.Selection) {
		// Going through each table.
		if s.Is("table") {
			// Storing the common data for each AC.
			var commontabledata models.CommonACTableData

			// Extract common AC Name and Total Votes for a particular table.
			s.Find("tr").Each(func(i int, tr *goquery.Selection) {
				tr.Find("td").Each(func(i int, td *goquery.Selection) {
					tdText := td.Text()

					if strings.Contains(tdText, "AC Name:") {
						commontabledata.AssemblyConstituency = strings.TrimSpace(strings.TrimPrefix(tdText, "AC Name:"))
					} else if strings.Contains(tdText, "Total AC Votes:") {
						number := strings.ReplaceAll(strings.TrimSpace(strings.TrimPrefix(tdText, "Total AC Votes:")), ",", "")

						if isNumeric(number) {
							commontabledata.TotalVotes, err = strconv.Atoi(number)
							if err != nil {
								log.Fatal(err)
								os.Exit(1)
							}
						}
					}
				})
			})

			// Now find the div next to table to get the votes data of each candidate.
			nextDiv := s.Next()
			if nextDiv.Is("div") {
				// Extract the data under this div.
				tbdoyNode := nextDiv.Find("table").Find("tbody")

				tbdoyNode.Find("tr").Each(func(i int, tr *goquery.Selection) {
					var differentACData models.DifferentACData

					differentACData.Candidate = tr.Find("td").Eq(0).Text()
					differentACData.CandidateParty = tr.Find("td").Eq(1).Text()
					number := strings.ReplaceAll(tr.Find("td").Eq(2).Text(), ",", "")

					if isNumeric(number) {
						differentACData.Votes, err = strconv.Atoi(number)
						if err != nil {
							log.Fatal(err)
							os.Exit(1)
						}
					}

					differentACData.VotesPercentage, err = strconv.ParseFloat(strings.TrimSuffix(tr.Find("td").Eq(3).Text(), "%"), 64)
					if err != nil {
						log.Fatal(err)
						os.Exit(1)
					}
					differentACData.District = ""
					differentACData.Type = "AC"

					// Compiling the data into one data to append in the database.
					var data models.ACData

					// Inserting the Common AC Data.
					data.CACData = commonData
					data.CACTableData = commontabledata
					data.DACData = differentACData

					// Appending data inside the main returning slice of data.
					datas = append(datas, data)
				})
			}
		}
	})

	return datas, nil
}

func ScrapeStoreACData(mongoDBObject *mongodb.MongoDB) {
	// Get the HTML content of the main page.
	htmlContent, err := getHTMLContent(config.ScrapeSiteUrl)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Trying to iterate through each a tag.
	doc.Find("a.detailResults").Each(func(i int, s *goquery.Selection) {
		// Checking whether it's a AC View link.
		if strings.Contains(s.Text(), "AC View") {
			link, exists := s.Attr("href")

			// If a link is there, then getting AC wise data for the PC and storing it.
			if exists {
				acHtmlContent, err := getHTMLContent(link)
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}

				// Getting the data.
				datas, err := getACData(acHtmlContent)
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}

				// Storing the data in the db.
				mongoDBObject.SetMongoDatabase("Election-Database")
				mongoDBObject.SetMongoCollection("Election-Collection")

				err = mongoDBObject.InsertIntoDB(datas)
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
			}
		}
	})
}
