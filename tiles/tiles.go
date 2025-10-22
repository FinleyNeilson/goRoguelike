package tiles

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Tile interface {
	Draw(screen *ebiten.Image, tileSize int)
	GetPosition() (x, y int)
	GetObjectId() (id int)
	SetPosition(x, y int)
}
