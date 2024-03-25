/*
Copyright © 2024 <https://github.com/gccruz93>
*/
//go:generate goversioninfo
package main

import (
	"bytes"
	"desktop-buddy/assets"
	_ "embed"
	"image"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	screenHeight    = 0
	screenWidth     = 0
	frameCount      = 0
	audioContext    *audio.Context
	pets            []*Monster
	mplusNormalFont font.Face
	cfg             Cfg
	nextSpawn       = 0
)

func init() {
	cfg.Load()

	audioContext = audio.NewContext(sampleRate)
	loadedGifs = make(map[string][]*ebiten.Image)
	loadedAudios = make(map[string][]byte)

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	mplusNormalFont, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    14,
		DPI:     62,
		Hinting: font.HintingVertical,
	})
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
	mx, my := ebiten.CursorPosition()

	// if frameCount%60 == 0 {
	// 	log.Println(fmt.Sprintln("pets alive:", len(pets)))
	// }

	if frameCount%(nextSpawn*ebiten.TPS()) == 0 && len(pets) < cfg.PetsMax {
		SpawnRandom(1)
	}

	petsAlive := pets[:0]
	for _, e := range pets {
		e.Update()

		// if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		// 	keepPets = append(keepPets, e)
		// }
		if isMouseHover(mx, my, e.x, e.y, e.x+float64(e.width), e.y+float64(e.height)) {
			// ebiten.SetCursorMode(ebiten.CursorModeHidden)
			e.drawName = true
			e.hp--
		} else {
			// ebiten.SetCursorMode(ebiten.CursorModeVisible)
			e.drawName = false
			if e.hp < e.maxhp {
				e.hp++
			}
		}

		if e.hp > 0 {
			petsAlive = append(petsAlive, e)
		}
	}
	pets = petsAlive

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, e := range pets {
		e.Draw(screen)
	}
}

func main() {
	ebiten.SetWindowTitle("Ragnarok Buddy")
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowMousePassthrough(true)
	sw, sh := ebiten.ScreenSizeInFullscreen()
	screenHeight = sh - cfg.ScreenPaddingBottom
	screenWidth = sw * cfg.ScreenMonitors
	ebiten.SetWindowSize(screenWidth, screenHeight)

	imgByte, err := assets.Assets.ReadFile("icon.jpg")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatal(err)
	}
	iconImages := []image.Image{img}
	ebiten.SetWindowIcon(iconImages)

	op := &ebiten.RunGameOptions{}
	op.ScreenTransparent = true
	// op.SkipTaskbar = true

	SpawnRandom(random(1, cfg.PetsMax))

	if err := ebiten.RunGameWithOptions(&Game{}, op); err != nil {
		log.Fatal(err)
	}
}
