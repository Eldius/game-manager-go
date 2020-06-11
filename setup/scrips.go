package setup

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/Eldius/game-manager-go/config"
)

const (
	baseScript = `#!/bin/bash

## header ##

env | grep PATH
env | grep HOME

echo "$( which pyenv )"

whoami

eval "$(pyenv init -)"
eval "$(pyenv virtualenv-init -)"

## header ##

pyenv commands

`

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

	//baseScriptParsed := fmt.Sprintf(baseScript, config.GetPyenvBinFolder())

	setupPythonScriptContent := strings.Join([]string{baseScript, setupPython}, "\n")

	log.Println("---\ngenerated script:\n", setupPythonScriptContent, "\n---")

	ioutil.WriteFile(config.GetInstallPythonFile(), []byte(setupPythonScriptContent), os.ModePerm)
}
