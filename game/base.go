package game

import (
	"fmt"
	"roguelike/tiles"
)

type Base struct {
	Tile       tiles.Tile
	ObjectID   int
	ObjectName string
}

func NewBase(baseTile tiles.Tile, name string) Base {
	return Base{
		Tile:       baseTile,
		ObjectID:   baseTile.GetObjectId(),
		ObjectName: name,
	}
}

func (obj *Base) GetId() int {
	return obj.ObjectID
}

func (obj *Base) GetName() string {
	return obj.ObjectName
}

func (obj *Base) OnPlayerEnter() bool {
	fmt.Println("im game object")
	return true
}
