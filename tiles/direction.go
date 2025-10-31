package tiles

type Direction int

const (
	NoDirection Direction = iota
	Up
	UpRight
	Right
	DownRight
	Down
	DownLeft
	Left
	UpLeft
)

func (d Direction) Vector() (dx, dy int) {
	switch d {
	case Up:
		return 0, -1
	case UpRight:
		return 1, -1
	case Right:
		return 1, 0
	case DownRight:
		return 1, 1
	case Down:
		return 0, 1
	case DownLeft:
		return -1, 1
	case Left:
		return -1, 0
	case UpLeft:
		return -1, -1
	case NoDirection:
		return 0, 0
	default:
		return 0, 0
	}
}
