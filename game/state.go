package game

import (
	"fmt"
	"roguelike/tiles"
	"roguelike/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameState struct {
	LevelMap     tiles.TileMap
	DynamicTiles []*tiles.DynamicTile
	GameObjectID map[int]Object
	InputState   *ui.InputState
	Player       *Player
	NewTurn      bool
}

func (state *GameState) currentTileMap() *tiles.TileMap {
	return state.LevelMap.Combine(state.DynamicTiles)
}

func (state *GameState) Update() error {
	state.NewTurn = false
	ui.SetInputState(state.InputState)

	// If space is pressed you do nothing for this turn
	if state.InputState.Turn {
		state.NewTurn = true
	}

	// Turn is set inside Player.Move function
	if state.InputState.Move {
		state.Player.Move(state)
	}

	if state.NewTurn {
		for _, gameObject := range state.GameObjectID {
			if dynamic, ok := gameObject.(*Dynamic); ok {
				dynamic.TakeTurn(state)
			}
		}
		fmt.Println("NewTurn")
	}

	return nil
}

func (state *GameState) Draw(screen *ebiten.Image) {
	state.currentTileMap().Draw(screen)
}

func (state *GameState) Layout(_, _ int) (int, int) {
	return 320, 240
}
