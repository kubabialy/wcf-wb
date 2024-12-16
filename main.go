package main

import (
	"fmt"
	"wave_function_collapse/render"
	"wave_function_collapse/tile"
)

const (
	width  = 100
	height = 100
)

func main() {
	grid := make([][]tile.Tile, height)
	for i := range grid {
		grid[i] = make([]tile.Tile, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grid[x][y] = tile.CollapseTile(x, y, grid)
		}
	}

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
