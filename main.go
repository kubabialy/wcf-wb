package main

import (
	"fmt"
	"wave_function_collapse/render"
	"wave_function_collapse/tile"
)

const (
	width  = 25
	height = 25
)

func main() {
	grid := tile.InitGrid(width, height)

	image, err := render.Render(grid)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := render.SaveImage(image, "output.png"); err != nil {
		fmt.Println(err)
		return
	}
}
