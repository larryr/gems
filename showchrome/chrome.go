package showchrome

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
)

type ChromeCtl struct {
}

func (ch *ChromeCtl) Launch(args []string) string {

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), chromedp.DefaultExecAllocatorOptions[:]...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// ensure that the browser process is started
	if err := chromedp.Run(taskCtx); err != nil {
		log.Fatal(err)
	}

	return "ok"
}
