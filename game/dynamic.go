package game

import (
	"fmt"
	"roguelike/tiles"
	"strconv"
)

type Dynamic struct {
	Base
	health  int
	IsEnemy bool
}

func (obj *Dynamic) GetTile() *tiles.DynamicTile {
	if tile, ok := obj.Tile.(*tiles.DynamicTile); ok {
		return tile
	}
	return nil
}

func (obj *Dynamic) OnAttack(player *Player) bool {
	obj.health = obj.health - player.GetDamage()
	fmt.Println("You just attacked " + obj.GetName())
	fmt.Println(obj.GetName() + "'s Health is now: " + strconv.Itoa(obj.health))
	return true
}

func (obj *Dynamic) OnPlayerEnter(player *Player) bool {
	return false
}

func CreateDynamicObject() *Dynamic {
	tile := tiles.NewDynamicTile(10, 10, "monster9")

	baseGameObject := NewBase(tile, "Death")
	dynamicObject := &Dynamic{
		Base:    baseGameObject,
		health:  100,
		IsEnemy: true}

	return dynamicObject
}
