package game

import (
	"roguelike/tiles"
)

type GameObject interface {
	getId() int
	getName() string
	onPlayerEnter() bool
}

type BaseGameObject struct {
	ObjectID   int
	ObjectName string
	Tile       tiles.Tile
}

func newBase(baseTile tiles.Tile) BaseGameObject {
	return BaseGameObject{
		Tile:       baseTile,
		ObjectID:   baseTile.GetObjectId(),
		ObjectName: "Default Name",
	}
}

func (obj *BaseGameObject) getId() int {
	return obj.ObjectID
}

func (obj *BaseGameObject) getName() string {
	return obj.ObjectName
}
