package tile

import "math/rand"

type Tile struct {
	ID int
}

const (
	Forrest  int = 0
	Mountain int = 1
	Water    int = 2
	Grass    int = 3
	River    int = 4
)

var tiles = []Tile{
	{Forrest}, {Mountain}, {Water}, {Grass}, {River},
}

var adjacencyRules = map[int][]int{
	Forrest:  {Forrest, Mountain, Grass},
	Mountain: {Mountain, Forrest, Grass},
	Water:    {Water, River, Grass},
	Grass:    {Mountain, Forrest, River, Water, Grass},
	River:    {Forrest, Grass, Water, River},
}

func InitGrid(width, height int) [][]Tile {
	grid := make([][]Tile, height)
	for i := range grid {
		grid[i] = make([]Tile, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grid[x][y] = CollapseTile(x, y, grid)
		}
	}

	return grid
}

func CollapseTile(x, y int, grid [][]Tile) Tile {
	possibleTiles := tiles
	if x > 0 {
		possibleTiles = filterTiles(possibleTiles, grid[y][x-1].ID)
	}
	if y > 0 {
		possibleTiles = filterTiles(possibleTiles, grid[y-1][x].ID)
	}

	return possibleTiles[rand.Intn(len(possibleTiles))]
}

func filterTiles(possibleTiles []Tile, adjacentTileID int) []Tile {
	var filteredTiles []Tile
	for _, t := range possibleTiles {
		for _, allowedID := range adjacencyRules[adjacentTileID] {
			if t.ID == allowedID {
				filteredTiles = append(filteredTiles, t)
			}
		}
	}

	return filteredTiles
}
