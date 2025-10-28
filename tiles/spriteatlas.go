package tiles

import (
	"encoding/json"
	"image"
	_ "image/png"
	"log"
	"os"

	"roguelike/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheet struct {
	size int
	cols int
	rows int
	path string
}

func DefaultSheet() SpriteSheet {
	return SpriteSheet{
		size: 16,
		cols: 88,
		rows: 39,
		path: "assets/tilesheet.png",
	}
}

var spritesByName = map[string]*ebiten.Image{}

func Get(name string) *ebiten.Image {
	return spritesByName[name]
}

func LoadTileSheet(spriteSheet SpriteSheet) {

	file, err := os.Open(spriteSheet.path)
	if err != nil {
		log.Fatal("failed to open tile sheet:", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal("failed to decode tile sheet:", err)
	}

	sheet := ebiten.NewImageFromImage(img)

	spriteNames, err := LoadSpriteMap("assets/sprites.json")
	if err != nil {
		log.Fatal(err)
	}

	for name, id := range spriteNames {
		x, y := utils.Unflatten(id, spriteSheet.cols)

		rect := image.Rect(
			x*spriteSheet.size,
			y*spriteSheet.size,
			(x+1)*spriteSheet.size,
			(y+1)*spriteSheet.size,
		)

		sub := sheet.SubImage(rect).(*ebiten.Image)
		spritesByName[name] = sub
	}

}

func LoadSpriteMap(path string) (map[string]int, error) {
	// Open sprite map
	file, err := os.Open(path)
	if err != nil {
		return nil, err

	}
	defer file.Close()

	var spriteMap map[string]int
	err = json.NewDecoder(file).Decode(&spriteMap)
	return spriteMap, err
}
