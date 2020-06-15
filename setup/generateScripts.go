package setup

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Eldius/game-manager-go/config"
	"github.com/Eldius/game-manager-go/scripts"
)

/*
GenerateScripts generate the scripts
to be used
*/
func GenerateScripts(cfg config.ManagerConfig) {
	//scriptsFolder := config.GetScriptsFolder()
	//_ = os.MkdirAll(scriptsFolder, os.ModePerm)

	engine := scripts.NewScriptEngine(cfg)
	for _, s := range engine.GetSetupScripts() {
		scriptFolder := filepath.Dir(s.Path)
		log.Printf("---\ngenerating script:\nfolder: %s\nfile:   %s\ntype: %s\n", scriptFolder, s.Path, s.Type)
		if err := os.MkdirAll(scriptFolder, os.ModePerm); err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		s.SaveToFile()
	}
}
