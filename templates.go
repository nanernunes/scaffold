package scaffold

import (
	"bytes"
	"embed"
	"fmt"
	"go/build"
	"go/format"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/fatih/color"
	"github.com/nanernunes/scaffold/pkg/helpers"
)

//go:embed pkg/templates/*
var templates embed.FS

type Template struct {
	ProjectName    *helpers.Name
	ResourceName   *helpers.Name
	ResourceFields map[*helpers.Name]string
	Imports        []string
}

func NewTemplate(project string) *Template {
	return &Template{
		ProjectName: helpers.NewName(project),
		Imports:     make([]string, 0),
	}
}

func (t *Template) CreateFile(src string, dst string, any interface{}) {
	currentDir, _ := os.Getwd()

	if path.Base(currentDir) != t.ProjectName.ToString() {
		currentDir = filepath.Join(currentDir, t.ProjectName.ToString())
	}

	targetPath := filepath.Join(currentDir, filepath.FromSlash(dst))

	os.MkdirAll(path.Dir(targetPath), os.ModePerm)

	source, _ := templates.ReadFile(src)

	fmt.Printf("    %s  %s\n", color.GreenString("create"), dst)

	tmpl, _ := template.New("temp").Parse(string(source))

	var buf bytes.Buffer

	_ = tmpl.Execute(&buf, any)

	if strings.HasSuffix(targetPath, ".go") {
		linted, _ := format.Source(buf.Bytes())
		buf = *bytes.NewBuffer(linted)
	}

	ioutil.WriteFile(targetPath, buf.Bytes(), 0644)

}

func (t *Template) CreateFiles(files map[string]string) {
	for source, target := range files {
		t.CreateFile(source, target, t)
	}
}

func (t *Template) Skel() {
	templates := map[string]string{
		"pkg/templates/.gitignore.tmpl":               ".gitignore",
		"pkg/templates/go.mod.tmpl":                   "go.mod",
		"pkg/templates/Makefile.tmpl":                 "Makefile",
		"pkg/templates/api/api.go.tmpl":               "api/api.go",
		"pkg/templates/api/swagger/swagger.go.tmpl":   "api/.swagger/swagger.go",
		"pkg/templates/api/controllers/index.go.tmpl": "api/controllers/index.go",
		"pkg/templates/cmd/api.go.tmpl":               "cmd/api.go",
		"pkg/templates/config/api.go.tmpl":            "config/api.go",
		"pkg/templates/config/application.go.tmpl":    "config/application.go",
		"pkg/templates/config/config.go.tmpl":         "config/config.go",
	}

	t.CreateFiles(templates)

	cmd := helpers.NewCommand(true, t.ProjectName.ToString())

	cmd.Run("make deps")
	cmd.Run("make build-swagger")
	cmd.Run("go mod tidy")
	cmd.Run("go get ./...")
	cmd.Run("git init .")
	cmd.Run("git add .")
}

func (t *Template) AddMVC(name string, fields map[*helpers.Name]string) {
	t.ResourceName = helpers.NewName(name)
	t.ResourceFields = fields

	var importTime bool

	for _, typ := range fields {
		if typ == "time.Time" {
			importTime = true
		}
	}

	if importTime {
		t.Imports = append(t.Imports, "time")
	}

	templates := map[string]string{
		"pkg/templates/api/controllers/controller.go.tmpl": fmt.Sprintf("api/controllers/%s.go", t.ResourceName.Lower().Plural().ToString()),
		"pkg/templates/app/models/model.go.tmpl":           fmt.Sprintf("app/models/%s.go", t.ResourceName.Lower().ToString()),
		"pkg/templates/app/services/service.go.tmpl":       fmt.Sprintf("app/services/%s.go", t.ResourceName.Lower().ToString()),
	}

	t.CreateFiles(templates)

	cmd := helpers.NewCommand(true, t.ProjectName.ToString())
	cmd.Run(fmt.Sprintf("%s/bin/swag fmt api/controllers/%s.go", build.Default.GOPATH, t.ResourceName.Lower().Plural().ToString()))

	helpers.AppendInFile(
		"api/api.go",
		"func NewAPI",
		"\treturn",
		fmt.Sprintf("\tcontrollers.New%sController(api.Router, application)\n", t.ResourceName.Camel().Plural().ToString()),
	)

	appConfig := "config/application.go"
	importModel := fmt.Sprintf(`"%s/app/models"`, t.ProjectName.Lower().ToString())

	if !helpers.FileContains(appConfig, importModel) {
		helpers.AppendInFile(appConfig, "import", ")", fmt.Sprintf("\t%s\n", importModel))
	}

	helpers.AppendInFile(
		appConfig,
		"\tmodels := []interface{}{",
		"\t}",
		fmt.Sprintf("\t\tmodels.%s{},\n", t.ResourceName.Camel().ToString()),
	)
}
