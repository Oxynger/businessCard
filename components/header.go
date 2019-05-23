package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Header struct {
	vecty.Core
	Title string
}

func (c *Header) Render() vecty.ComponentOrHTML {
	return elem.Header(
		elem.Div(
			vecty.Markup(vecty.Class("greeter")),
			elem.Heading1(vecty.Text(c.Title)),
		),
	)
}
