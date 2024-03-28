package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) UpdateGamemode0() error {
	mx, my := ebiten.CursorPosition()

	if frameCount%(nextSpawn*ebiten.TPS()) == 0 && len(mobs) < cfg.MobsSpawnMax {
		SpawnRandom(1)
	}

	initalPress := inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	if initalPress {
		log.Println("mouse left")
	}

	mobsAlive := mobs[:0]
	for _, e := range mobs {
		e.Update()

		if isMouseHover(mx, my, e.x, e.y, e.x+float64(e.width), e.y+float64(e.height)) {
			// ebiten.SetCursorMode(ebiten.CursorModeHidden)
			e.drawName = true
			if cfg.CursorCanHit {
				e.hp--
			}
		} else {
			// ebiten.SetCursorMode(ebiten.CursorModeVisible)
			e.drawName = false
			if e.hp < e.maxhp {
				e.hp++
			}
		}

		if e.hp > 0 {
			mobsAlive = append(mobsAlive, e)
		}
	}
	mobs = mobsAlive

	return nil
}

func (g *Game) UpdateGamemode1() error {
	return nil
}
