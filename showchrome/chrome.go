package showchrome

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"net/url"
)

type ChromeCtl struct {
}

var myExecAllocatorOptions = [...]chromedp.ExecAllocatorOption{
	chromedp.NoFirstRun,
	chromedp.NoDefaultBrowserCheck,

	// After Puppeteer's default behavior.
	chromedp.Flag("disable-background-networking", true),
	chromedp.Flag("enable-features", "NetworkService,NetworkServiceInProcess"),
	chromedp.Flag("disable-background-timer-throttling", true),
	chromedp.Flag("disable-backgrounding-occluded-windows", true),
	chromedp.Flag("disable-breakpad", true),
	chromedp.Flag("disable-client-side-phishing-detection", true),
	chromedp.Flag("disable-default-apps", true),
	chromedp.Flag("disable-dev-shm-usage", true),
	chromedp.Flag("disable-extensions", true),
	chromedp.Flag("disable-features", "site-per-process,Translate,BlinkGenPropertyTrees"),
	chromedp.Flag("disable-hang-monitor", true),
	chromedp.Flag("disable-ipc-flooding-protection", true),
	chromedp.Flag("disable-popup-blocking", true),
	chromedp.Flag("disable-prompt-on-repost", true),
	chromedp.Flag("disable-renderer-backgrounding", true),
	chromedp.Flag("disable-sync", true),
	chromedp.Flag("force-color-profile", "srgb"),
	chromedp.Flag("metrics-recording-only", true),
	chromedp.Flag("safebrowsing-disable-auto-update", true),
	chromedp.Flag("enable-automation", true),
	chromedp.Flag("password-store", "basic"),
	chromedp.Flag("use-mock-keychain", true),
}

func (ch *ChromeCtl) Launch(args []string) string {

	log.Printf("launching chrome: %v", args)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), myExecAllocatorOptions[:]...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	// run browser
	page, err := url.Parse("http://10.42.0.2:4646/")
	pageTitle := ""
	if err != nil {
		log.Fatal("oops")
	}
	if err := chromedp.Run(taskCtx,
		chromedp.Navigate(page.String()),
		chromedp.Title(&pageTitle),
	); err != nil {
		log.Fatal(err)
	}
	log.Printf("page title:%s", pageTitle)

	<-allocCtx.Done()
	return "ok"
}
