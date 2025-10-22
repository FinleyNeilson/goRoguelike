package game

import (
	"fmt"
	"roguelike/tiles"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	Tile *tiles.DynamicTile
	X, Y int
}

func (player *Player) Move(state *GameState) {
	state.currentTileMap()
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyW):
		player.move(tiles.Up, state)

	case inpututil.IsKeyJustPressed(ebiten.KeyS):
		player.move(tiles.Down, state)

	case inpututil.IsKeyJustPressed(ebiten.KeyA):
		player.move(tiles.Left, state)

	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		player.move(tiles.Right, state)

	default:
		return
	}
}

func (player *Player) move(direction tiles.Direction, state *GameState) {
	moveX, moveY := player.Tile.GetAdjacentTileXY(direction)
	moveTile := state.currentTileMap().Get(moveX, moveY)
	if moveTile == nil {
		return
	}

	moveObject := state.GameObjectID[(moveTile.GetObjectId())]

	if ground, ok := moveObject.(*StaticObject); ok && ground.isSolid() {
		fmt.Println("You hit a tree")
	} else {
		player.Tile.Move(direction)
	}
}
