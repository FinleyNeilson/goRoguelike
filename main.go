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

type gameState struct {
	Map tiles.TileMap

	GameState []*tiles.DynamicTile
	Player    *game.Player
}

func newGame() *gameState {
	playerTile := tiles.NewDynamicTile(5, 5, "hero1")
	player := &game.Player{Tile: playerTile}

	g := &gameState{
		Map:       tiles.LoadTileMap("assets/level1.json"),
		GameState: []*tiles.DynamicTile{playerTile},
		Player:    player,
	}
	return g
}

func (g *gameState) Update() error {
	g.Player.Move()
	return nil
}

func (g *gameState) Draw(screen *ebiten.Image) {
	mapToDraw := g.Map.Combine(g.GameState)
	mapToDraw.Draw(screen)
}

func (g *gameState) Layout(_, _ int) (int, int) {
	return internalWidth, internalHeight
}

func main() {
	ebiten.SetWindowSize(internalWidth*3, internalHeight*3)
	ebiten.SetWindowTitle("0xGUILD")

	sprites.LoadTilesheet(sprites.DefaultSheet())

	Game := newGame()
	if err := ebiten.RunGame(Game); err != nil {
		log.Fatal(err)
	}

}
