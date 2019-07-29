package main

import (
	"math/rand"
	"time"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func random(quant int) []int {
	rand.Seed(time.Now().UnixNano())
	return rand.Perm(quant)[:quant]
}

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "SorTei.oN",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(random)
	app.Run()
}
