package main

import (
	"bytes"
	"fmt"
	"github.com/gobuffalo/packr/v2"
	"regexp"

	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	app := cli.NewApp()
	app.Name = ""
	app.Usage = "SubScan Plugin gen"
	app.UsageText = "pluginName [options]"
	app.HideVersion = true
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "d",
			Value:       "",
			Usage:       "Specify the plugin where the project is located",
			Destination: &p.path,
		},
	}
	if len(os.Args) < 2 || strings.HasPrefix(os.Args[1], "-") {
		_ = app.Run([]string{"-h"})
		return
	}
	p.Name = os.Args[1]
	app.Action = runNew
	args := append([]string{os.Args[0]}, os.Args[2:]...)
	err := app.Run(args)
	if err != nil {
		panic(err)
	}
}

type Plugin struct {
	Name      string
	ModClass  string
	ModPrefix string
	path      string
	ModPath   string
}

var p Plugin

func runNew(_ *cli.Context) (err error) {
	if p.path != "" {
		if p.path, err = filepath.Abs(p.path); err != nil {
			return
		}
		p.path = filepath.Join(p.path, p.Name)
	} else {
		pwd, _ := os.Getwd()
		p.path = filepath.Join(pwd, p.Name)
	}
	p.ModPrefix = p.Name
	p.ModPath = modPath(p.path)
	pathSlice := strings.Split(p.Name, "/")

	p.Name = pathSlice[len(pathSlice)-1]
	p.ModClass = upperCamel(p.Name)
	if err := create(); err != nil {
		return err
	}
	fmt.Printf("Plugin: %s\n", p.Name)
	fmt.Printf("Directory: %s\n\n", p.path)
	return nil
}

//go:generate packr2
func create() (err error) {
	box := packr.New("all", "./templates/plugin")
	if err = os.MkdirAll(p.path, 0755); err != nil {
		return
	}
	for _, name := range box.List() {
		if p.ModPrefix == "" {
			continue
		}
		tmpl, _ := box.FindString(name)
		i := strings.LastIndex(name, string(os.PathSeparator))
		if i > 0 {
			dir := name[:i]
			if err = os.MkdirAll(filepath.Join(p.path, dir), 0755); err != nil {
				return
			}
		}
		if strings.HasSuffix(name, ".tmpl") {
			name = strings.TrimSuffix(name, ".tmpl")
		}
		if err = write(filepath.Join(p.path, name), tmpl); err != nil {
			return
		}
	}

	if err = generate("./..."); err != nil {
		return
	}
	return
}

func generate(path string) error {
	cmd := exec.Command("go", "generate", "-x", path)
	cmd.Dir = p.path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return nil
}

func write(path, tpl string) (err error) {
	data, err := parse(tpl)
	if err != nil {
		return
	}
	return ioutil.WriteFile(path, data, 0644)
}

func parse(s string) ([]byte, error) {
	t, err := template.New("").Parse(s)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err = t.Execute(&buf, p); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func upperCamel(s string) string {
	if len(s) == 0 {
		return ""
	}
	s = strings.ToUpper(string(s[0])) + string(s[1:])
	return s
}

func modPath(p string) string {
	dir := filepath.Dir(p)
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			content, _ := ioutil.ReadFile(filepath.Join(dir, "go.mod"))
			mod := regexpReplace(`module\s+(?P<name>[\S]+)`, string(content), "$name")
			name := strings.TrimPrefix(filepath.Dir(p), dir)
			name = strings.TrimPrefix(name, string(os.PathSeparator))
			if name == "" {
				return fmt.Sprintf("%s/", mod)
			}
			return fmt.Sprintf("%s/%s/", mod, name)
		}
		parent := filepath.Dir(dir)
		if dir == parent {
			return ""
		}
		dir = parent
	}
}
func regexpReplace(reg, src, temp string) string {
	var result []byte
	pattern := regexp.MustCompile(reg)
	for _, submatches := range pattern.FindAllStringSubmatchIndex(src, -1) {
		result = pattern.ExpandString(result, temp, src, submatches)
	}
	return string(result)
}
