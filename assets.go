package main

import (
	"bytes"
	"desktop-buddy/assets"
	"image"
	"image/draw"
	"image/gif"
	"io"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

var loadedGifs map[string][]*ebiten.Image
var loadedAudios map[string][]byte

const sampleRate = 48000

func loadGif(name string) {
	if _, ok := loadedGifs[name]; ok {
		return
	}
	file, err := assets.Assets.ReadFile(name)
	if err != nil {
		log.Println("loadGif ERR: " + name)
		return
	}
	loadedGif, _ := gif.DecodeAll(bytes.NewReader(file))
	loadedGifs[name] = splitAnimatedGIF(loadedGif)
}

func loadAudio(name string) {
	if _, ok := loadedAudios[name]; ok {
		return
	}
	file, err := assets.Assets.ReadFile(name)
	if err != nil {
		log.Println("loadAudio ERR: " + name)
		return
	}
	s, _ := wav.DecodeWithSampleRate(sampleRate, bytes.NewReader(file))
	b, _ := io.ReadAll(s)
	loadedAudios[name] = b
}

func splitAnimatedGIF(gif *gif.GIF) []*ebiten.Image {
	var frames []*ebiten.Image
	imgWidth, imgHeight := getGifDimensions(gif)

	for _, srcImg := range gif.Image {
		frame := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
		draw.Draw(frame, frame.Bounds(), srcImg, image.Point{}, draw.Over)
		frames = append(frames, ebiten.NewImageFromImage(frame))
	}

	return frames
}
func getGifDimensions(gif *gif.GIF) (x, y int) {
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
