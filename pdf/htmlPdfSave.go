package pdf

/**
 * @title htmlPdfSave
 * @author CH00SE1
 * @date 2022-09-02 20:21
 */

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
)

func PdfSaveLocal() {
	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// capture pdf
	var buf []byte
	if err := chromedp.Run(ctx, printToPDF(`https://github.com/CH00SE1/httpParse`, &buf)); err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("httpParse.pdf", buf, 0644); err != nil {
		log.Fatal(err)
	}

}

// print a specific pdf page.
func printToPDF(urlstr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
