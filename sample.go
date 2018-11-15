// +build !wasm

package main

import (
	"log"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/gopherwasm/js"
	"github.com/nobonobo/vecty-components/qrcode"
	"github.com/nobonobo/vecty-components/qrscanner"
	"github.com/nobonobo/vecty-components/camera"
	"github.com/nobonobo/vecty-components/spinner"
	"github.com/nobonobo/vecty-components/toggle"
)

// Index ...
type Index struct {
	vecty.Core
}

// Render ...
func (c *Index) Render() vecty.ComponentOrHTML {
	qc := &qrcode.QRCode{Data: "http://www.google.co.jp"}
	return elem.Body(
		elem.Div(
			elem.Heading3(vecty.Text("QR Code(detected)")),
			qc,
			&spinner.Spinner{},
			&toggle.Toggle{Label:"Toggle Switch!"},
		),
		elem.Div(
			elem.Heading3(vecty.Text("QR Scanner Camera")),
			&qrscanner.Scanner{
				ID: "scanner",
				Callback: js.NewCallback(func(args []js.Value) {
					qc.Data = args[0].String()
					log.Println(qc.Data)
					vecty.Rerender(qc)
				}),
			},
		),
		elem.Div(
			elem.Heading3(vecty.Text("Normal Camera")),
			&camera.Camera{ID:"camera"},
		),
	)
}

func main() {
	log.Println("main start")
	vecty.RenderBody(&Index{})
}
