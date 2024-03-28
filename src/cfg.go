package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type Cfg struct {
	EffectsMuted          bool    `ini:"effects_muted" json:"effects_muted"`
	EffectsVolume         float64 `ini:"effects_volume" json:"effects_volume"`
	MusicMuted            bool    `ini:"music_muted" json:"music_muted"`
	MusicVolume           float64 `ini:"music_volume" json:"music_volume"`
	Gamemode              int     `ini:"gamemode" json:"gamemode"`
	SkipTaskbar           bool    `ini:"skip_taskbar" json:"skip_taskbar"`
	ScreenPaddingBottom   int     `ini:"screen_padding_bottom" json:"screen_padding_bottom"`
	ScreenMonitors        int     `ini:"screen_monitors" json:"screen_monitors"`
	CursorCanHit          bool    `ini:"cursor_can_hit" json:"cursor_can_hit"`
	MobsSpawn             bool    `ini:"mobs_spawn" json:"mobs_spawn"`
	MobsDespawn           bool    `ini:"mobs_despawn" json:"mobs_despawn"`
	MobsSpawnMax          int     `ini:"mobs_spawn_max" json:"mobs_spawn_max"`
	MobsSpawnSecondsMin   int     `ini:"mobs_spawn_seconds_min" json:"mobs_spawn_seconds_min"`
	MobsSpawnSecondsMax   int     `ini:"mobs_spawn_seconds_max" json:"mobs_spawn_seconds_max"`
	MobsDespawnSecondsMin int     `ini:"mobs_despawn_seconds_min" json:"mobs_despawn_seconds_min"`
	MobsDespawnSecondsMax int     `ini:"mobs_despawn_seconds_max" json:"mobs_despawn_seconds_max"`
	Debug                 bool    `ini:"debug" json:"debug"`

	// gamemode = 0
	MobsAllowed string `ini:"mobs_allowed" json:"mobs_allowed"`
	MobsBlocked string `ini:"mobs_blocked" json:"mobs_blocked"`

	// gamemode = 1
	Map      string `ini:"map" json:"map"`
	MapCycle bool   `ini:"map_cycle" json:"map_cycle"`
}

func (c *Cfg) Load() {
	ini.PrettyFormat = false
	c.LoadDefaults()

	cfg, err := ini.Load("cfg.ini")
	if err != nil {
		cfg = ini.Empty()
		_ = cfg.ReflectFrom(c)
		_ = cfg.SaveTo("cfg.ini")
	}

	err = cfg.MapTo(&c)
	if err != nil {
		fmt.Printf("Fail to map file: %v", err)
		os.Exit(1)
	}

	c.MusicVolume = c.FloatRange(c.MusicVolume, 0, 1, 0.2)
	c.EffectsVolume = c.FloatRange(c.EffectsVolume, 0, 1, 0.1)
	c.Gamemode = c.IntRange(c.Gamemode, 0, 1, 0)
	c.ScreenMonitors = c.IntPositive(c.ScreenMonitors, 1)
	c.MobsSpawnMax = c.IntPositive(c.MobsSpawnMax, 6)
	c.MobsSpawnSecondsMin = c.IntPositive(c.MobsSpawnSecondsMin, 5)
	c.MobsSpawnSecondsMax = c.IntPositive(c.MobsSpawnSecondsMax, 20)

	c.Save()
}

func (c *Cfg) Save() {
	// _ = os.Remove("cfg.bkp.ini")
	// _ = os.Rename("cfg.ini", "cfg.bkp.ini")
	cfg := ini.Empty()
	_ = cfg.ReflectFrom(c)
	err := cfg.SaveTo("cfg.ini")
	if err != nil {
		fmt.Printf("Fail to save file: %v", err)
		os.Exit(1)
	}
}

func (c *Cfg) LoadDefaults() {
	c.MusicVolume = 0.2
	c.EffectsVolume = 0.1
	c.ScreenPaddingBottom = 62
	c.ScreenMonitors = 1
	c.CursorCanHit = true
	c.MobsSpawn = true
	c.MobsDespawn = false
	c.MobsSpawnMax = 6
	c.MobsSpawnSecondsMin = 5
	c.MobsSpawnSecondsMax = 20
	c.MobsDespawnSecondsMin = 10
	c.MobsDespawnSecondsMax = 20

	// gamemode = 1
	c.MapCycle = true
}

func (c *Cfg) IntRange(val, min, max, dfault int) int {
	if val < min {
		return dfault
	} else if val > max {
		return dfault
	}
	return val
}
func (c *Cfg) IntPositive(val, dfault int) int {
	if val <= 0 {
		return dfault
	}
	return val
}
func (c *Cfg) FloatRange(val, min, max, dfault float64) float64 {
	if val < min {
		return dfault
	} else if val > max {
		return dfault
	}
	return val
}
