package main

import (
	"fmt"
	"math/rand"
)

const (
	width  = 10
	height = 10
)

type Tile struct {
	id int
}

var tiles = []Tile{
	{0}, {1}, {2}, {3},
}

var adjacencyRules = map[int][]int{
	0: {1, 2},
	1: {0, 3},
	2: {0, 3},
	3: {1, 2},
}

func main() {
	grid := make([][]Tile, height)
	for i := range grid {
		grid[i] = make([]Tile, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grid[x][y] = collapseTile(x, y, grid)
		}
	}

	printGrid(grid)
}

func collapseTile(x, y int, grid [][]Tile) Tile {
	possibleTiles := tiles
	if x > 0 {
		possibleTiles = filterTiles(possibleTiles, grid[y][x-1].id)
	}
	if y > 0 {
		possibleTiles = filterTiles(possibleTiles, grid[y-1][x].id)
	}

	if len(possibleTiles) == 0 {
		return Tile{id: -1}
	}

	return possibleTiles[rand.Intn(len(possibleTiles))]
}

func filterTiles(possibleTiles []Tile, adjacentTileID int) []Tile {
	var filteredTiles []Tile
	for _, tile := range possibleTiles {
		for _, allowedID := range adjacencyRules[adjacentTileID] {
			if tile.id == allowedID {
				filteredTiles = append(filteredTiles, tile)
			}
		}
	}

	return filteredTiles
}

func printGrid(grid [][]Tile) {
	for _, row := range grid {
		for _, tile := range row {
			fmt.Print(tile.id, " ")
		}
		fmt.Println()
	}
}
