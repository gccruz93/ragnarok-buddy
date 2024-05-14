package core

import (
	"ragnarok-buddy/internal/config"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	Title        = "Ragnarok Buddy"
	IsRunning    = true
	ScreenHeight = 0.0
	ScreenWidth  = 0.0
	FrameTick    = 0
	AudioContext *audio.Context
	FaceSource   *text.GoTextFaceSource
	NormalFace   *text.GoTextFace
)

var (
	Cfg config.Config
)

const SampleRate = 48000

var Mx int
var My int
