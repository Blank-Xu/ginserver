package routers

import (
	"fmt"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
)

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(fmt.Sprintf("Load Templates Layouts err: [%v]", err))
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.tpl")
	if err != nil {
		panic(fmt.Sprintf("Load Templates Includes err: [%v]", err))
	}

	includes2, err := filepath.Glob(templatesDir + "/includes/**/*.tpl")
	if err != nil {
		panic(fmt.Sprintf("Load Templates Includes err: [%v]", err))
	}
	includes = append(includes, includes2...)
	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}

	templates, err := filepath.Glob(templatesDir + "/*.html")
	if err != nil {
		panic(fmt.Sprintf("Load Templates err: [%v]", err))
	}
	for _, template := range templates {
		r.AddFromFiles(filepath.Base(template), template)
	}
	return r
}
