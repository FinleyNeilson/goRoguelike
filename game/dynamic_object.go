package game

import (
	"roguelike/tiles"
	"roguelike/utils"
)

type DynamicObject struct {
	BaseGameObject
	Name    string
	isEnemy bool
}

func (obj *DynamicObject) GetName() string {
	return obj.Name
}

func (obj *DynamicObject) GetTile() *tiles.DynamicTile {
	if tile, ok := obj.Tile.(*tiles.DynamicTile); ok {
		return tile
	}
	return nil
}

func CreateDynamicObject() *DynamicObject {
	tile := tiles.NewDynamicTile(10, 10, "monster9")
	objectID := utils.UniqueID()

	baseGameObject := BaseGameObject{
		Tile:     tile,
		ObjectID: objectID}

	dynamicObject := &DynamicObject{
		BaseGameObject: baseGameObject,
		isEnemy:        true}

	return dynamicObject
}
