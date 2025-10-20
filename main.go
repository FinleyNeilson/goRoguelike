package main

import (
	"log"

	"roguelike/game"
	"roguelike/sprites"
	"roguelike/tiles"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	internalWidth  = 320
	internalHeight = 240
)

// TODO: I think we have to move this to the game package Game logic requires 
// access to game state to make sense make it make sense

type gameState struct {
	Map tiles.TileMap

	GameState []*tiles.DynamicTile
	Player    *game.Player
}

func newGame() *gameState {
	playerTile := tiles.NewDynamicTile(5, 5, "hero1")
	player := &game.Player{Tile: playerTile}

	g := &gameState{
		Map:       tiles.LoadTMJ("assets/level1.tmj"),
		GameState: []*tiles.DynamicTile{playerTile},
		Player:    player,
	}
	return g
}

// ---

func (g *gameState) Update() error {
	g.Player.Move()
	return nil
}

func (g *gameState) Draw(screen *ebiten.Image) {
	mapToDraw := g.Map.Combine(g.GameState)
	mapToDraw.Draw(screen)
}

func (g *gameState) Layout(outsideWidth, outsideHeight int) (int, int) {
	return internalWidth, internalHeight
}

func main() {
	ebiten.SetWindowSize(internalWidth*3, internalHeight*3)
	ebiten.SetWindowTitle("0xGUILD")

	sprites.LoadTilesheet(sprites.DefaultSheet())

	game := newGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
