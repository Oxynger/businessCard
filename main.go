package main

import (
	"businessCard/components"
	"businessCard/store"

	"github.com/gopherjs/vecty"
)

func main() {
	must(store.FetchCV())

	vecty.SetTitle("Home")
	vecty.AddStylesheet("main.css")
	vecty.RenderBody(&components.App{})
}

func must(e error) {
	if e != nil {
		panic(e)
	}
}
