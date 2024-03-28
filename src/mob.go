package main

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Mob struct {
	Entity
	name                          string
	speed                         float64
	moveFuel, idleTime, idleCount int
	maxhp, hp                     int

	// draw
	drawName bool

	// idle
	idleFrametime  int
	idleLeftGif    string
	idleRightGif   string
	idleAudio      string
	idleAudioFrame int
	// walk
	walkFrametime  int
	walkLeftGif    string
	walkRightGif   string
	walkAudio      string
	walkAudioFrame int
}

func (m *Mob) Update() {
	if m.moveFuel > 0 {
		m.x += m.vx
		if m.x+float64(m.width) >= float64(screenWidth) {
			m.SetWalkLeft()
		} else if m.x <= 1 {
			m.SetWalkRight()
		}
		m.moveFuel--
	} else {
		if m.idleCount == 0 {
			m.SetIdle()
			m.idleTime = random(200, 400)
		}

		m.idleCount++

		if m.idleCount >= m.idleTime {
			m.idleCount = 0
			m.idleTime = 0

			if random(0, 1) == 0 {
				m.SetWalkLeft()
			} else {
				m.SetWalkRight()
			}

			steps := random(1, 8)
			m.moveFuel = (m.frameSpeed * m.frameLength * steps) - steps
		}
	}

	m.Entity.Update()
}
func (m *Mob) Draw(screen *ebiten.Image) {
	m.Entity.Draw(screen)

	if m.drawName {
		text.Draw(screen, fmt.Sprint(m.name+" ", m.hp), mplusNormalFont, int(m.x), int(m.y-float64(m.height)/2), color.White)
	}
}
func (m *Mob) SetIdle() {
	if m.vx > 0 {
		m.SetIdleRight()
	} else {
		m.SetIdleLeft()
	}
}
func (m *Mob) SetIdleLeft() {
	m.SetGif(m.idleLeftGif, m.idleFrametime)
	if m.idleAudio != "" {
		m.SetAudio(m.idleAudio, m.idleAudioFrame)
	}
	m.vx = 0
}
func (m *Mob) SetIdleRight() {
	m.SetGif(m.idleRightGif, m.idleFrametime)
	if m.idleAudio != "" {
		m.SetAudio(m.idleAudio, m.idleAudioFrame)
	}
	m.vx = 0
}
func (m *Mob) SetWalkLeft() {
	m.SetGif(m.walkLeftGif, m.walkFrametime)
	if m.walkAudio != "" {
		m.SetAudio(m.walkAudio, m.walkAudioFrame)
	}
	m.vx = -m.speed
}
func (m *Mob) SetWalkRight() {
	m.SetGif(m.walkRightGif, m.walkFrametime)
	if m.walkAudio != "" {
		m.SetAudio(m.walkAudio, m.walkAudioFrame)
	}
	m.vx = m.speed
}
func (m *Mob) SetSpawn() {
	m.x = float64(random(0, screenWidth-m.width))
}
func (m *Mob) SetGifs(name string) {
	m.idleLeftGif = "assets/mob/" + name + "_idle_left.gif"
	m.idleRightGif = "assets/mob/" + name + "_idle_right.gif"
	m.walkLeftGif = "assets/mob/" + name + "_walk_left.gif"
	m.walkRightGif = "assets/mob/" + name + "_walk_right.gif"
}
func (m *Mob) SetIdleFrametime(frametime int) {
	m.idleFrametime = frametime
}
func (m *Mob) SetWalkFrametime(frametime int) {
	m.walkFrametime = frametime
}
func (m *Mob) SetIdleAudio(name string, frame int) {
	m.idleAudio = "assets/sound/effect/" + name + "_idle.wav"
	m.idleAudioFrame = frame
}
func (m *Mob) SetWalkAudio(name string, frame int) {
	m.walkAudio = "assets/sound/effect/" + name + "_move.wav"
	m.walkAudioFrame = frame
}

func SpawnRandom(n int) {
	nextSpawn = random(cfg.MobsSpawnSecondsMin, cfg.MobsSpawnSecondsMax)

	entry := []string{"angeling", "baphometjr", "ghostring", "kobold_axe", "kobold_hammer", "kobold_mace", "lunatic", "poring", "smokie", "spore"}

	if cfg.MobsBlocked != "" {
		canSpawn := false
		for _, pet := range entry {
			if !strings.Contains(cfg.MobsBlocked, pet) {
				canSpawn = true
				break
			}
		}
		if !canSpawn {
			return
		}
	}

	for n > 0 {
		pet := entry[random(0, len(entry)-1)]
		if cfg.MobsBlocked != "" && strings.Contains(cfg.MobsBlocked, pet) {
			continue
		}
		switch pet {
		case "angeling":
			mobs = append(mobs, NewAngeling())
		case "baphometjr":
			mobs = append(mobs, NewBaphometjr())
		case "ghostring":
			mobs = append(mobs, NewGhostring())
		case "kobold_axe":
			mobs = append(mobs, NewKoboldAxe())
		case "kobold_hammer":
			mobs = append(mobs, NewKoboldHammer())
		case "kobold_mace":
			mobs = append(mobs, NewKoboldMace())
		case "lunatic":
			mobs = append(mobs, NewLunatic())
		case "poring":
			mobs = append(mobs, NewPoring())
		case "smokie":
			mobs = append(mobs, NewSmokie())
		case "spore":
			mobs = append(mobs, NewSpore())
		default:
		}
		n--
	}
}

