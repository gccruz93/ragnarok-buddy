package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	x, y, vx, vy                                    float64
	height, width                                   int
	frameIndex, frameSpeed, frameAudio, frameLength int
	gif                                             string
	audio                                           string
	alpha                                           float32
}

func (e *Entity) Update() {
	e.y = float64(screenHeight - e.height)
	e.y += e.vy

	if frameCount%e.frameSpeed == 0 {
		e.frameIndex = (e.frameIndex + 1) % e.frameLength

		if e.frameIndex == e.frameAudio && e.audio != "" && cfg.MainVolume > 0 {
			p := audioContext.NewPlayerFromBytes(loadedAudios[e.audio])
			p.SetVolume(cfg.MainVolume)
			p.Play()
		}
	}

	if e.alpha != 1 {
		e.alpha += 0.01
		if e.alpha >= 1 {
			e.alpha = 1
		}
	}
}

func (e *Entity) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(e.alpha)
	op.GeoM.Translate(e.x, e.y)
	screen.DrawImage(loadedGifs[e.gif][e.frameIndex], op)
}

func (e *Entity) SetGif(path string, speed int) {
	loadGif(path)
	e.gif = path
	e.width = loadedGifs[e.gif][0].Bounds().Dx()
	e.height = loadedGifs[e.gif][0].Bounds().Dy()
	e.frameLength = len(loadedGifs[e.gif])
	e.frameSpeed = speed
	e.frameIndex = 0
}

func (e *Entity) SetAudio(path string, frame int) {
	if path != "" {
		loadAudio(path)
	}
	e.audio = path
	e.frameAudio = frame
}
