package tiles

import (
	"roguelike/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	MapWidth  = 20
	MapHeight = 15
	TileSize  = 16
)

type TileMap struct {
	Tiles []Tile
}

func (tm *TileMap) Get(x, y int) Tile {
	if x < 0 || x >= MapWidth || y < 0 || y >= MapHeight {
		return nil
	}

	index := utils.Flatten(x, y, MapWidth)
	return tm.Tiles[index]
}

func (tm *TileMap) Place(tile Tile) {
	x, y := tile.GetPosition()
	tm.Tiles[utils.Flatten(x, y, MapWidth)] = tile
}

func (tm *TileMap) Draw(screen *ebiten.Image) {
	for _, tile := range tm.Tiles {
		tile.Draw(screen, TileSize)
	}
}

func (tm *TileMap) Combine(dynamicEntities []*DynamicTile) *TileMap {
	resultTileMap := &TileMap{
		Tiles: make([]Tile, MapWidth*MapHeight),
	}
	copy(resultTileMap.Tiles, tm.Tiles)

	for _, entity := range dynamicEntities {
		resultTileMap.Place(entity)
	}

	return resultTileMap
}
