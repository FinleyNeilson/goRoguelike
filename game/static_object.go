package game

type StaticObject struct {
	BaseGameObject
	Name  string
	Solid bool
}

func (obj StaticObject) getName() string {
	return obj.Name
}

func (obj StaticObject) isSolid() bool {
	return obj.Solid
}
