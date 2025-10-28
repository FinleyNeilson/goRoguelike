package game

import (
	"roguelike/tiles"
	"strings"
)

func LoadLevel(path string) (tiles.TileMap, map[int]GameObject) {
	tileMap := tiles.LoadTileMap(path)

	objects := make(map[int]GameObject)

	for y := 0; y < tiles.MapHeight; y++ {
		for x := 0; x < tiles.MapWidth; x++ {
			tile := tileMap.Get(x, y)

			baseTile, ok := tile.(*tiles.BaseTile)
			if !ok {
				continue
			}

			switch {
			case strings.HasPrefix(baseTile.Name, "tree"):
				tree := &StaticObject{
					BaseGameObject: newBase(baseTile),
					Name:           "tree",
					Solid:          true,
				}
				objects[baseTile.GetObjectId()] = tree
			case strings.HasPrefix(baseTile.Name, "ground"):
				ground := &StaticObject{
					BaseGameObject: newBase(baseTile),
					Name:           "ground",
					Solid:          false,
				}
				objects[baseTile.GetObjectId()] = ground
			}
		}
	}

	return tileMap, objects
}
