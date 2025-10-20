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
	return tm.Tiles[utils.Flatten(x, y, MapWidth)]
}

func (tm *TileMap) Set(tile Tile) {
	x, y := tile.Position()
	tm.Tiles[utils.Flatten(x, y, MapWidth)] = tile
}

func (tm TileMap) Draw(screen *ebiten.Image) {
	for _, tile := range tm.Tiles {
		tile.Draw(screen, TileSize)
	}
}

func (staticTileMap TileMap) Combine(dynamicEntities []*DynamicTile) TileMap {
	resultTileMap := TileMap{
		Tiles: make([]Tile, MapWidth*MapHeight),
	}
	copy(resultTileMap.Tiles, staticTileMap.Tiles)

	for _, entity := range dynamicEntities {
		resultTileMap.Set(entity)
	}

	return resultTileMap
}

func DefaultTileMap() TileMap {
	tileMap := TileMap{
		Tiles: make([]Tile, MapWidth*MapHeight),
	}

	for i := 0; i < MapWidth*MapHeight; i++ {
		x, y := utils.Unflatten(i, MapWidth)
		tileMap.Tiles[i] = NewBaseTile(x, y, "tree17")
	}

	return tileMap
}
