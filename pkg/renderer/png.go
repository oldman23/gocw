package renderer

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/font/inconsolata"
	"golang.org/x/image/math/fixed"

	"gocrossword/pkg/crossword"
)

const CellSize = 40

type PNGRenderer struct{}

func NewPNGRenderer() *PNGRenderer {
	return &PNGRenderer{}
}

func (r *PNGRenderer) Render(c *crossword.Crossword, outputPath string) error {
	return r.renderInternal(c, outputPath, false)
}

func (r *PNGRenderer) RenderEmpty(c *crossword.Crossword, outputPath string) error {
	return r.renderInternal(c, outputPath, true)
}

func (r *PNGRenderer) renderInternal(c *crossword.Crossword, outputPath string, empty bool) error {
	width := c.Cols*CellSize + (c.Cols-1)
	height := c.Rows*CellSize + (c.Rows-1)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	
	white := color.RGBA{255, 255, 255, 255}
	black := color.RGBA{0, 0, 0, 255}
	grey := color.RGBA{128, 128, 128, 255}

	draw.Draw(img, img.Bounds(), &image.Uniform{grey}, image.Point{}, draw.Src)

	for row := 0; row < c.Rows; row++ {
		for col := 0; col < c.Cols; col++ {
			x := col*(CellSize+1)
			y := row*(CellSize+1)
			
			cellRect := image.Rect(x, y, x+CellSize, y+CellSize)
			
			cell := strings.TrimSpace(c.GetCell(row, col))
			if cell == "" {
				draw.Draw(img, cellRect, &image.Uniform{black}, image.Point{}, draw.Src)
			} else {
				draw.Draw(img, cellRect, &image.Uniform{white}, image.Point{}, draw.Src)
				
				number := strings.TrimSpace(c.GetNumber(row, col))
				if number != "" {
					r.drawNumber(img, number, x, y, black)
				}
				
				if !empty && cell != "" {
					r.drawCharacter(img, cell, x, y, CellSize, black)
				}
			}
			
			if col < c.Cols-1 {
				lineRect := image.Rect(x+CellSize, y, x+CellSize+1, y+CellSize)
				draw.Draw(img, lineRect, &image.Uniform{grey}, image.Point{}, draw.Src)
			}
			
			if row < c.Rows-1 {
				lineRect := image.Rect(x, y+CellSize, x+CellSize, y+CellSize+1)
				draw.Draw(img, lineRect, &image.Uniform{grey}, image.Point{}, draw.Src)
			}
		}
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	return png.Encode(outFile, img)
}

func (r *PNGRenderer) drawNumber(img *image.RGBA, number string, x, y int, col color.Color) {
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
	}
	
	d.Dot = fixed.Point26_6{
		X: fixed.Int26_6((x + 2) * 64),
		Y: fixed.Int26_6((y + 10) * 64),
	}
	
	d.DrawString(number)
}

func (r *PNGRenderer) drawCharacter(img *image.RGBA, char string, x, y, cellSize int, col color.Color) {
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: inconsolata.Bold8x16,
	}
	
	charWidth := d.MeasureString(strings.ToUpper(char))
	charX := x + (cellSize-charWidth.Round())/2
	charY := y + cellSize/2 + 6
	
	d.Dot = fixed.Point26_6{
		X: fixed.Int26_6(charX * 64),
		Y: fixed.Int26_6(charY * 64),
	}
	
	d.DrawString(strings.ToUpper(char))
}