package game

import (
	"roguelike/ui"
)

type Static struct {
	Base
	Solid bool
}

func (obj Static) IsSolid() bool {
	return obj.Solid
}

func (obj Static) OnPlayerEnter(_ *Player) bool {
	if obj.Solid {
		ui.LogMessage("You can't go that way theres a " + obj.GetName() + " in the way")
		return false
	}
	return true
}
