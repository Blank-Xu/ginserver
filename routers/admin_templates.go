package routers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/multitemplate"
)

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	templatesDir = filepath.Base(templatesDir)

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.*")
	if err != nil {
		panic(fmt.Sprintf("Load Templates Layouts err: [%v]", err))
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.*")
	if err != nil {
		panic(fmt.Sprintf("Load Templates Includes err: [%v]", err))
	}

	includesSub, err := filepath.Glob(templatesDir + "/includes/**/*.*")
	if err != nil {
		panic(fmt.Sprintf("Load Templates Includes err: [%v]", err))
	}
	includes = append(includes, includesSub...)
	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		include = strings.Replace(include, string(os.PathSeparator), "/", -1)
		include = strings.TrimPrefix(include, templatesDir+"/includes/")
		r.AddFromFiles(include, files...)
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
