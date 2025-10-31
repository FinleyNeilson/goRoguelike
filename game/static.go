package game

import (
	"fmt"
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
		fmt.Println("You hit a " + obj.GetName())
		return false
	}
	return true
}
