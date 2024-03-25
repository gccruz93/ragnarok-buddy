package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type Cfg struct {
	MainVolume          float64 `ini:"main_volume" json:"main_volume"`
	ScreenPaddingBottom int     `ini:"screen_padding_bottom" json:"screen_padding_bottom"`
	ScreenMonitors      int     `ini:"screen_monitors" json:"screen_monitors"`
	PetsMax             int     `ini:"pets_max" json:"pets_max"`
	PetsBlocked         string  `ini:"pets_blocked" json:"pets_blocked"`
	PetsSpawnSecondsMin int     `ini:"pets_spawn_seconds_min" json:"pets_spawn_seconds_min"`
	PetsSpawnSecondsMax int     `ini:"pets_spawn_seconds_max" json:"pets_spawn_seconds_max"`
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

	if c.MainVolume < 0 {
		c.MainVolume = 0
	} else if c.MainVolume > 1 {
		c.MainVolume = 1
	}

	if c.ScreenMonitors <= 0 {
		c.ScreenMonitors = 1
	} else if c.ScreenMonitors > 3 {
		c.ScreenMonitors = 3
	}

	if c.PetsMax <= 0 {
		c.PetsMax = 1
	}

	c.Save()
}

func (c *Cfg) Save() {
	_ = os.Remove("cfg.bkp.ini")
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
	c.MainVolume = 0.1
	c.ScreenPaddingBottom = 62
	c.ScreenMonitors = 1
	c.PetsMax = 6
	c.PetsBlocked = ""
	c.PetsSpawnSecondsMin = 5
	c.PetsSpawnSecondsMax = 20
}
