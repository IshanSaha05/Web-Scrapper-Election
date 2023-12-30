package scraper

import (
	"context"
	"time"

	"github.com/chromedp/chromedp"
)

func GetHTMLContent(url string, timeout int, sleep int) (string, error) {
	var cancelFuncs []context.CancelFunc

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), chromedp.Flag("headless", true), chromedp.WindowSize(1920, 1080))
	cancelFuncs = append(cancelFuncs, cancel)

	ctx, cancel = context.WithTimeout(ctx, time.Second*time.Duration(timeout))
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
		chromedp.Sleep(time.Duration(sleep)),
		chromedp.Evaluate(`document.documentElement.outerHTML`, &htmlContent),
	)

	if err != nil {
		return "", err
	}

	return htmlContent, nil
}
