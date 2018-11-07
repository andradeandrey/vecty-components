package camera

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/gowasm/vecty/prop"
	"github.com/gowasm/gopherwasm/js"
)

// Camera ...
type Camera struct {
	vecty.Core
	ID     string
	stream *js.Object
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
	mediaDevices := js.Global.Get("navigator").Get("mediaDevices")
	mediaDevices.Call("getUserMedia", js.M{"video": true}).Call("then", func(stream *js.Object) {
		c.stream = stream
		//src := js.Global.Get("window").Get("HTMLMediaElement").Call("srcObject", stream)
		video := js.Global.Get("document").Call("getElementById", c.ID)
		video.Set("srcObject", stream)
	})
}

// Unmount ...
func (c *Camera) Unmount() {
	if c.stream != nil {
		c.stream.Call("getTracks").Call("forEach", func(s *js.Object) {
			s.Call("stop")
		})
	}
}
