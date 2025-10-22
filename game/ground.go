package game

type StaticObject struct {
	BaseGameObject
	Name  string
	Solid bool
}

func (g StaticObject) getName() string {
	return g.Name
}

func (g StaticObject) isSolid() bool {
	return g.Solid
}
