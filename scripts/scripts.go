package scripts

import (
	"log"
	"bytes"
	"strings"
	"text/template"

	"github.com/gobuffalo/packr"
	"github.com/Eldius/game-manager-go/config"
)

const (
	scriptHeader = "shell/header.sh"
)

var (
	box = packr.NewBox("./templates")
	engine = template.New("scripts_parser")
)

/*
RenderScript returns a script from template
passed as parameter
*/
func RenderScript(path string) string {
	header := GetScriptTemplate(scriptHeader)
	script := GetScriptTemplate(path)

	templateMap := map[string]string {
		"WORKSPACE_PATH": config.WorkspaceFolder(),
	}

	tmpl, err := engine.Parse(strings.Join([]string{header, script}, "\n#---\n"))
	if err != nil {
		log.Panic(err.Error())
	}
	buf := new(bytes.Buffer)
	tmpl.Execute(buf, templateMap)
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
