package game

import (
	"fmt"
	"roguelike/tiles"
)

type Player struct {
	Tile   *tiles.DynamicTile
	Damage int
}

func NewPlayer() *Player {
	return &Player{}
}

func (player *Player) GetDamage() int {
	return player.Damage
}

func (player *Player) Move(state *GameState) bool {
	direction := state.InputState.Dir

	// This can be improved
	x, y := player.Tile.GetAdjacentTileXY(direction)
	moveTile, ok := state.currentTileMap().Get(x, y)

	if !ok {
		fmt.Println("no no no, cant go there")
		return false
	}

	moveObject := state.GameObjectID[(moveTile.GetObjectId())]
	if moveObject == nil {
		return false
	}

	if _, ok := moveObject.(*Static); ok {
		if moveObject.OnPlayerEnter(player) {
			player.Tile.Move(direction)
			return true
		}
	} else if dynamicObject, ok := moveObject.(*Dynamic); ok {
		if dynamicObject.OnAttack(player) {
			// I mean this is confusing change later, probably.
			// It didn't actually move but its counted as a turn.
			return true
		}
	}

	return false
}
