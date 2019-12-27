package routers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/multitemplate"
)

func loadTemplates(templateDir string) multitemplate.Renderer {
	render := multitemplate.NewRenderer()
	// load template
	template, err := filepath.Glob(templateDir + "/*.*")
	if err != nil {
		panic(fmt.Sprintf("Load Template err: [%v]", err))
	}
	for _, v := range template {
		render.AddFromFiles(filepath.Base(v), v)
	}

	layout, err := filepath.Glob(templateDir + "/layout/*.*")
	if err != nil {
		panic(fmt.Sprintf("Load Template Layout err: [%v]", err))
	}

	include, err := filepath.Glob(templateDir + "/include/*.*")
	if err != nil {
		panic(fmt.Sprintf("Load Template Include err: [%v]", err))
	}

	includeSub, err := filepath.Glob(templateDir + "/include/**/*.*")
	if err != nil {
		panic(fmt.Sprintf("Load Template Include err: [%v]", err))
	}
	include = append(include, includeSub...)
	// Generate our template map from our layout/ and v/ directories
	for _, v := range include {
		layoutCopy := make([]string, len(layout))
		copy(layoutCopy, layout)
		files := append(layoutCopy, v)
		v = strings.Replace(v, string(os.PathSeparator), "/", -1)
		v = strings.TrimPrefix(v, templateDir+"/include/")
		render.AddFromFiles(v, files...)
	}
	return render
}
