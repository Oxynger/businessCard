package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type App struct {
	vecty.Core
}

func (c *App) Render() vecty.ComponentOrHTML {
	return elem.Body(
		&Header{
			Title: "cat oxynger",
		},
		&Page{
			Content: &CV{},
		},
	)
}
