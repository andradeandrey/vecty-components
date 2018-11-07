package toggle

import (
	"github.com/gowasm/vecty"
	"github.com/gowasm/vecty/elem"
	"github.com/gowasm/vecty/event"
	"github.com/gowasm/vecty/prop"
	"github.com/vincent-petithory/dataurl"
)

const style = `
.form-check {
	position: relative;
	display: block;
	margin-bottom: 0.75rem;
}

.form-check-toggle {
	position: relative;
	padding-left: 0;
	line-height: 30px;
}

.form-check-toggle input {
	display: block;
	position: absolute;
	top: 0;
	right: 0;
	bottom: 0;
	left: 0;
	width: 0%;
	height: 0%;
	margin: 0;
	cursor: pointer;
	opacity: 0;
	filter: alpha(opacity=0);
}

.form-check-toggle input + span {
	cursor: pointer;
	user-select: none;
	height: 30px;
	margin-left: 70px;
	display: block;
}

.form-check-toggle input + span:before {
	content: "";
	position: absolute;
	left: 0;
	display: inline-block;
	height: 30px;
	width: 50px;
	background: #fff;
	border: solid 1px #babbbb;
	transition: background 0.3s ease-in-out, border-color 0.3s ease-in-out;
	border-radius: 15px;
}

.form-check-toggle input + span:after {
	width: 28px;
	height: 28px;
	margin-top: 1px;
	margin-left: 1px;
	border-radius: 50%;
	position: absolute;
	left: 0;
	top: 0;
	display: block;
	background: #fff;
	transition: margin-left 0.1s ease-in-out, box-shadow 0.1s ease-in-out;
	text-align: center;
	font-weight: bold;
	content: "";
	box-shadow: 0 0 2px rgba(0, 0, 0, 0.2), 0 0 5px rgba(0, 0, 0, 0.05);
}

.form-check-toggle input:checked + span:after {
	content: "";
	margin-left: 21px;
	box-shadow: none;
}

.form-check-toggle input:checked + span:before {
	background-color: #86be4e;
	border-color: #86be4e;
	transition: background 0.1s ease-in-out, border-color 0.1s ease-in-out;
}
`

func init() {
	vecty.AddStylesheet(dataurl.New([]byte(style), "text/css").String())
}

// Toggle ...
type Toggle struct {
	vecty.Core
	Label string `vecty:"prop"`
	Value bool   `vecty:"prop"`
}

// Render ...
func (c *Toggle) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("form-check"),
		),
		elem.Label(
			vecty.Markup(
				vecty.ClassMap{
					"form-check-label":  true,
					"form-check-toggle": true,
				},
			),
			elem.Input(
				vecty.Markup(
					vecty.Class("form-check-input"),
					prop.Type("checkbox"),
					prop.Checked(c.Value),
					event.Change(func(ev *vecty.Event) {
						c.Value = ev.Get("target").Get("checked").Bool()
					}),
				),
			),
			elem.Span(
				vecty.Text(c.Label),
			),
		),
	)
}
