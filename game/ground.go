package game

import (
	"roguelike/tiles"
)

type Ground struct {
	Tile *tiles.BaseTile
	Solid bool
}
