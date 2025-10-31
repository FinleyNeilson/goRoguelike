package ui

import (
	"roguelike/tiles"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// InputState The state set by the input
type InputState struct {
	Move bool
	Dir  tiles.Direction
}

func NewInputState() *InputState {
	return &InputState{
		Move: false,
		Dir:  tiles.NoDirection,
	}
}

func SetInputState(state *InputState) {
	resetInputState(state)

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyW):
		state.movePlayer(tiles.Up)

	case inpututil.IsKeyJustPressed(ebiten.KeyS):
		state.movePlayer(tiles.Down)

	case inpututil.IsKeyJustPressed(ebiten.KeyA):
		state.movePlayer(tiles.Left)

	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		state.movePlayer(tiles.Right)
	}
}

func resetInputState(state *InputState) {
	state.Move = false
	state.Dir = tiles.NoDirection
}

func (state *InputState) movePlayer(direction tiles.Direction) {
	state.Move = true
	state.Dir = direction
}
