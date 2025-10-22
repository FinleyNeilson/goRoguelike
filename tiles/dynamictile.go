package tiles

type DynamicTile struct {
	BaseTile
}

func (t *DynamicTile) GetAdjacentTileXY(dir Direction) (int, int) {
	dx, dy := dir.Vector()
	return t.X + dx, t.Y + dy
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
			Image: Get(spriteName),
		},
	}
}
