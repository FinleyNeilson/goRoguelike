package tiles

import (
	"roguelike/sprites"
)

type DynamicTile struct {
	BaseTile
}

func (t *DynamicTile) Move(dir Direction) {
	dx, dy := dir.Vector()
	t.SetPosition(t.X+dx, t.Y+dy)
}
 
func NewDynamicTile(x, y int, spriteName string) *DynamicTile {
	return &DynamicTile{
		BaseTile: BaseTile{
			X:     x,
			Y:     y,
			Image: sprites.Get(spriteName),
		},
	}
}
