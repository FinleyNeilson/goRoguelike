package main

import (
	"log"
	"roguelike/game"
	"roguelike/tiles"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	internalWidth  = 320
	internalHeight = 240
)

func newGame() *game.GameState {
	playerTile := tiles.NewDynamicTile(5, 5, "hero2")
	player := &game.Player{Tile: playerTile}

	tileMap, gameObject := game.LoadLevel("assets/level1.json")

	enemyObject := game.CreateDynamicObject()
	gameObject[enemyObject.ObjectID] = enemyObject
	enemyTile := enemyObject.GetTile()

	g := &game.GameState{
		LevelMap:     tileMap,
		GameObjectID: gameObject,
		DynamicTiles: []*tiles.DynamicTile{playerTile, enemyTile},
		Player:       player,
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
