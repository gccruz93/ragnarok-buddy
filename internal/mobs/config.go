package mobs

import (
	"embed"
	"encoding/json"
	"ragnarok-buddy/pkg/helpers"
)

type MobConfig struct {
	Name              string  `json:"name"`
	Asset             string  `json:"asset"`
	Speed             float64 `json:"speed"`
	Rarity            int     `json:"rarity"`
	AnchorMargin      float64 `json:"anchor_margin"`
	ActionFrametime   int     `json:"action_frametime"`
	ActionAudioFrame  int     `json:"action_audioframe"`
	Action2Frametime  int     `json:"action2_frametime"`
	Action2AudioFrame int     `json:"action2_audioframe"`
	Action3Frametime  int     `json:"action3_frametime"`
	Action3AudioFrame int     `json:"action3_audioframe"`
	Action4Frametime  int     `json:"action4_frametime"`
	Action4AudioFrame int     `json:"action4_audioframe"`
	IdleFrametime     int     `json:"idle_frametime"`
	IdleAudioFrame    int     `json:"idle_audioframe"`
	WalkFrametime     int     `json:"walk_frametime"`
	WalkAudioFrame    int     `json:"walk_audioframe"`
	AttackFrametime   int     `json:"attack_frametime"`
	AttackAudioFrame  int     `json:"attack_audioframe"`
	DamageFrametime   int     `json:"damage_frametime"`
	DamageAudioFrame  int     `json:"damage_audioframe"`
	DieFrametime      int     `json:"die_frametime"`
	DieAudioFrame     int     `json:"die_audioframe"`
}

var (
	//go:embed gifs/*
	gifs embed.FS

	//go:embed mobs.json
	jsonFile []byte

	cachedConfig []*MobConfig
)

func LoadConfig() {
	List = nil
	rarityList = nil
	cachedConfig = nil
	cachedGifs = nil
	cachedGifs = make(map[string]*helpers.CustomGif)
	nextSpawnTick = 1

	var mobs []*MobConfig
	if err := json.Unmarshal(jsonFile, &mobs); err != nil {
		panic(err)
	}
	for i, mob := range mobs {
		cachedConfig = append(cachedConfig, mob)
		for j := 0; j <= mob.Rarity; j++ {
			rarityList = append(rarityList, i)
		}
	}
}
