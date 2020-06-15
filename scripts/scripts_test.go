package scripts

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"io/ioutil"

	"github.com/Eldius/game-manager-go/config"
	"github.com/spf13/viper"
)

func TestGetHeaderTemplate(t *testing.T) {

	expectedHeaderContent := getHeaderText()

	header := GetScriptTemplate(scriptHeader)

	if header != expectedHeaderContent {
		t.Errorf("---\nExpected header content was:\n%s\n\nbut received:\n%s\n---", expectedHeaderContent, header)
	}

}

func TestRenderScript(t *testing.T) {

	cfg := config.GetAppConfig()
	tmpDir, err := ioutil.TempDir("/tmp", "game-manager-test")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpDir)

	fmt.Println("using temp folder", tmpDir)

	viper.SetDefault("workspace", tmpDir)

	s := cfg.GetScriptInfo("install_python_env")

	renderedScript := RenderScript(s, ScriptTemplateVars{
		WorkspacePath: tmpDir,
	})

	if strings.Contains(renderedScript, "(-- header --)") {
		t.Errorf("---\nExpected header to have two '-- header --' line\n---\n%s\n---\n", renderedScript)
	}

	if strings.Index(renderedScript, "(eval \"$\\(pyenv init -\\)\")") > 0 {
		t.Errorf("---\nExpected header to have a 'eval \"$(pyenv init -)\"' line\n---\n%s\n---", renderedScript)
	}
}

func getHeaderText() string {
	return `#!/bin/bash

## -- header -- ##
eval "$(pyenv init -)"
eval "$(pyenv virtualenv-init -)"

## -- header -- ##
`
}