func SpawnAll() {
	mobs = append(mobs, NewAngeling())
	mobs = append(mobs, NewBaphometjr())
	mobs = append(mobs, NewGhostring())
	mobs = append(mobs, NewKoboldAxe())
	mobs = append(mobs, NewKoboldHammer())
	mobs = append(mobs, NewKoboldMace())
	mobs = append(mobs, NewLunatic())
	mobs = append(mobs, NewPoring())
	mobs = append(mobs, NewSmokie())
	mobs = append(mobs, NewSpore())
}

func NewPoring() *Mob {
	m := &Mob{
		name:  "Poring",
		speed: 1.3,
		maxhp: 100,
		hp:    100,
	}
	m.SetGifs("poring")
	m.SetIdleFrametime(6)
	m.SetWalkFrametime(3)
	m.SetWalkAudio("poring", 7)
	m.SetSpawn()
	m.SetIdle()
	return m
}

func NewLunatic() *Mob {
	m := &Mob{
		name:  "Lunatic",
		speed: 1.9,
		maxhp: 100,
		hp:    100,
	}
	m.SetGifs("lunatic")
	m.SetIdleFrametime(3)
	m.SetWalkFrametime(2)
	m.SetSpawn()
	m.SetIdle()
	return m
}

func NewAngeling() *Mob {
	m := &Mob{
		name:  "Angeling",
		speed: 2.5,
		maxhp: 100,
		hp:    100,
	}
	m.SetGifs("angeling")
	m.SetIdleFrametime(4)
	m.SetWalkFrametime(2)
	m.SetWalkAudio("poring", 7)
	m.SetSpawn()
	m.SetIdle()
	return m
}

func NewBaphometjr() *Mob {
	m := &Mob{
		name:  "Baphomet Jr.",
		speed: 4,
		maxhp: 100,
		hp:    100,
	}
	m.SetGifs("baphometjr")
	m.SetIdleFrametime(5)
	m.SetWalkFrametime(3)
	m.SetIdleAudio("baphometjr", 3)
	m.SetSpawn()
	m.SetIdle()
	return m
}

func NewGhostring() *Mob {
	m := &Mob{
		name:  "Ghostring",
		speed: 2,
		maxhp: 100,
		hp:    100,
	}
	m.SetGifs("ghostring")
	m.SetIdleFrametime(5)
	m.SetWalkFrametime(5)
	m.SetSpawn()
	m.SetIdle()
	return m
}

func NewKoboldAxe() *Mob {
	m := &Mob{
		name:  "Kobold",
		speed: 1.3,
		maxhp: 100,
		hp:    100,
	}
	m.SetGifs("kobold_axe")
	m.SetIdleFrametime(5)
	m.SetWalkFrametime(3)
	m.SetSpawn()
	m.SetIdle()
	return m
}

func NewKoboldHammer() *Mob {
	m := &Mob{
		name:  "Kobold",
		speed: 1.3,
		maxhp: 100,
		hp:    100,
	}
	m.SetGifs("kobold_hammer")
	m.SetIdleFrametime(5)
	m.SetWalkFrametime(3)
	m.SetSpawn()
	m.SetIdle()
	return m
}

func NewKoboldMace() *Mob {
	m := &Mob{
		name:  "Kobold",
		speed: 1.3,
		maxhp: 100,
		hp:    100,
	}
	m.SetGifs("kobold_mace")
	m.SetIdleFrametime(5)
	m.SetWalkFrametime(3)
	m.SetSpawn()
	m.SetIdle()
	return m
}

func NewSmokie() *Mob {
	m := &Mob{
		name:  "Smokie",
		speed: 1.5,
		maxhp: 100,
		hp:    100,
	}
	m.SetGifs("smokie")
	m.SetIdleFrametime(5)
	m.SetWalkFrametime(3)
	m.SetSpawn()
	m.SetIdle()
	return m
}

func NewSpore() *Mob {
	m := &Mob{
		name:  "Spore",
		speed: 1.5,
		maxhp: 100,
		hp:    100,
	}
	m.SetGifs("spore")
	m.SetIdleFrametime(4)
	m.SetWalkFrametime(2)
	// m.SetWalkAudio("spore", 1)
	m.SetSpawn()
	m.SetIdle()
	return m
}
