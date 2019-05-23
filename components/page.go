package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Page struct {
	vecty.Core
	Content vecty.Component
}

func (c *Page) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(vecty.Class("content")),
		c.Content,
	)
}
