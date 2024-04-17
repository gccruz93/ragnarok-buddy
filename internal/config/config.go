package config

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type Config struct {
	Effects       bool    `ini:"effects" json:"effects"`
	EffectsVolume float64 `ini:"effects_volume" json:"effects_volume"`
	Music         bool    `ini:"music" json:"music"`
	MusicVolume   float64 `ini:"music_volume" json:"music_volume"`

	SkipTaskbar         bool `ini:"skip_taskbar" json:"skip_taskbar"`
	ScreenPaddingBottom int  `ini:"screen_padding_bottom" json:"screen_padding_bottom"`
	ScreenMonitors      int  `ini:"screen_monitors" json:"screen_monitors"`
	CursorAttack        bool `ini:"cursor_attack" json:"cursor_attack"`

	MobsSpawnCycle        bool   `ini:"mobs_cycle" json:"mobs_cycle"`
	MobsSpawnTotal        int    `ini:"mobs_spawn_max" json:"mobs_spawn_max"`
	MobsSpawnSecondsMin   int    `ini:"mobs_spawn_seconds_min" json:"mobs_spawn_seconds_min"`
	MobsSpawnSecondsMax   int    `ini:"mobs_spawn_seconds_max" json:"mobs_spawn_seconds_max"`
	MobsDespawnSecondsMin int    `ini:"mobs_despawn_seconds_min" json:"mobs_despawn_seconds_min"`
	MobsDespawnSecondsMax int    `ini:"mobs_despawn_seconds_max" json:"mobs_despawn_seconds_max"`
	MobsAllowed           string `ini:"mobs_allowed" json:"mobs_allowed"`
	MobsBlocked           string `ini:"mobs_blocked" json:"mobs_blocked"`

	EmoteSpawnSecondsMin int `ini:"emote_spawn_seconds_min" json:"emote_spawn_seconds_min"`
	EmoteSpawnSecondsMax int `ini:"emote_spawn_seconds_max" json:"emote_spawn_seconds_max"`

	Map      string `ini:"map" json:"map"`
	MapCycle bool   `ini:"map_cycle" json:"map_cycle"`

	Debug bool `ini:"debug" json:"debug"`
}

func (c *Config) Load() {
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

	c.ScreenMonitors = c.IntPositive(c.ScreenMonitors, 1)
	c.MobsSpawnTotal = c.IntPositive(c.MobsSpawnTotal, 6)

	c.Save()
}

func (c *Config) Save() {
	cfg := ini.Empty()
	_ = cfg.ReflectFrom(c)
	err := cfg.SaveTo("cfg.ini")
	if err != nil {
		fmt.Printf("Fail to save file: %v", err)
		os.Exit(1)
	}
}

func (c *Config) LoadDefaults() {
	c.Effects = false
	c.EffectsVolume = 0.1
	c.Music = false
	c.MusicVolume = 0.2

	c.ScreenPaddingBottom = 62
	c.ScreenMonitors = 1
	c.CursorAttack = false

	c.MobsSpawnCycle = true
	c.MobsSpawnTotal = 4
	c.MobsSpawnSecondsMin = 3
	c.MobsSpawnSecondsMax = 10
	c.MobsDespawnSecondsMin = 60
	c.MobsDespawnSecondsMax = 120
	c.EmoteSpawnSecondsMin = 45
	c.EmoteSpawnSecondsMax = 120

	c.MapCycle = false
}

func (c *Config) IntRange(val, min, max, dfault int) int {
	if val < min {
		return dfault
	} else if val > max {
		return dfault
	}
	return val
}
func (c *Config) IntPositive(val, dfault int) int {
	if val <= 0 {
		return dfault
	}
	return val
}
func (c *Config) FloatRange(val, min, max, dfault float64) float64 {
	if val < min {
		return dfault
	} else if val > max {
		return dfault
	}
	return val
}
