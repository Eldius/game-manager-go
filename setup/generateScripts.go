package setup

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Eldius/game-manager-go/config"
	"github.com/Eldius/game-manager-go/scripts"
)

const (
	helpCommand = `pyenv help`
)

/*
GenerateScripts generate the scripts
to be used
*/
func GenerateScripts() {
	//scriptsFolder := config.GetScriptsFolder()
	//_ = os.MkdirAll(scriptsFolder, os.ModePerm)

	templateVars := scripts.GetTemplateVars()
	for _, s := range config.GetAllScripts() {
		scriptFolder := filepath.Dir(s.Path)
		log.Println(s)
		log.Printf("---\ngenerating script:\nfolder: %s\nfile:   %s\n", scriptFolder, s.Path)
		if err := os.MkdirAll(scriptFolder, os.ModePerm); err != nil {
			log.Println(err.Error())
		}
		ioutil.WriteFile(s.Path, []byte(scripts.RenderScript(s, templateVars)), getFileMode(s))
	}
}

func getFileMode(s config.ScriptDef) os.FileMode {
	if strings.HasSuffix(s.Path, "sh") {
		return os.ModePerm
	}
	return 0666
}
