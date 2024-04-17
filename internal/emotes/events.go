package emotes

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func Update() {
	if EmoteActive == nil {
		return
	}
	EmoteActive.Update()

	if EmoteActive.ended {
		if EmoteActive.alpha > 0 {
			EmoteActive.alpha -= 0.06
		}
		if EmoteActive.alpha <= 0 {
			EmoteActive = nil
		}
	}
}

func Draw(screen *ebiten.Image) {
	if EmoteActive != nil {
		EmoteActive.Draw(screen)
	}
}
