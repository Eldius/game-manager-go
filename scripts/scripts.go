package scripts

import (
	"bytes"
	"log"
	"strings"
	"text/template"

	"github.com/Eldius/game-manager-go/config"
	"github.com/gobuffalo/packr"
)

const (
	scriptHeader = "shell/header.sh"
)

var (
	box    = packr.NewBox("./templates")
	engine = template.New("scripts_parser")
)

/*
ScriptTemplateVars is a representation of
vars to parse script templates
*/
type ScriptTemplateVars struct {
	WorkspacePath string
}

/*
RenderScript returns a script from template
passed as parameter
*/
func RenderScript(s config.ScriptDef, tmplVars ScriptTemplateVars) string {

	script := GetScriptTemplate(s.Template)

	var tmpl *template.Template
	var err error

	if strings.HasSuffix(s.Path, "sh") {
		header := GetScriptTemplate(scriptHeader)
		tmpl, err = engine.Parse(strings.Join([]string{header, script}, "\n#---\n"))
	} else {
		tmpl, err = engine.Parse(script)
	}
	if err != nil {
		log.Panic(err.Error())
	}

	buf := new(bytes.Buffer)
	tmpl.Execute(buf, tmplVars)
	return buf.String()
}

/*
GetScriptTemplate returns a script template
*/
func GetScriptTemplate(path string) string {
	//box := packr.NewBox("./templates")
	s, err := box.FindString(path)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(s)
	return s
}

/*
GetTemplateVars generate the variables to
parse script templates
*/
func GetTemplateVars() ScriptTemplateVars {
	return ScriptTemplateVars{
		WorkspacePath: config.WorkspaceFolder(),
	}
}
