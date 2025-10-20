package tiles

import (
	"encoding/json"
	"log"
	"os"
)

// The json is storing a 2d array of strings
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
	tm := TileMap{
		Tiles: make([]Tile, MapWidth*MapHeight),
	}
	for y, row := range level.Tiles {
		for x, name := range row {
			tile := NewBaseTile(x, y, name)
			tm.Set(tile)
		}
	}
	return tm
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
