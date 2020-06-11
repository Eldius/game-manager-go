package setup

import (
	"io/ioutil"
	"log"
	"os"
	"fmt"
	"strings"

	"github.com/Eldius/game-manager-go/config"
	"github.com/gobuffalo/packr"
)

const (
	scriptHeader = "header.sh"
	setupPythonEnvironment = "setup/setup_python_environment.sh"
	setupPython = `pyenv install -l`
	helpCommand = `pyenv help`
)

/*
GenerateScripts generate the scripts
to be used
*/
func GenerateScripts() {
	scriptsFolder := config.GetScriptsFolder()
	_ = os.MkdirAll(scriptsFolder, os.ModePerm)

	setupPythonScriptContent := RenderScript(setupPythonEnvironment)

	log.Printf("---\ngenerated script:\n%s\n\n#%s\n---\n", setupPythonScriptContent, config.InstallPythonEnvScript())

	ioutil.WriteFile(config.InstallPythonEnvScript(), []byte(setupPythonScriptContent), os.ModePerm)
}

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
	box := packr.NewBox("./scripts")
	s, err := box.FindString(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
	return s
}
