package scripts

import (
	"testing"
)

func TestGetHeaderTemplate(t *testing.T) {

	expectedHeaderContent := `#!/bin/bash

## -- header -- ##
eval "$(pyenv init -)"
eval "$(pyenv virtualenv-init -)"

## -- header -- ##
`

	header := GetScriptTemplate(scriptHeader)

	if header != expectedHeaderContent {
		t.Errorf("---\nExpected header content was:\n%s\n\nbut received:\n%s\n---", expectedHeaderContent, header)
	}

}