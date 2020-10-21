package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func Test_regexpReplace(t *testing.T) {
	mod := regexpReplace(`module\s+(?P<name>[\S]+)`, `module github.com/itering/subscan-plugin`, "$name")
	assert.Equal(t, "github.com/itering/subscan-plugin", mod)
	mod = regexpReplace(`module\s+(?P<name>[\S]+)`, `xx`, "$name")
	assert.Equal(t, "", mod)
}

func Test_modPath(t *testing.T) {
	path, _ := filepath.Abs("")
	assert.Equal(t, "github.com/itering/subscan-plugin/tools/", modPath(path))
}

func Test_upperCamel(t *testing.T) {
	assert.Equal(t, "Faaaaa", upperCamel("faaaaa"))
	assert.Equal(t, "Hello", upperCamel("hello"))
	assert.Equal(t, "", upperCamel(""))
}

func Test_parse(t *testing.T) {
	cc, err := parse(`package {{.Name}}`)
	assert.NoError(t, err)
	fmt.Println(string(cc))
	assert.Equal(t, "package ", string(cc))
}
