package emotes

import (
	"bytes"
	"image/gif"
	"ragnarok-buddy/internal/core"
	"ragnarok-buddy/pkg/helpers"

	"github.com/hajimehoshi/ebiten/v2"
)

func loadAsset(name string) {
	file, _ := gifs.ReadFile("gifs/" + name + ".gif")
	loadedGif, _ := gif.DecodeAll(bytes.NewReader(file))
	EmoteActive.gif = helpers.SplitAnimatedGIF(loadedGif)
}

func SpawnRandom() {
	NextSpawnTick = helpers.Random(core.Cfg.EmoteSpawnSecondsMin, core.Cfg.EmoteSpawnSecondsMax) * ebiten.TPS()

	if len(cachedConfig) == 0 {
		return
	}

	emoteConfig := cachedConfig[rarityList[helpers.Random(0, len(rarityList)-1)]]
	EmoteActive = &Emote{
		anchorMargin: emoteConfig.AnchorMargin,
		frameTime:    emoteConfig.Frametime,
		loops:        emoteConfig.Loops,
	}
	loadAsset(emoteConfig.Asset)
}
