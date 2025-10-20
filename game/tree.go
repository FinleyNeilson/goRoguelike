package game

import (
	"roguelike/tiles"
)

type Tree struct {
	Tile *tiles.BaseTile
	Solid bool
}
