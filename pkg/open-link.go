package pkg

import (
	"context"

	"github.com/chromedp/chromedp"
)

func OpenLinkInChrome(urlToOpen string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(urlToOpen),
	)
	if err != nil {
		return err
	}

	return nil
}
