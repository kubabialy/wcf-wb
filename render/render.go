package render

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"wave_function_collapse/tile"
)

func loadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func getImagePath(tileID int) string {
	return fmt.Sprintf("assets/tiles/%d.png", tileID)
}

func Render(grid [][]tile.Tile) (image.Image, error) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return nil, fmt.Errorf("grid is empty")
	}

	firstImagePath := getImagePath(grid[0][0].ID)
	firstImage, err := loadImage(firstImagePath)
	if err != nil {
		return nil, err
	}

	imgWidth := firstImage.Bounds().Dx()
	imgHeight := firstImage.Bounds().Dy()

	outputWidth := imgWidth * len(grid[0])
	outputHeight := imgHeight * len(grid)
	outputImage := image.NewRGBA(image.Rect(0, 0, outputWidth, outputHeight))

	for y, row := range grid {
		for x, t := range row {
			img, err := loadImage(getImagePath(t.ID))
			if err != nil {
				return nil, err
			}

			offset := image.Pt(x*imgWidth, y*imgHeight)
			draw.Draw(outputImage, img.Bounds().Add(offset), img, image.Point{}, draw.Over)
		}
	}

	return outputImage, nil
}

func SaveImage(img image.Image, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}
