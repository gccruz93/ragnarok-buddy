//go:generate goversioninfo
package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/jpeg"
	"log"
	"ragnarok-buddy/assets"
	"ragnarok-buddy/internal/core"
	"ragnarok-buddy/internal/emotes"
	"ragnarok-buddy/internal/mobs"
	"ragnarok-buddy/pkg/helpers"

	"github.com/energye/systray"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func init() {
	core.Cfg.Load()
	mobs.LoadConfig()
	emotes.LoadConfig()

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	core.Font, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    14,
		DPI:     62,
		Hinting: font.HintingVertical,
	})

	core.AudioContext = audio.NewContext(core.SampleRate)
}

type Game struct{}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	if !core.IsRunning {
		return ebiten.Termination
	}

	core.FrameTick++
	core.Mx, core.My = ebiten.CursorPosition()

	mobs.Update()
	emotes.Update()

	if core.FrameTick%emotes.NextSpawnTick == 0 && emotes.EmoteActive == nil {
		emotes.SpawnRandom()
		mobs.List[helpers.Random(0, len(mobs.List)-1)].Emote = emotes.EmoteActive
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	mobs.Draw(screen)
	emotes.Draw(screen)
}

func setScreenArea() {
	sw, sh := ebiten.ScreenSizeInFullscreen()
	screenHeight := sh - core.Cfg.ScreenPaddingBottom
	screenWidth := sw * core.Cfg.ScreenMonitors
	ebiten.SetWindowSize(screenWidth, screenHeight)
	core.ScreenHeight = float64(screenHeight)
	core.ScreenWidth = float64(screenWidth)
}

func main() {
	ebiten.SetWindowTitle(core.Title)
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowMousePassthrough(true)
	setScreenArea()

	img, _, err := image.Decode(bytes.NewReader(assets.Icon))
	if err != nil {
		log.Fatal(err)
	}
	iconImages := []image.Image{img}
	ebiten.SetWindowIcon(iconImages)

	op := &ebiten.RunGameOptions{}
	op.ScreenTransparent = true
	op.SkipTaskbar = core.Cfg.SkipTaskbar

	go func() {
		systray.Run(onReady, onExit)
	}()
	if err := ebiten.RunGameWithOptions(&Game{}, op); err != nil {
		log.Fatal(err)
	}
}
