package tiles

import (
	"log"
	"roguelike/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type BaseTile struct {
	X, Y     int
	Image    *ebiten.Image
	Name     string
	ObjectID int
}

func (tile *BaseTile) GetPosition() (int, int) {
	return tile.X, tile.Y
}

func (tile *BaseTile) SetPosition(x, y int) {
	if x < MapWidth && y < MapHeight && x >= 0 && y >= 0 {
		tile.X = x
		tile.Y = y
	}
}

func (tile *BaseTile) GetDistanceTo(other Tile) int {
	otherX, otherY := other.GetPosition()
	dx := absInt(tile.X - otherX)
	dy := absInt(tile.Y - otherY)
	if dx > dy {
		return dx
	}
	return dy
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (tile *BaseTile) GetObjectId() (id int) {
	return tile.ObjectID
}

func (tile *BaseTile) Draw(screen *ebiten.Image, tileSize int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(tile.X*tileSize), float64(tile.Y*tileSize))
	screen.DrawImage(tile.Image, op)
}

func NewBaseTile(x, y int, spriteName string) *BaseTile {
	sprite := Get(spriteName)
	if sprite == nil {
		log.Printf("sprite '%s' not found!", spriteName)
	}
	return &BaseTile{
		X:        x,
		Y:        y,
		Image:    sprite,
		Name:     spriteName,
		ObjectID: utils.UniqueID(),
	}
}
