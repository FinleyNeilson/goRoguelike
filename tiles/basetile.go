package tiles

import (
	"log"
	"roguelike/sprites"

	"github.com/hajimehoshi/ebiten/v2"
)

type BaseTile struct {
	X, Y  int
	Image *ebiten.Image
	Name  string
}

func (t *BaseTile) Draw(screen *ebiten.Image, tileSize int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.X*tileSize), float64(t.Y*tileSize))
	screen.DrawImage(t.Image, op)
}

func (t BaseTile) Position() (int, int) {
	return t.X, t.Y
}

func (t *BaseTile) SetPosition(x, y int) {
	if x < MapWidth && y < MapHeight && x >= 0 && y >= 0 {
		t.X = x
		t.Y = y
	}
}

func NewBaseTile(x, y int, spriteName string) Tile {
	sprite := sprites.Get(spriteName)
	if sprite == nil {
		log.Printf("sprite '%s' not found!", spriteName)
	}
	return &BaseTile{
		X:     x,
		Y:     y,
		Image: sprite,
		Name:  spriteName,
	}
}
