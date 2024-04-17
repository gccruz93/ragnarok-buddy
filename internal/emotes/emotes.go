package emotes

import (
	"ragnarok-buddy/internal/core"
	"ragnarok-buddy/pkg/helpers"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	EmoteActive *Emote
	rarityList  []int
)

var NextSpawnTick = 1

type Emote struct {
	alpha                 float32
	anchorMargin          float64
	frameIndex, frameTime int
	gif                   *helpers.CustomGif
	X, Y                  float64
	loops, loopsCount     int
	ended                 bool
}

func (e *Emote) Update() {
	if core.FrameTick%e.frameTime == 0 {
		if !e.ended || e.ended && e.loops == 0 {
			e.frameIndex = (e.frameIndex + 1) % e.gif.Length
			if e.frameIndex == 0 {
				e.loopsCount++
				e.ended = e.loopsCount >= e.loops
			}
		}
	}

	if !e.ended && e.alpha < 1 {
		e.alpha += 0.04
	}
}

func (e *Emote) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(e.alpha)
	op.GeoM.Translate(e.X, e.Y-e.gif.Height-e.anchorMargin)
	screen.DrawImage(e.gif.Frames[e.frameIndex], op)
}
