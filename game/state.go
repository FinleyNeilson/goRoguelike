package game

import (
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
}

func (state *GameState) currentTileMap() *tiles.TileMap {
	return state.LevelMap.Combine(state.DynamicTiles)
}

func (state *GameState) Update() error {
	ui.SetInputState(state.InputState)

	if state.InputState.Move {
		if state.Player.Move(state) {

		}
	}

	return nil
}

func (state *GameState) Draw(screen *ebiten.Image) {
	state.currentTileMap().Draw(screen)
}

func (state *GameState) Layout(_, _ int) (int, int) {
	return 320, 240
}
