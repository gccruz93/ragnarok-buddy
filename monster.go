package main

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Monster struct {
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

func (m *Monster) Update() {
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
func (m *Monster) Draw(screen *ebiten.Image) {
	m.Entity.Draw(screen)

	if m.drawName {
		text.Draw(screen, fmt.Sprint(m.name+" ", m.hp), mplusNormalFont, int(m.x), int(m.y-float64(m.height)/2), color.White)
	}
}
func (m *Monster) SetIdle() {
	if m.vx > 0 {
		m.SetIdleRight()
	} else {
		m.SetIdleLeft()
	}
}
func (m *Monster) SetIdleLeft() {
	m.SetGif(m.idleLeftGif, m.idleFrametime)
	if m.idleAudio != "" {
		m.SetAudio(m.idleAudio, m.idleAudioFrame)
	}
	m.vx = 0
}
func (m *Monster) SetIdleRight() {
	m.SetGif(m.idleRightGif, m.idleFrametime)
	if m.idleAudio != "" {
		m.SetAudio(m.idleAudio, m.idleAudioFrame)
	}
	m.vx = 0
}
func (m *Monster) SetWalkLeft() {
	m.SetGif(m.walkLeftGif, m.walkFrametime)
	if m.walkAudio != "" {
		m.SetAudio(m.walkAudio, m.walkAudioFrame)
	}
	m.vx = -m.speed
}
func (m *Monster) SetWalkRight() {
	m.SetGif(m.walkRightGif, m.walkFrametime)
	if m.walkAudio != "" {
		m.SetAudio(m.walkAudio, m.walkAudioFrame)
	}
	m.vx = m.speed
}
func (m *Monster) SetSpawn() {
	m.x = float64(random(0, screenWidth-m.width))
}
func (m *Monster) SetGifs(name string) {
	m.idleLeftGif = "monster/" + name + "_idle_left.gif"
	m.idleRightGif = "monster/" + name + "_idle_right.gif"
	m.walkLeftGif = "monster/" + name + "_walk_left.gif"
	m.walkRightGif = "monster/" + name + "_walk_right.gif"
}
func (m *Monster) SetIdleFrametime(frametime int) {
	m.idleFrametime = frametime
}
func (m *Monster) SetWalkFrametime(frametime int) {
	m.walkFrametime = frametime
}
func (m *Monster) SetIdleAudio(name string, frame int) {
	m.idleAudio = "sound/" + name + "_idle.wav"
	m.idleAudioFrame = frame
}
func (m *Monster) SetWalkAudio(name string, frame int) {
	m.walkAudio = "sound/" + name + "_move.wav"
	m.walkAudioFrame = frame
}

func SpawnRandom(n int) {
	nextSpawn = random(cfg.PetsSpawnSecondsMin, cfg.PetsSpawnSecondsMax)

	entry := []string{"angeling", "baphometjr", "ghostring", "kobold_axe", "kobold_hammer", "kobold_mace", "lunatic", "poring", "smokie", "spore"}

	if cfg.PetsBlocked != "" {
		canSpawn := false
		for _, pet := range entry {
			if !strings.Contains(cfg.PetsBlocked, pet) {
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
		if cfg.PetsBlocked != "" && strings.Contains(cfg.PetsBlocked, pet) {
			continue
		}
		switch pet {
		case "angeling":
			pets = append(pets, NewAngeling())
		case "baphometjr":
			pets = append(pets, NewBaphometjr())
		case "ghostring":
			pets = append(pets, NewGhostring())
		case "kobold_axe":
			pets = append(pets, NewKoboldAxe())
		case "kobold_hammer":
			pets = append(pets, NewKoboldHammer())
		case "kobold_mace":
			pets = append(pets, NewKoboldMace())
		case "lunatic":
			pets = append(pets, NewLunatic())
		case "poring":
			pets = append(pets, NewPoring())
		case "smokie":
			pets = append(pets, NewSmokie())
		case "spore":
			pets = append(pets, NewSpore())
		default:
		}
		n--
	}
}

func SpawnAll() {
	pets = append(pets, NewAngeling())
	pets = append(pets, NewBaphometjr())
	pets = append(pets, NewGhostring())
	pets = append(pets, NewKoboldAxe())
	pets = append(pets, NewKoboldHammer())
	pets = append(pets, NewKoboldMace())
	pets = append(pets, NewLunatic())
	pets = append(pets, NewPoring())
	pets = append(pets, NewSmokie())
	pets = append(pets, NewSpore())
}

func NewPoring() *Monster {
	m := &Monster{
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

func NewLunatic() *Monster {
	m := &Monster{
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

func NewAngeling() *Monster {
	m := &Monster{
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

func NewBaphometjr() *Monster {
	m := &Monster{
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

func NewGhostring() *Monster {
	m := &Monster{
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

func NewKoboldAxe() *Monster {
	m := &Monster{
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

func NewKoboldHammer() *Monster {
	m := &Monster{
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

func NewKoboldMace() *Monster {
	m := &Monster{
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

func NewSmokie() *Monster {
	m := &Monster{
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

func NewSpore() *Monster {
	m := &Monster{
		name:  "Spore",
		speed: 1.5,
		maxhp: 100,
		hp:    100,
	}
	m.SetGifs("spore")
	m.SetIdleFrametime(4)
	m.SetWalkFrametime(2)
	m.SetSpawn()
	m.SetIdle()
	return m
}
