package tiles

import (
	"encoding/json"
	"log"
	"os"
)

type LevelData struct {
	Tiles [][]string `json:"tiles"`
}

func LoadTileMap(path string) TileMap {
	// Load the level json
	level, err := LoadJson(path)
	if err != nil {
		log.Fatalf("failed to load level: %v", err)
	}

	// Create a tilemap of the appropriate size to load with json data
	tileMap := TileMap{
		Tiles: make([]Tile, MapWidth*MapHeight),
	}

	// Create base tiles and place in tilemap
	for y, row := range level.Tiles {
		for x, name := range row {
			tile := NewBaseTile(x, y, name)
			tileMap.Place(tile)
		}
	}
	return tileMap
}

func LoadJson(path string) (LevelData, error) {
	file, err := os.Open(path)
	if err != nil {
		return LevelData{}, err
	}
	defer file.Close()

	var level LevelData
	err = json.NewDecoder(file).Decode(&level)
	return level, err
}
