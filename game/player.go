package game

import (
	"roguelike/tiles"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	Tile *tiles.DynamicTile
}

func (player *Player) Move() {
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyW):
		player.Tile.Move(tiles.Up)

	case inpututil.IsKeyJustPressed(ebiten.KeyS):
		player.Tile.Move(tiles.Down)

	case inpututil.IsKeyJustPressed(ebiten.KeyA):
		player.Tile.Move(tiles.Left)

	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		player.Tile.Move(tiles.Right)

	default:
		return
	}
}
