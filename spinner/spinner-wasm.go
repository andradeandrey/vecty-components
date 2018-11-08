// +build wasm

package spinner

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/vincent-petithory/dataurl"
)

const style = `
.loadingspinner {
  margin: auto;
  pointer-events: none;
  width: 2.5em;
  height: 2.5em;
  border: 0.4em solid transparent;
  border-color: #eee;
  border-top-color: #3e67ec;
  border-radius: 50%;
  -webkit-animation: loadingspin 1s linear infinite;
  animation: loadingspin 1s linear infinite;
}

@-webkit-keyframes loadingspin {
  100% {
    -webkit-transform: rotate(360deg);
    transform: rotate(360deg);
  }
}

@keyframes loadingspin {
  100% {
    -webkit-transform: rotate(360deg);
    transform: rotate(360deg);
  }
}
`

func init() {
	vecty.AddStylesheet(dataurl.New([]byte(style), "text/css").String())
}

// Spinner ...
type Spinner struct {
	vecty.Core
}

// Render ...
func (c *Spinner) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("d-flex", "justify-content-center", "align-content-center"),
			vecty.Style("min-height", "2.5em"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("loadingspinner"),
			),
		),
	)
}
