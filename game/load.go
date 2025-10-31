package game

import (
	"roguelike/tiles"
	"strings"
)

func LoadLevel(path string) (tiles.TileMap, map[int]Object) {
	tileMap := tiles.LoadTileMap(path)
	objectMap := make(map[int]Object)

	for y := 0; y < tiles.MapHeight; y++ {
		for x := 0; x < tiles.MapWidth; x++ {
			tile, _ := tileMap.Get(x, y)

			baseTile, ok := tile.(*tiles.BaseTile)
			if !ok {
				continue
			}

			switch {
			case strings.HasPrefix(baseTile.Name, "tree"):
				tree := &Static{
					Base:  NewBase(baseTile, "tree"),
					Solid: true,
				}
				objectMap[baseTile.GetObjectId()] = tree
			case strings.HasPrefix(baseTile.Name, "ground"):
				ground := &Static{
					Base:  NewBase(baseTile, "ground"),
					Solid: false,
				}
				objectMap[baseTile.GetObjectId()] = ground
			}
		}
	}

	return tileMap, objectMap
}
