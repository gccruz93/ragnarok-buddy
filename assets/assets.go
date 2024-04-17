package assets

import (
	"embed"
)

//go:embed sounds/*
//go:embed ba_frostjoke.txt
var Embeded embed.FS

var (
	//go:embed icon.jpg
	Icon []byte
	//go:embed icontray.ico
	Icontray []byte

	//go:embed ui/al_warp.ico
	IconAlWarp []byte
	//go:embed ui/card.ico
	IconCard []byte
	//go:embed ui/cursor_attack.ico
	IconCursorAttack []byte
	//go:embed ui/dc_humming.ico
	IconDcHumming []byte
	//go:embed ui/dead_branch.ico
	IconDeadBranch []byte
	//go:embed ui/mo_extremityfist.ico
	IconMoExtremityFirst []byte
	//go:embed ui/pr_resurrection.ico
	IconPrResurrection []byte
	//go:embed ui/tf_hiding.ico
	IconTfHiding []byte
)
