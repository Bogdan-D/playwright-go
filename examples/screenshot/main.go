package main

import (
	"log"

	"github.com/mxschmitt/playwright-go"
)

func assertErrorToNilf(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}
}

func main() {
	pw, err := playwright.Run()
	assertErrorToNilf("could not launch playwright: %w", err)
	browser, err := pw.Chromium.Launch()
	assertErrorToNilf("could not launch Chromium: %w", err)
	context, err := browser.NewContext()
	assertErrorToNilf("could not create context: %w", err)
	page, err := context.NewPage()
	assertErrorToNilf("could not create page: %w", err)
	_, err = page.Goto("http://whatsmyuseragent.org/")
	assertErrorToNilf("could not goto: %w", err)
	_, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("foo.png"),
	})
	assertErrorToNilf("could not create screenshot: %w", err)
	err = browser.Close()
	assertErrorToNilf("could not close browser: %w", err)
	err = pw.Stop()
	assertErrorToNilf("could not stop Playwright: %w", err)
}