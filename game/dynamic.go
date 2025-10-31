package game

import (
	"roguelike/tiles"
	"roguelike/ui"
	"strconv"
)

//Maybe the names are bad should be entity? or enemy?

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
	ui.LogMessage("You just attacked " + obj.GetName())
	ui.LogMessage(obj.GetName() + "'s Health is now: " + strconv.Itoa(obj.health))
	return true
}

func (obj *Dynamic) TakeTurn(state *GameState) {
	if obj.Tile.GetDistanceTo(state.Player.Tile) == 1 {
		state.Player.TakeDamage(5)
	}

	ui.LogMessage("Im Death and im taking my turn")
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
