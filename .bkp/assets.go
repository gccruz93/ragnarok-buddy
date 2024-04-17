
// var LoadedAudios map[string][]byte

// func LoadAudio(name string) {
// 	if _, ok := LoadedAudios[name]; ok {
// 		return
// 	}
// 	file, err := assets.ReadFile("sound/effect/" + name)
// 	if err != nil {
// 		log.Println("loadAudio ERR: " + name)
// 		return
// 	}
// 	s, _ := wav.DecodeWithSampleRate(sampleRate, bytes.NewReader(file))
// 	b, _ := io.ReadAll(s)
// 	LoadedAudios[name] = b
// }
