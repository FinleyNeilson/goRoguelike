package game

import (
	"roguelike/tiles"
	"strconv"
)

type GameObject interface {
	getId() string
	getName() string
}

type BaseGameObject struct {
	ObjectID int
	Tile     tiles.Tile
}

func (obj BaseGameObject) getId() string {
	return strconv.Itoa(obj.ObjectID)
}

func (obj BaseGameObject) getName() string {
	return "Default Name"
}

func newBase(baseTile tiles.Tile) BaseGameObject {
	return BaseGameObject{
		Tile:     baseTile,
		ObjectID: baseTile.GetObjectId(),
	}
}
