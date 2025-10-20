package game

import (
	"roguelike/tiles"
)

func LoadLevel(path string) (tiles.TileMap, []GameObject) {
	tileMap := tiles.LoadTileMap(path)

	var objects []GameObject

	for y := 0; y < tiles.MapHeight; y++ {
		for x := 0; x < tiles.MapWidth; x++ {
			tile := tileMap.Get(x, y)

			// Type assertion
			baseTile, ok := tile.(*tiles.BaseTile)
			if !ok {
				continue
			}

			// TODO: This should become its own method
			switch baseTile.Name {
			case "tree":
				tree := &Tree{
					Tile:  baseTile,
					Solid: true,
				}
				objects = append(objects, tree)
			}
		}
	}

	return tileMap, objects
}
