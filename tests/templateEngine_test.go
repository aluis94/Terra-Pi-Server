package tests

import (
	"testing"

	"github.com/aluis94/terra-pi-server/templateEngine"
	"github.com/stretchr/testify/assert"
)

func TestCreateFile(t *testing.T) {
	scriptname := "test.txt"

	templateEngine.CreateFile(scriptname)

	assert.FileExists(t, scriptname, "file not found %s")
}

func TestDeleteFile(t *testing.T) {
	scriptname := "test.txt"

	templateEngine.DeleteFile(scriptname)

	assert.NoFileExists(t, scriptname, "file still exists %s")
}
