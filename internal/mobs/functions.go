package mobs

import (
	"bytes"
	"image/gif"
	"ragnarok-buddy/internal/core"
	"ragnarok-buddy/pkg/helpers"

	"github.com/hajimehoshi/ebiten/v2"
)

var cachedGifs map[string]*helpers.CustomGif

func loadAsset(name string) {
	if _, ok := cachedGifs[name]; ok {
		return
	}
	file, _ := gifs.ReadFile("gifs/" + name + ".gif")
	loadedGif, _ := gif.DecodeAll(bytes.NewReader(file))
	cachedGifs[name] = helpers.SplitAnimatedGIF(loadedGif)
}

func SpawnRandom(n int) {
	nextSpawnTick = helpers.Random(core.Cfg.MobsSpawnSecondsMin, core.Cfg.MobsSpawnSecondsMax) * ebiten.TPS()

	if len(cachedConfig) == 0 {
		return
	}

	for n > 0 {
		mobConfig := cachedConfig[rarityList[helpers.Random(0, len(rarityList)-1)]]
		mob := &Mob{
			name:              mobConfig.Name,
			assetName:         mobConfig.Asset,
			speed:             mobConfig.Speed,
			anchorMargin:      mobConfig.AnchorMargin,
			actionFrametime:   mobConfig.ActionFrametime,
			actionAudioFrame:  mobConfig.ActionAudioFrame,
			action2Frametime:  mobConfig.Action2Frametime,
			action2AudioFrame: mobConfig.Action2Frametime,
			action3Frametime:  mobConfig.Action3Frametime,
			action3AudioFrame: mobConfig.Action3Frametime,
			action4Frametime:  mobConfig.Action4Frametime,
			action4AudioFrame: mobConfig.Action4Frametime,
			idleFrametime:     mobConfig.IdleFrametime,
			idleAudioFrame:    mobConfig.IdleAudioFrame,
			walkFrametime:     mobConfig.WalkFrametime,
			walkAudioFrame:    mobConfig.WalkAudioFrame,
			attackFrametime:   mobConfig.AttackFrametime,
			attackAudioFrame:  mobConfig.AttackAudioFrame,
			damageFrametime:   mobConfig.DamageFrametime,
			damageAudioFrame:  mobConfig.DamageAudioFrame,
			dieFrametime:      mobConfig.DieFrametime,
			dieAudioFrame:     mobConfig.DieAudioFrame,
			rarity:            mobConfig.Rarity,
			hp:                100,
		}
		mob.anchorMargin = mobConfig.AnchorMargin

		if mobConfig.ActionFrametime > 0 {
			loadAsset(mobConfig.Asset + "_action")
		}
		if mobConfig.Action2Frametime > 0 {
			loadAsset(mobConfig.Asset + "_action2")
		}
		if mobConfig.Action3Frametime > 0 {
			loadAsset(mobConfig.Asset + "_action3")
		}
		if mobConfig.Action4Frametime > 0 {
			loadAsset(mobConfig.Asset + "_action4")
		}
		if mobConfig.IdleFrametime > 0 {
			loadAsset(mobConfig.Asset + "_idle")
		}
		if mobConfig.WalkFrametime > 0 {
			loadAsset(mobConfig.Asset + "_walk")
		}
		if mobConfig.AttackFrametime > 0 {
			loadAsset(mobConfig.Asset + "_attack")
		}
		if mobConfig.DamageFrametime > 0 {
			loadAsset(mobConfig.Asset + "_damage")
		}
		if mobConfig.DieFrametime > 0 {
			loadAsset(mobConfig.Asset + "_die")
		}

		mob.setStatus("idle")
		mob.setSpawn()

		List = append(List, mob)
		n--
	}
}
