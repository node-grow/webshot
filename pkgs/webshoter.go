package pkgs

import (
	"context"
	"github.com/chromedp/chromedp"
	"strconv"
)

type Webshoter struct {
}

func (w *Webshoter) Shot(url string, options map[string]string, buf *[]byte) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// capture entire browser viewport, returning png with quality=90
	err := chromedp.Run(ctx, w.fullScreenshot(url, options, buf))
	if err != nil {
		panic(err)
	}
}

// FullScreenshot takes a screenshot of the entire browser viewport.
//
// Note: chromedp.FullScreenshot overrides the device's emulation settings. Use
// device.Reset to reset the emulation and viewport settings.
func (w *Webshoter) fullScreenshot(urlstr string, options map[string]string, res *[]byte) chromedp.Tasks {
	width, _ := strconv.ParseInt(options["width"], 10, 0)
	height, _ := strconv.ParseInt(options["height"], 10, 0)
	quality, _ := strconv.Atoi(options["quality"])
	println(width, height, quality)

	return chromedp.Tasks{
		chromedp.EmulateViewport(width, height),
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}
