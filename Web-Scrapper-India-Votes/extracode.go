package webscrapperindiavotes

/*
func main() {
	filepath := "C:\\Users\\User\\Desktop\\GoLang\\Web-Scrapper-India-Votes\\outer.html"

	fileContent, err := os.ReadFile(filepath)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBufferString(string(fileContent)))

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	doc.Find("td.numberTable").Each(func(i int, s *goquery.Selection) {
		fmt.Println("Content ", i, ": ", s.Text())
	})

}
*/
/*
func main() {
	filePath := "C:\\Users\\User\\Desktop\\GoLang\\Web-Scrapper-India-Votes\\outer.html" // Replace with your file path

	htmlContent, err := getDesiredLine(filePath)
	if err != nil {
		log.Fatal(err)
	}

	modifiedHTML := replaceGtWithNewline(htmlContent)

	fileName := "C:\\Users\\User\\Desktop\\GoLang\\Web-Scrapper-India-Votes\\output1.html"
	if err := os.WriteFile(fileName, []byte(modifiedHTML), 0644); err != nil {
		log.Fatal(err)
	}
}

func replaceGtWithNewline(content string) string {

	modifiedContent := strings.ReplaceAll(content, ">", ">\n")
	return modifiedContent
}

func getDesiredLine(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Increase the buffer size for the scanner (1 << 20 is an example size, adjust as needed)
	const maxCapacity = 1 << 20
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	var lastLine string
	var i int = 0
	for scanner.Scan() {
		i++
		if i == 1838 {
			lastLine = scanner.Text()
			if err := scanner.Err(); err != nil {
				return "", err
			}

			return lastLine, nil

		}
	}
	return lastLine, nil
}
*/
/*
func main() {
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
		chromedp.Navigate("https://www.indiavotes.com/pc/acwisedetails/29/7873/16"),
		chromedp.Sleep(time.Second*10),
		chromedp.Evaluate(`document.documentElement.outerHTML`, &htmlContent),
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(htmlContent)

	fileName := "C:\\Users\\User\\Desktop\\GoLang\\Web-Scrapper-India-Votes\\outputnew.html"
	if err := os.WriteFile(fileName, []byte(htmlContent), 0644); err != nil {
		log.Fatal(err)
	}
}
*/
/*
func main() {
	filePath := "C:\\Users\\User\\Desktop\\GoLang\\Web-Scrapper-India-Votes\\output.html" // Replace with your file path

	// Read the contents of the file
	htmlContent, err := readFileToString(filePath)
	if err != nil {
		log.Fatal(err)
	}

	cleanedHTML := cleanHTMLContent(htmlContent)

	fileName := "C:\\Users\\User\\Desktop\\GoLang\\Web-Scrapper-India-Votes\\output1.html"
	if err := os.WriteFile(fileName, []byte(cleanedHTML), 0644); err != nil {
		log.Fatal(err)
	}

}

func readFileToString(filePath string) (string, error) {
	// Read the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// Convert byte slice to string
	fileContent := string(content)

	return fileContent, nil
}

func cleanHTMLContent(content string) string {
	// Remove tabs
	content = strings.ReplaceAll(content, "\t", "")

	// Remove newline characters
	content = strings.ReplaceAll(content, "\n", "")

	// Replace multiple spaces with a single space
	content = strings.ReplaceAll(content, "  ", " ")

	return content
}
*/
