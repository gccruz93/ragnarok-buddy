package emotes

import (
	"embed"
	"encoding/json"
)

type EmoteConfig struct {
	Asset        string  `json:"asset"`
	Frametime    int     `json:"frametime"`
	Rarity       int     `json:"rarity"`
	AnchorMargin float64 `json:"anchor_margin"`
	Loops        int     `json:"loops"`
}

var (
	//go:embed gifs/*
	gifs embed.FS

	//go:embed emotes.json
	jsonFile []byte

	cachedConfig []*EmoteConfig
)

func LoadConfig() {
	rarityList = nil
	cachedConfig = nil
	NextSpawnTick = 1

	var emotes []*EmoteConfig
	if err := json.Unmarshal(jsonFile, &emotes); err != nil {
		panic(err)
	}
	for i, mob := range emotes {
		cachedConfig = append(cachedConfig, mob)
		for j := 0; j <= mob.Rarity; j++ {
			rarityList = append(rarityList, i)
		}
	}
}
