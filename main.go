package main

import (
	"log"
	"roguelike/game"
	"roguelike/tiles"
	"roguelike/ui"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	internalWidth  = 320
	internalHeight = 240
)

func newGame() *game.GameState {
	playerTile := tiles.NewDynamicTile(9, 10, "hero2")
	player := &game.Player{Tile: playerTile, Damage: 5}

	tileMap, gameObject := game.LoadLevel("assets/level1.json")

	enemyObject := game.CreateDynamicObject()
	gameObject[enemyObject.ObjectID] = enemyObject
	enemyTile := enemyObject.GetTile()

	inputState := ui.NewInputState()

	g := &game.GameState{
		LevelMap:     tileMap,
		GameObjectID: gameObject,
		DynamicTiles: []*tiles.DynamicTile{playerTile, enemyTile},
		Player:       player,
		InputState:   inputState,
	}
	return g
}

func main() {
	ebiten.SetWindowSize(internalWidth*3, internalHeight*3)
	ebiten.SetWindowTitle("0xGUILD")

	tiles.LoadTileSheet(tiles.DefaultSheet())

	gs := newGame()
	if err := ebiten.RunGame(gs); err != nil {
		log.Fatal(err)
	}
}
