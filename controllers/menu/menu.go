package menu

type Menu struct {
	Id        int
	Name      string
	Path      string
	Icon      string
	ChildMenu []*Menu
}
