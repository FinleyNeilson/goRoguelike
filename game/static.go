package game

import (
	"roguelike/ui"
)

type Static struct {
	Base
	Walkable bool
}

func (obj Static) IsWalkable() bool {
	return obj.Walkable
}

func (obj Static) OnPlayerEnter(_ *Player) bool {
	if !obj.Walkable {
		ui.LogMessage("You can't go that way theres a " + obj.GetName() + " in the way")
		return false
	}
	return true
}
