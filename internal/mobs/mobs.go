package mobs

import (
	"fmt"
	"image/color"
	"ragnarok-buddy/internal/core"
	"ragnarok-buddy/internal/emotes"
	"ragnarok-buddy/pkg/helpers"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var (
	List       []*Mob
	rarityList []int
)

var nextSpawnTick = 1

type Mob struct {
	Emote                       *emotes.Emote
	Height, Width, X, Y, Vx, Vy float64

	name                                           string
	assetName                                      string
	speed                                          float64
	moveFuel, idleTime, idleTick                   int
	status                                         string // "idle", "walk", "attack", "damage", "die", "action", "action2", "action3", "action4"
	lifeTime, hp                                   int
	alpha                                          float32
	frameIndex, frameLength, frameTime, frameAudio int
	frameLooping                                   bool
	gifName                                        string
	invert                                         bool
	anchorMargin                                   float64
	rarity                                         int
	damageTick                                     int

	idleFrametime     int
	idleAudioFrame    int
	walkFrametime     int
	walkAudioFrame    int
	attackFrametime   int
	attackAudioFrame  int
	damageFrametime   int
	damageAudioFrame  int
	dieFrametime      int
	dieAudioFrame     int
	actionFrametime   int
	actionAudioFrame  int
	action2Frametime  int
	action2AudioFrame int
	action3Frametime  int
	action3AudioFrame int
	action4Frametime  int
	action4AudioFrame int

	// draw
	showName bool
}

func (m *Mob) Update() {
	if m.speed > 0 {
		if m.moveFuel > 0 {
			m.X += m.Vx
			if m.X+m.Width >= core.ScreenWidth {
				// walk left
				m.Vx = -m.speed
				m.invert = false
			} else if m.X <= 1 {
				// walk right
				m.Vx = m.speed
				m.invert = true
			}
			m.moveFuel--
		} else {
			if m.idleTick == 0 {
				m.setIdle()
			}

			m.idleTick++

			if m.idleTick >= m.idleTime {
				m.idleTick = 0
				m.idleTime = 0

				if helpers.Random(0, 1) == 0 {
					// walk left
					m.Vx = -m.speed
					m.invert = false
				} else {
					// walk right
					m.Vx = m.speed
					m.invert = true
				}
				m.setWalk()

				steps := helpers.Random(1, 6)
				m.moveFuel = m.frameTime*m.frameLength*steps - m.frameTime
			}
		}
	}

	m.Y = core.ScreenHeight - m.Height
	m.Y += m.Vy

	if m.frameIndex == m.frameLength-1 && !m.frameLooping {
		m.setIdle()
	}

	if core.FrameTick%m.frameTime == 0 {
		m.frameIndex = (m.frameIndex + 1) % m.frameLength

		// if core.Cfg.EffectsVolume > 0 && m.frameIndex == m.frameAudio && e.audioPath != "" {
		// 	p := core.AudioContext.NewPlayerFromBytes(assets.LoadedAudios[e.audioPath])
		// 	p.SetVolume(core.Cfg.EffectsVolume)
		// 	p.Play()
		// }
	}

	if m.alpha < 1 {
		m.alpha += 0.01
	}
}
func (m *Mob) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(m.alpha)
	if m.invert {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(m.Width, 0)
	}
	op.GeoM.Translate(m.X, m.Y-m.anchorMargin)
	screen.DrawImage(cachedGifs[m.gifName].Frames[m.frameIndex], op)

	if m.showName {
		rarity := "comum"
		if m.rarity == 10 {
			rarity = "incomum"
		} else if m.rarity == 1 {
			rarity = "RARO"
		}
		if core.Cfg.CursorAttack {
			text.Draw(screen, fmt.Sprint(m.name+" ("+rarity+") ", m.hp), core.Font, int(m.X), int(m.Y-15), color.White)
		} else {
			text.Draw(screen, m.name+" ("+rarity+")", core.Font, int(m.X), int(m.Y-15), color.White)
		}
	}
}

func (m *Mob) setStatus(status string) {
	m.status = status
	m.gifName = m.assetName + "_" + m.status
	m.frameIndex = 0
	m.frameLength = cachedGifs[m.gifName].Length
	m.Height = cachedGifs[m.gifName].Height
	m.Width = cachedGifs[m.gifName].Width
}
func (m *Mob) setIdle() {
	m.frameTime = m.idleFrametime
	m.frameAudio = m.idleAudioFrame
	m.frameLooping = true
	m.Vx = 0
	m.setStatus("idle")

	loops := helpers.Random(8, 12)
	m.idleTime = m.frameTime * m.frameLength * loops
}
func (m *Mob) setWalk() {
	if m.walkFrametime == 0 {
		return
	}
	m.frameTime = m.walkFrametime
	m.frameAudio = m.walkAudioFrame
	m.frameLooping = true
	m.setStatus("walk")
}
func (m *Mob) setAttack() {
	if m.attackFrametime == 0 {
		return
	}
	m.frameTime = m.attackFrametime
	m.frameAudio = m.attackAudioFrame
	m.frameLooping = false
	m.Vx = 0
	m.setStatus("attack")
}
func (m *Mob) setDamage() {
	if m.damageFrametime == 0 {
		return
	}
	m.frameTime = m.damageFrametime
	m.frameAudio = m.damageAudioFrame
	m.frameLooping = false
	m.Vx = 0
	m.setStatus("damage")
}
func (m *Mob) setDie() {
	if m.dieFrametime == 0 {
		return
	}
	m.frameTime = m.dieFrametime
	m.frameAudio = m.dieAudioFrame
	m.frameLooping = false
	m.Vx = 0
	m.setStatus("die")
}
func (m *Mob) setAction() {
	if m.actionFrametime == 0 {
		return
	}
	m.frameTime = m.actionFrametime
	m.frameAudio = m.actionAudioFrame
	m.frameLooping = false
	m.Vx = 0
	m.setStatus("action")
}
func (m *Mob) setAction2() {
	if m.action2Frametime == 0 {
		return
	}
	m.frameTime = m.action2Frametime
	m.frameAudio = m.action2AudioFrame
	m.frameLooping = false
	m.Vx = 0
	m.setStatus("action2")
}
func (m *Mob) setAction3() {
	if m.action3Frametime == 0 {
		return
	}
	m.frameTime = m.action3Frametime
	m.frameAudio = m.action3AudioFrame
	m.frameLooping = false
	m.Vx = 0
	m.setStatus("action3")
}
func (m *Mob) setAction4() {
	if m.action4Frametime == 0 {
		return
	}
	m.frameTime = m.action4Frametime
	m.frameAudio = m.action4AudioFrame
	m.frameLooping = false
	m.Vx = 0
	m.setStatus("action4")
}

func (m *Mob) setSpawn() {
	m.X = float64(helpers.Random(0, int(float64(core.Cfg.ScreenMonitors)*core.ScreenWidth-2*m.Width)))
	m.lifeTime = helpers.Random(core.Cfg.MobsDespawnSecondsMin, core.Cfg.MobsDespawnSecondsMin) * ebiten.TPS()
	m.setIdle()
	m.invert = helpers.Random(0, 1) == 1
}
