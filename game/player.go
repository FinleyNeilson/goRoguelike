package game

import (
	"fmt"
	"roguelike/tiles"
	"roguelike/ui"
	"strconv"
)

type Player struct {
	Tile   *tiles.DynamicTile
	Health int
	Damage int
}

func NewPlayer() *Player {
	playerTile := tiles.NewDynamicTile(9, 10, "hero2")
	return &Player{
		Tile:   playerTile,
		Health: 100,
		Damage: 5,
	}
}

func (player *Player) GetDamage() int {
	return player.Damage
}

func (player *Player) TakeDamage(i int) {
	ui.LogMessage("Taking damage")
	ui.LogMessage("my health is " + strconv.Itoa(player.Health))
	player.Health = player.Health - i
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
			state.NewTurn = true
			return true
		}
	} else if dynamicObject, ok := moveObject.(*Dynamic); ok {
		if dynamicObject.OnAttack(player) {
			state.NewTurn = true
			return false
		}
	}

	return false
}
