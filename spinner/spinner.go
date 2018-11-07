package spinner

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/vincent-petithory/dataurl"
)

const style = `
.spinner {
	-webkit-animation: rotator 1.4s linear infinite;
	animation: rotator 1.4s linear infinite;
  }
  
  @-webkit-keyframes rotator {
	0% {
	  -webkit-transform: rotate(0deg);
	  transform: rotate(0deg);
	}
	100% {
	  -webkit-transform: rotate(270deg);
	  transform: rotate(270deg);
	}
  }
  
  @keyframes rotator {
	0% {
	  -webkit-transform: rotate(0deg);
	  transform: rotate(0deg);
	}
	100% {
	  -webkit-transform: rotate(270deg);
	  transform: rotate(270deg);
	}
  }
  .path {
	stroke-dasharray: 187;
	stroke-dashoffset: 0;
	-webkit-transform-origin: center;
	transform-origin: center;
	-webkit-animation: dash 1.4s ease-in-out infinite,
	  colors 5.6s ease-in-out infinite;
	animation: dash 1.4s ease-in-out infinite, colors 5.6s ease-in-out infinite;
  }
  
  @-webkit-keyframes colors {
	0% {
	  stroke: #4285f4;
	}
	25% {
	  stroke: #de3e35;
	}
	50% {
	  stroke: #f7c223;
	}
	75% {
	  stroke: #1b9a59;
	}
	100% {
	  stroke: #4285f4;
	}
  }
  
  @keyframes colors {
	0% {
	  stroke: #4285f4;
	}
	25% {
	  stroke: #de3e35;
	}
	50% {
	  stroke: #f7c223;
	}
	75% {
	  stroke: #1b9a59;
	}
	100% {
	  stroke: #4285f4;
	}
  }
  @-webkit-keyframes dash {
	0% {
	  stroke-dashoffset: 187;
	}
	50% {
	  stroke-dashoffset: 46.75;
	  -webkit-transform: rotate(135deg);
	  transform: rotate(135deg);
	}
	100% {
	  stroke-dashoffset: 187;
	  -webkit-transform: rotate(450deg);
	  transform: rotate(450deg);
	}
  }
  @keyframes dash {
	0% {
	  stroke-dashoffset: 187;
	}
	50% {
	  stroke-dashoffset: 46.75;
	  -webkit-transform: rotate(135deg);
	  transform: rotate(135deg);
	}
	100% {
	  stroke-dashoffset: 187;
	  -webkit-transform: rotate(450deg);
	  transform: rotate(450deg);
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
			vecty.Attribute("style", "margin: auto"),
		),
		vecty.Tag("svg",
			vecty.Markup(
				vecty.Class("spinner"),
				vecty.Attribute("width", "65px"),
				vecty.Attribute("height", "65px"),
				vecty.Attribute("viewBox", "0 0 66 66"),
				vecty.Attribute("xmlns", "http://www.w3.org/2000/svg"),
			),
			vecty.Tag("circle",
				vecty.Markup(
					vecty.Class("path"),
					vecty.Attribute("fill", "none"),
					vecty.Attribute("stroke-width", "6"),
					vecty.Attribute("stroke-linecap", "round"),
					vecty.Attribute("cx", "33"),
					vecty.Attribute("cy", "33"),
					vecty.Attribute("r", "30"),
				),
			),
		),
	)
}
