package core

import (
	"ragnarok-buddy/internal/config"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"golang.org/x/image/font"
)

var (
	Title        = "Ragnarok Buddy"
	IsRunning    = true
	ScreenHeight = 0.0
	ScreenWidth  = 0.0
	FrameTick    = 0
	AudioContext *audio.Context
	Font         font.Face
)

var (
	Cfg config.Config
)

const SampleRate = 48000

var Mx int
var My int
