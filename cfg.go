package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type Cfg struct {
	MainVolume float64 `ini:"main_volume" json:"main_volume"`
	Max        int     `ini:"max" json:"max"`
}

func (c *Cfg) Load() {
	ini.PrettyFormat = false
	cfg, err := ini.Load("cfg.ini")

	if err != nil {
		cfg = ini.Empty()
		c.LoadDefaults()
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
}

func (c *Cfg) Save() {
	_ = os.Remove("cfg.bkp.ini")
	_ = os.Rename("cfg.ini", "cfg.bkp.ini")
	cfg := ini.Empty()
	_ = cfg.ReflectFrom(c)
	err := cfg.SaveTo("cfg.ini")
	if err != nil {
		fmt.Printf("Fail to save file: %v", err)
		os.Exit(1)
	}
}

func (c *Cfg) LoadDefaults() {
	c.MainVolume = 0.5
	c.Max = 10
}
