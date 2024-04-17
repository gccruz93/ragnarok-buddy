package helpers

import (
	_ "embed"
	"image"
	"image/draw"
	"image/gif"

	"github.com/hajimehoshi/ebiten/v2"
)

type CustomGif struct {
	Height, Width float64
	Length        int
	Frames        []*ebiten.Image
}

func SplitAnimatedGIF(gif *gif.GIF) *CustomGif {
	imgWidth, imgHeight := GetGifDimensions(gif)

	customGif := &CustomGif{
		Height: float64(imgHeight),
		Width:  float64(imgWidth),
	}

	for _, srcImg := range gif.Image {
		frame := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
		draw.Draw(frame, frame.Bounds(), srcImg, image.Point{}, draw.Over)
		customGif.Frames = append(customGif.Frames, ebiten.NewImageFromImage(frame))
	}
	customGif.Length = len(customGif.Frames)

	return customGif
}

func GetGifDimensions(gif *gif.GIF) (x, y int) {
	var lowestX, lowestY, highestX, highestY int
	for _, img := range gif.Image {
		if img.Rect.Min.X < lowestX {
			lowestX = img.Rect.Min.X
		}
		if img.Rect.Min.Y < lowestY {
			lowestY = img.Rect.Min.Y
		}
		if img.Rect.Max.X > highestX {
			highestX = img.Rect.Max.X
		}
		if img.Rect.Max.Y > highestY {
			highestY = img.Rect.Max.Y
		}
	}
	return highestX - lowestX, highestY - lowestY
}
