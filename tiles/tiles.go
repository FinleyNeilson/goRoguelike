package tiles

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Tile interface {
	Draw(screen *ebiten.Image, tileSize int)
	GetPosition() (x, y int)
	SetPosition(x, y int)
	GetDistanceTo(other Tile) int
	GetObjectId() (id int)
}
