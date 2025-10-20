package tiles

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Tile interface {
	Draw(screen *ebiten.Image, tileSize int)
	Position() (x, y int)
	SetPosition(x, y int)
}
