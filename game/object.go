package game

// Object Might be split up into entity items enemy etc. at some point
type Object interface {
	GetId() int
	GetName() string
	OnPlayerEnter(player *Player) bool
}
