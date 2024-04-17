package main

import (
	"fmt"
	"ragnarok-buddy/assets"
	"ragnarok-buddy/internal/core"
	"ragnarok-buddy/internal/emotes"
	"ragnarok-buddy/internal/mobs"

	"github.com/energye/systray"
)

func onReady() {
	systray.SetIcon(assets.Icontray)
	systray.SetTitle(core.Title)
	systray.SetTooltip(core.Title)
	systray.SetOnRClick(func(menu systray.IMenu) {
		menu.ShowMenu()
	})

	systray.AddMenuItem("v0.2.0 - @twpax", "v0.2.0 - @twpax").Disable()

	mQuit := systray.AddMenuItem("Abandonar os bixin :(", "Fechar")
	mQuit.SetIcon(assets.IconAlWarp)
	mQuit.Click(func() {
		core.Cfg.Save()
		core.IsRunning = false
		systray.Quit()
	})

	systray.AddSeparator()

	mTaskbar := systray.AddMenuItem("Esconder da barra de tarefas", "Esconder da barra de tarefas")
	mTaskbar.SetIcon(assets.IconTfHiding)
	mTaskbar.Click(func() {
		core.Cfg.SkipTaskbar = !core.Cfg.SkipTaskbar
		if core.Cfg.SkipTaskbar {
			mTaskbar.Check()
		} else {
			mTaskbar.Uncheck()
		}
	})
	if core.Cfg.SkipTaskbar {
		mTaskbar.Check()
	}

	mReload := systray.AddMenuItem("Recarregar", "Recarregar")
	mReload.SetIcon(assets.IconMoExtremityFirst)
	mReload.Click(func() {
		core.Cfg.Load()
		mobs.LoadConfig()
		emotes.LoadConfig()
	})

	mCursorAttack := systray.AddMenuItem("Dano com cursor", "Dano com cursor")
	mCursorAttack.SetIcon(assets.IconCursorAttack)
	mCursorAttack.Click(func() {
		core.Cfg.CursorAttack = !core.Cfg.CursorAttack
		if core.Cfg.CursorAttack {
			mCursorAttack.Check()
		} else {
			mCursorAttack.Uncheck()
		}
	})
	if core.Cfg.CursorAttack {
		mCursorAttack.Check()
	}

	systray.AddSeparator()

	/**
	* ========== SOUNDS ==========
	 */
	mSounds := systray.AddMenuItem("Sons", "Sons")
	mSounds.SetIcon(assets.IconDcHumming)
	mSounds_Music := mSounds.AddSubMenuItem("Música", "Música")
	mSounds_Music.Click(func() {
		core.Cfg.Music = !core.Cfg.Music
		if core.Cfg.Music {
			mSounds_Music.Check()
		} else {
			mSounds_Music.Uncheck()
		}
	})
	if core.Cfg.Music {
		mSounds_Music.Check()
	}
	mSounds_Effects := mSounds.AddSubMenuItem("Efeitos", "Efeitos")
	mSounds_Effects.Click(func() {
		core.Cfg.Effects = !core.Cfg.Effects
		if core.Cfg.Effects {
			mSounds_Effects.Check()
		} else {
			mSounds_Effects.Uncheck()
		}
	})
	if core.Cfg.Effects {
		mSounds_Effects.Check()
	}

	systray.AddSeparator()

	/**
	* ========== MAPS ==========
	 */
	mMaps := systray.AddMenuItem("Mapas", "Mapas")
	mMapsCycle := mMaps.AddSubMenuItem("Ciclo", "")
	mMapsCycle.Click(func() {
		core.Cfg.MapCycle = !core.Cfg.MapCycle
		if core.Cfg.MapCycle {
			mMapsCycle.Check()
		} else {
			mMapsCycle.Uncheck()
		}
	})
	if core.Cfg.MapCycle {
		mMapsCycle.Check()
	}

	mMapsDiv1 := mMaps.AddSubMenuItem("", "")
	mMapsDiv1.Disable()

	mMapGeffenFields := mMaps.AddSubMenuItem("Arredores de Geffen", "")
	mMapGeffenFields.Click(func() {
		core.Cfg.Map = "geffen_fields"
		mMapGeffenFields.Check()
	})
	if core.Cfg.Map == "geffen_fields" {
		mMapGeffenFields.Check()
	}

	/**
	* ========== MOBS ==========
	 */
	mMobs := systray.AddMenuItem("Mobs", "Mobs")
	mMobsSpawnTotal := mMobs.AddSubMenuItem(fmt.Sprintf("Total: %d", core.Cfg.MobsSpawnTotal), "")
	mMobsSpawnTotal.Disable()

	mMobs.AddSubMenuItem("Total++", "").Click(func() {
		core.Cfg.MobsSpawnTotal++
		mMobsSpawnTotal.SetTitle(fmt.Sprintf("Total: %d", core.Cfg.MobsSpawnTotal))
	})

	mMobs.AddSubMenuItem("Total--", "").Click(func() {
		core.Cfg.MobsSpawnTotal--
		mMobsSpawnTotal.SetTitle(fmt.Sprintf("Total: %d", core.Cfg.MobsSpawnTotal))
	})

	mMobsSpawnCycle := mMobs.AddSubMenuItem("Ciclo", "Ciclo")
	mMobsSpawnCycle.Click(func() {
		core.Cfg.MobsSpawnCycle = !core.Cfg.MobsSpawnCycle
		if core.Cfg.MobsSpawnCycle {
			mMobsSpawnCycle.Check()
		} else {
			mMobsSpawnCycle.Uncheck()
		}
	})
	if core.Cfg.MobsSpawnCycle {
		mMobsSpawnCycle.Check()
	}

	mMobsSpawn := systray.AddMenuItem("Usar galho seco", "Usar galho seco")
	mMobsSpawn.SetIcon(assets.IconDeadBranch)
	mMobsSpawn.Click(func() {
		mobs.SpawnRandom(1)
	})
}

func onExit() {}
