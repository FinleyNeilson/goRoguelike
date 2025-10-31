package ui

import (
	"roguelike/tiles"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// InputState The state set by the input
type InputState struct {
	Turn bool
	Move bool
	Dir  tiles.Direction
}

func NewInputState() *InputState {
	return &InputState{
		Turn: false,
		Move: false,
		Dir:  tiles.NoDirection,
	}
}

func resetInputState(state *InputState) {
	state.Turn = false
	state.Move = false
	state.Dir = tiles.NoDirection
}

func SetInputState(inputState *InputState) {
	resetInputState(inputState)

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyQ):
		inputState.moveDirection(tiles.UpLeft)
	case inpututil.IsKeyJustPressed(ebiten.KeyW):
		inputState.moveDirection(tiles.Up)
	case inpututil.IsKeyJustPressed(ebiten.KeyE):
		inputState.moveDirection(tiles.UpRight)
	case inpututil.IsKeyJustPressed(ebiten.KeyS):
		inputState.moveDirection(tiles.Down)
	case inpututil.IsKeyJustPressed(ebiten.KeyZ):
		inputState.moveDirection(tiles.DownLeft)
	case inpututil.IsKeyJustPressed(ebiten.KeyC):
		inputState.moveDirection(tiles.DownRight)
	case inpututil.IsKeyJustPressed(ebiten.KeyA):
		inputState.moveDirection(tiles.Left)
	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		inputState.moveDirection(tiles.Right)

	case inpututil.IsKeyJustPressed(ebiten.KeySpace):
		inputState.Turn = true
	}
}

func (state *InputState) moveDirection(direction tiles.Direction) {
	state.Move = true
	state.Dir = direction
}
