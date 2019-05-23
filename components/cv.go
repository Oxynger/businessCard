package components

import (
	"businessCard/store"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type CV struct {
	vecty.Core
	CV string
}

func (c *CV) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(vecty.Class("cv")),
		vecty.Text(c.CV),
	)
}

func (c *CV) Mount() {
	c.CV = store.GetCV()
	vecty.Rerender(c)
}
