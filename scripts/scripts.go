package scripts

import (
	"log"
	"strings"

	"github.com/gobuffalo/packr"
)

const (
	scriptHeader = "shell/header.sh"
)

var box = packr.NewBox("./templates")

/*
RenderScript returns a script from template
passed as parameter
*/
func RenderScript(path string) string {
	header := GetScriptTemplate(scriptHeader)
	script := GetScriptTemplate(path)

	return strings.Join([]string{header, script}, "\n#---\n")
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
