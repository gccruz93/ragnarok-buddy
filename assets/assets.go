package assets

import "embed"

//go:embed cursors/*
//go:embed monster/*
//go:embed npc/*
//go:embed sound/*
//go:embed ba_frostjoke.txt
//go:embed icon.jpg
var Assets embed.FS
