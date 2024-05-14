package mobs

import (
	"ragnarok-buddy/internal/core"
	"ragnarok-buddy/pkg/helpers"

	"github.com/hajimehoshi/ebiten/v2"
)

func Update() {
	eventSpawnRandom()

	mobsAlive := List[:0]
	for _, e := range List {
		e.Update()

		if helpers.IsMouseHover(core.Mx, core.My, e.X, e.Y, e.X+e.Width, e.Y+e.Height) {
			e.isHovering = true

			if core.Cfg.CursorAttack {
				e.damageTick++
				if e.damageTick%ebiten.TPS() == 0 {
					e.hp -= 35
					e.setDamage()
				}
			}
		} else {
			e.isHovering = false
			e.damageTick = 0
			if e.hp < 100 {
				e.hp++
			}
		}

		if core.Cfg.MobsSpawnCycle {
			e.lifeTime--
		}

		if e.Emote != nil {
			e.Emote.X = e.X
			e.Emote.Y = e.Y
		}

		if e.lifeTime > 0 && e.hp > 0 {
			mobsAlive = append(mobsAlive, e)
		}
	}
	List = mobsAlive
}

func Draw(screen *ebiten.Image) {
	for _, e := range List {
		e.Draw(screen)
	}
}

func eventSpawnRandom() {
	if core.FrameTick%nextSpawnTick == 0 && len(List) < core.Cfg.MobsSpawnTotal {
		SpawnRandom(1)
	}
}
