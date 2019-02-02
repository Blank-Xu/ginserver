package func_map

import (
	"html/template"
)

var funMap = make(template.FuncMap)

func GetFunMap() template.FuncMap {
	return funMap
}
