package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"main/pkgs"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()

	r.GET("/*url", Handle)

	r.Run(":8080")
}

func Handle(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}
		return
	}()

	trans := pkgs.UrlTransformer{}
	options, url := trans.Handle(c.Param("url"))

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var buf []byte

	// capture entire browser viewport, returning png with quality=90
	err := chromedp.Run(ctx, FullScreenshot(url, options, &buf))
	if err != nil {
		panic(err)
	}

	contentType := "image/png"
	if options["quality"] != "100" {
		contentType = "image/jpg"
	}

	c.Data(http.StatusOK, contentType, buf)
}

// FullScreenshot takes a screenshot of the entire browser viewport.
//
// Note: chromedp.FullScreenshot overrides the device's emulation settings. Use
// device.Reset to reset the emulation and viewport settings.
func FullScreenshot(urlstr string, options map[string]string, res *[]byte) chromedp.Tasks {
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
