
// func spawnRandom(n int) {
// 	nextSpawn = random(cfg.MobsSpawnSecondsMin, cfg.MobsSpawnSecondsMax)

// 	entry := []string{"angeling", "baphometjr", "ghostring", "kobold_axe", "kobold_hammer", "kobold_mace", "lunatic", "poring", "smokie", "spore"}

// 	if cfg.MobsBlocked != "" {
// 		canSpawn := false
// 		for _, pet := range entry {
// 			if !strings.Contains(cfg.MobsBlocked, pet) {
// 				canSpawn = true
// 				break
// 			}
// 		}
// 		if !canSpawn {
// 			return
// 		}
// 	}

// 	for n > 0 {
// 		pet := entry[random(0, len(entry)-1)]
// 		if cfg.MobsBlocked != "" && strings.Contains(cfg.MobsBlocked, pet) {
// 			continue
// 		}
