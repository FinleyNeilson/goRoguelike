package tiles

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
)

type TMJLayer struct {
	Data   []int `json:"data"`
	Width  int   `json:"width"`
	Height int   `json:"height"`

	Name string `json:"name"`
}

type TMJMap struct {
	Width  int `json:"width"`
	Height int `json:"height"`

	TileWidth  int        `json:"tilewidth"`
	TileHeight int        `json:"tileheight"`
	Layers     []TMJLayer `json:"layers"`
}

func LoadTMJ(path string) TileMap {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open TMJ: %v", err)
	}
	defer f.Close()

	var tmj TMJMap
	if err := json.NewDecoder(f).Decode(&tmj); err != nil {
		log.Fatalf("failed to decode TMJ: %v", err)
	}

	layer := tmj.Layers[0]
	tm := TileMap{
		Tiles: make([]Tile, tmj.Width*tmj.Height),
	}

	for y := 0; y < layer.Height; y++ {
		for x := 0; x < layer.Width; x++ {
			id := layer.Data[y*layer.Width+x]
			if id == 0 {
				continue
			}
			tile := NewBaseTile(x, y, strconv.Itoa(id))
			tm.Set(tile)
		}
	}

	return tm
}
