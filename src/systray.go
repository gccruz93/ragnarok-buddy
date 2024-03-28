package main

import (
	"fmt"
	"log"

	"github.com/energye/systray"
)

func onReady() {
	iconBytes, err := assets.ReadFile("assets/favicon.ico")
	if err != nil {
		log.Fatal(err)
	}
	systray.SetIcon(iconBytes)
	systray.SetTitle(title)
	systray.SetTooltip(title)
	systray.SetOnRClick(func(menu systray.IMenu) {
		menu.ShowMenu()
	})
	// desabilitado pq trava a execução do app
	// systray.SetOnDClick(func(menu systray.IMenu) {
	// 	if ebiten.IsWindowMaximized() {
	// 		ebiten.MinimizeWindow()
	// 		} else {
	// 		ebiten.RestoreWindow()
	// 	}
	// })

	mGamemode := systray.AddMenuItem("Gamemode", "Gamemode")
	mGamemode.Click(func() {
		mobs = nil
	})

	mCursorCanHit := systray.AddMenuItem("Dano com cursor", "Dano com cursor")
	mCursorCanHit.Click(func() {
		cfg.CursorCanHit = !cfg.CursorCanHit
		if cfg.CursorCanHit {
			mCursorCanHit.Check()
		} else {
			mCursorCanHit.Uncheck()
		}
	})
	if cfg.CursorCanHit {
		mCursorCanHit.Check()
	}

	mTaskbar := systray.AddMenuItem("Esconder da barra de tarefas", "Esconder da barra de tarefas")
	mTaskbar.Click(func() {
		cfg.SkipTaskbar = !cfg.SkipTaskbar
		if cfg.SkipTaskbar {
			mTaskbar.Check()
		} else {
			mTaskbar.Uncheck()
		}
	})
	if cfg.SkipTaskbar {
		mTaskbar.Check()
	}

	/**
	* ========== MOBS ==========
	 */
	mMobs := systray.AddMenuItem("Mobs", "Mobs")
	mMobsMaximo := mMobs.AddSubMenuItem(fmt.Sprintf("Máximo: %d", cfg.MobsSpawnMax), "")
	mMobsMaximo.Disable()

	mMobs.AddSubMenuItem("Máximo++", "").Click(func() {
		cfg.MobsSpawnMax++
		mMobsMaximo.SetTitle(fmt.Sprintf("Máximo: %d", cfg.MobsSpawnMax))
	})

	mMobs.AddSubMenuItem("Máximo--", "").Click(func() {
		cfg.MobsSpawnMax--
		mMobsMaximo.SetTitle(fmt.Sprintf("Máximo: %d", cfg.MobsSpawnMax))
	})

	mMobsSpawn := mMobs.AddSubMenuItem("Spawn automático", "Spawn automático")
	mMobsSpawn.Click(func() {
		cfg.MobsSpawn = !cfg.MobsSpawn
		if cfg.MobsSpawn {
			mMobsSpawn.Check()
		} else {
			mMobsSpawn.Uncheck()
		}
	})
	if cfg.MobsSpawn {
		mMobsSpawn.Check()
	}

	mMobsDespawn := mMobs.AddSubMenuItem("Despawn automático", "Despawn automático")
	mMobsDespawn.Click(func() {
		cfg.MobsDespawn = !cfg.MobsDespawn
		if cfg.MobsDespawn {
			mMobsDespawn.Check()
		} else {
			mMobsDespawn.Uncheck()
		}
	})
	if cfg.MobsDespawn {
		mMobsDespawn.Check()
	}

	mMobs.AddSubMenuItem("Chama +1", "Chama +1").Click(func() {
		SpawnRandom(1)
	})

	mMobs.AddSubMenuItem("Renovar", "Renovar").Click(func() {
		mobs = nil
	})

	/**
	* ========== SOUNDS ==========
	 */
	mSounds := systray.AddMenuItem("Sons", "Sons")
	mSoundsMusicMuted := mSounds.AddSubMenuItem("Mutar música", "Mutar música")
	mSoundsMusicMuted.Click(func() {
		cfg.MusicMuted = !cfg.MusicMuted
		if cfg.MusicMuted {
			mSoundsMusicMuted.Check()
		} else {
			mSoundsMusicMuted.Uncheck()
		}
	})
	if cfg.MusicMuted {
		mSoundsMusicMuted.Check()
	}
	mSoundsEffectsMuted := mSounds.AddSubMenuItem("Mutar efeitos", "Mutar efeitos")
	mSoundsEffectsMuted.Click(func() {
		cfg.EffectsMuted = !cfg.EffectsMuted
		if cfg.EffectsMuted {
			mSoundsEffectsMuted.Check()
		} else {
			mSoundsEffectsMuted.Uncheck()
		}
	})
	if cfg.EffectsMuted {
		mSoundsEffectsMuted.Check()
	}

	/**
	* ========== END ==========
	 */
	systray.AddSeparator()

	mQuit := systray.AddMenuItem("Abandonar os bixin :(", "Fechar")
	mQuit.Click(func() {
		cfg.Save()
		systray.Quit()
	})
}

func onExit() {}
