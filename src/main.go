/*
Copyright © 2024 <https://github.com/gccruz93>
*/
//go:generate goversioninfo
package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/jpeg"
	"log"

	"github.com/energye/systray"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	screenHeight    = 0
	screenWidth     = 0
	title           = "Ragnarok Buddy"
	frameCount      = 0
	audioContext    *audio.Context
	mobs            []*Mob
	mplusNormalFont font.Face
	cfg             Cfg
	nextSpawn       = 0
)

func init() {
	cfg.Load()

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	mplusNormalFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    14,
		DPI:     62,
		Hinting: font.HintingVertical,
	})

	audioContext = audio.NewContext(sampleRate)
	loadedGifs = make(map[string][]*ebiten.Image)
	loadedAudios = make(map[string][]byte)
}

type Game struct{}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func isMouseHover(mx, my int, x, y, x2, y2 float64) bool {
	return (float64(mx) >= x && float64(mx) <= x2) && (float64(my) >= y && float64(my) <= y2)
}

func (g *Game) Update() error {
	frameCount++

	switch cfg.Gamemode {
	case 0:
		// normal
		g.UpdateGamemode0()
	case 1:
		// maps
		g.UpdateGamemode1()
	default:
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, e := range mobs {
		e.Draw(screen)
	}
}

func main() {
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowMousePassthrough(true)
	sw, sh := ebiten.ScreenSizeInFullscreen()
	screenHeight = sh - cfg.ScreenPaddingBottom
	screenWidth = sw * cfg.ScreenMonitors
	ebiten.SetWindowSize(screenWidth, screenHeight)

	iconBytes, err := assets.ReadFile("assets/icon.jpg")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(bytes.NewReader(iconBytes))
	if err != nil {
		log.Fatal(err)
	}
	iconImages := []image.Image{img}
	ebiten.SetWindowIcon(iconImages)

	op := &ebiten.RunGameOptions{}
	op.ScreenTransparent = true
	op.SkipTaskbar = cfg.SkipTaskbar

	nextSpawn = 1

	trayStart, trayEnd := systray.RunWithExternalLoop(onReady, onExit)
	trayStart()
	if err := ebiten.RunGameWithOptions(&Game{}, op); err != nil {
		trayEnd()
		log.Fatal(err)
	}
}
