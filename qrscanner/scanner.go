// +build !wasm

package qrscanner

import (
	"log"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/gopherjs/gopherwasm/js"
	"github.com/vincent-petithory/dataurl"
)

func init() {
	ch := make(chan struct{})
	b, err := FSByte(false, "/dist/main.js")
	if err != nil {
		panic(err)
	}
	s := js.Global().Get("document").Call("createElement", "script")
	s.Set("src", dataurl.New(b, "application/javascript").String())
	js.Global().Get("document").Get("head").Call("appendChild", s)
	s.Call("addEventListener", "load", js.NewCallback(func([]js.Value) {
		close(ch)
	}))
	<-ch
	log.Println("script loaded")
}

// Scanner ...
type Scanner struct {
	vecty.Core
	ID       string
	scanner  js.Value
	Callback js.Callback
}

// Render ...
func (c *Scanner) Render() vecty.ComponentOrHTML {
	return elem.Video(
		vecty.Markup(
			prop.ID(c.ID),
			vecty.Attribute("autoplay", true),
			vecty.Attribute("playsinline", true),
		),
	)
}

// Mount ...
func (c *Scanner) Mount() {
	video := js.Global().Get("document").Call("getElementById", c.ID)
	qr := js.Global().Get("QrScanner")
	log.Println(qr)
	scanner := qr.New(video, c.Callback)
	scanner.Call("start")
	c.scanner = scanner
	log.Println("scanner start")
}

// Unmount ...
func (c *Scanner) Unmount() {
	if c.scanner != js.Null() {
		c.scanner.Call("stop")
		log.Println("scanner stop")
	}
}
