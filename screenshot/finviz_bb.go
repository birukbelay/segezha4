package screenshot

import (
	"bytes"
	"context"
	"image"
	"image/png"
	"log"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

// MakeScreenshotForFinvizBB description
func MakeScreenshotForFinvizBB(linkURL string) []byte {
	ctx0, cancel0 := chromedp.NewRemoteAllocator(context.Background(), getWebSocketDebuggerUrl())
	defer cancel0()
	ctx1, cancel1 := chromedp.NewContext(ctx0)
	defer cancel1()
	// start the browser without a timeout
	if err := chromedp.Run(ctx1); err != nil {
		log.Println(err)
		return nil
	}
	ctx2, cancel2 := context.WithTimeout(ctx1, 50*time.Second)
	defer cancel2()
	var buf1, buf2 []byte
	sel1 := "body > div.content.is-index > div.container > table > tbody > tr > td table:nth-child(1)"
	sel2 := "body > div.content.is-index > div.container > table > tbody > tr > td > #homepage > table > tbody > tr > td > table"
	if err := chromedp.Run(ctx2, func() chromedp.Tasks {
		return chromedp.Tasks{
			network.SetBlockedURLS([]string{"https://dggaenaawxe8z.cloudfront.net/cmp_v2/admiral/finviz.js"}),
			chromedp.Emulate(device.KindleFireHDX),
			chromedp.Navigate(linkURL),
			chromedp.WaitReady("body #homepage"),
			// chromedp.SetAttributeValue(sel1, "style", "margin: 20px 0 0"),
			chromedp.Screenshot(sel1, &buf1, chromedp.NodeVisible),
			chromedp.SetAttributeValue(sel2, "style", "margin: 0 0 4px"),
			chromedp.Screenshot(sel2, &buf2, chromedp.NodeVisible),
		}
	}()); err != nil {
		log.Println(err)
		return nil
	}
	var src image.Image
	if err := glueForFinvizBB(buf1, buf2, &src); err != nil {
		log.Println(err)
		return nil
	}
	buf1 = nil
	buf2 = nil
	// resize to width 800 using Bicubic resampling
	// and preserve aspect ratio
	// res := resize.Resize(800, 0, src, resize.Bicubic)
	// encode
	out := &bytes.Buffer{}
	if err := png.Encode(out, src); err != nil {
		log.Println(err)
		return nil
	}
	src = nil
	// res = nil
	return out.Bytes()
}

func glueForFinvizBB(buf1, buf2 []byte, src *image.Image) error {
	img1, _, err := image.Decode(bytes.NewReader(buf1))
	if err != nil {
		return err
	}
	buf1 = nil
	img2, _, err := image.Decode(bytes.NewReader(buf2))
	if err != nil {
		return err
	}
	buf2 = nil
	glueImages(img1, img2, src)
	return nil
}
