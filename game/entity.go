package game

import (
	"roguelike/tiles"
	"roguelike/ui"
	"strconv"
)

type Entity struct {
	Base
	health  int
	IsEnemy bool
}

func (obj *Entity) GetTile() *tiles.DynamicTile {
	if tile, ok := obj.Tile.(*tiles.DynamicTile); ok {
		return tile
	}
	return nil
}

func (obj *Entity) OnAttack(player *Player) bool {
	obj.health = obj.health - player.GetDamage()
	ui.LogMessage("You just attacked " + obj.GetName())
	ui.LogMessage(obj.GetName() + "'s Health is now: " + strconv.Itoa(obj.health))
	return true
}

func (obj *Entity) TakeTurn(state *GameState) {
	if obj.Tile.GetDistanceTo(state.Player.Tile) == 1 {
		state.Player.TakeDamage(5)
	}

	ui.LogMessage("Im Death and im taking my turn")
}

func (obj *Entity) OnPlayerEnter(player *Player) bool {
	return false
}

func CreateDynamicObject() *Entity {
	tile := tiles.NewDynamicTile(10, 10, "monster9")

	baseGameObject := NewBase(tile, "Death")
	dynamicObject := &Entity{
		Base:    baseGameObject,
		health:  100,
		IsEnemy: true}

	return dynamicObject
}
