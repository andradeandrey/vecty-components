package qrscanner

//go:generate bash -c "npm install -D webpack webpack-cli raw-loader && node_modules/.bin/webpack --mode=production src/index.js --output index.inc.js"

import (
	"log"

	"github.com/gowasm/gopherwasm/js"
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/gowasm/vecty/prop"
)

// Scanner ...
type Scanner struct {
	vecty.Core
	ID       string
	scanner  *js.Object
	Callback func(result *js.Object)
}

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
	video := js.Global.Get("document").Call("getElementById", c.ID)
	qr := js.Global.Get("QrScanner")
	scanner := qr.New(video, c.Callback)
	scanner.Call("start")
	c.scanner = scanner
	log.Println("scanner start")
}

// Unmount ...
func (c *Scanner) Unmount() {
	if c.scanner != nil {
		c.scanner.Call("stop")
		log.Println("scanner stop")
	}
}
