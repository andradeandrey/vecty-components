// +build !wasm

package camera

import (
	"github.com/gopherjs/gopherwasm/js"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
)

// Camera ...
type Camera struct {
	vecty.Core
	ID         string
	Constraint map[string]interface{}
	stream     js.Value
}

// Render ...
func (c *Camera) Render() vecty.ComponentOrHTML {
	return elem.Video(
		vecty.Markup(
			prop.ID(c.ID),
			vecty.Attribute("autoplay", true),
			vecty.Attribute("playsinline", true),
		),
	)
}

// Mount ...
func (c *Camera) Mount() {
	if c.Constraint == nil {
		c.Constraint = map[string]interface{}{
			"video": true,
			"audio": true,
		}
	}
	mediaDevices := js.Global().Get("navigator").Get("mediaDevices")
	mediaDevices.Call("getUserMedia", c.Constraint).Call(
		"then",
		js.NewCallback(func(stream []js.Value) {
			c.stream = stream[0]
			video := js.Global().Get("document").Call("getElementById", c.ID)
			video.Set("srcObject", c.stream)
		}),
	)
}

// Unmount ...
func (c *Camera) Unmount() {
	if c.stream != js.Null() {
		c.stream.Call("getTracks").Call(
			"forEach",
			js.NewCallback(func(s []js.Value) {
				s[0].Call("stop")
			}),
		)
	}
}
