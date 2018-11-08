// +build wasm

package qrcode

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/gowasm/vecty/prop"
	qrcode "github.com/skip2/go-qrcode"
	"github.com/vincent-petithory/dataurl"
)

// QRCode ...
type QRCode struct {
	vecty.Core
	Data string
	Size int
}

func (c *QRCode) Render() vecty.ComponentOrHTML {
	if c.Size == 0 {
		c.Size = 256
	}
	png, _ := qrcode.Encode(c.Data, qrcode.Medium, c.Size)
	uri := dataurl.New([]byte(png), "image/png").String()
	return elem.Image(
		vecty.Markup(
			vecty.MarkupIf(len(c.Data) > 0,
				prop.Src(uri),
			),
		),
	)
}
