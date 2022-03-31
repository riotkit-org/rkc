package generate

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"html/template"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Templating struct{}

// RenderTemplate renders a Go-formatted template in order from (stops on first found):
//              1. ~/.rkc/backups/templates/{backup,restore}/{name}.tmpl
//              2. ~/.rkc/backups/templates/base/{backup,restore}/{name}.tmpl
//
//              Templates in first directory are replaced only if the user has not modified them.
func (t *Templating) RenderTemplate(name string, operation string, variables interface{}) (string, error) {
	// load raw template content
	content, err := t.loadTemplate(name, operation)
	if err != nil {
		return "", errors.New(fmt.Sprintf("cannot render template: %s", err))
	}

	// parse
	tpl := template.New(name)
	parsed, parseErr := tpl.Parse(string(content))
	if parseErr != nil {
		return "", errors.New(fmt.Sprintf("cannot render template: %s", err))
	}

	// render
	textBuffer := bytes.NewBufferString("")
	if err := parsed.Execute(textBuffer, variables); err != nil {
		return "", errors.New(fmt.Sprintf("cannot render template, execution failed: %s", err))
	}
	return textBuffer.String(), nil
}

// loadTemplate is reading templates in order from (stops on first found):
//              1. ~/.rkc/backups/templates/{backup,restore}/{name}.tmpl
//              2. ~/.rkc/backups/templates/base/{backup,restore}/{name}.tmpl
//
//              Templates in first directory are replaced only if the user has not modified them.
func (t *Templating) loadTemplate(name string, operation string) ([]byte, error) {
	paths := []string{
		"./cmd/backups/generate/templates/" + operation + "/" + name + ".tmpl", // only in testing
		"~/.rkc/backups/templates/" + operation + "/" + name + ".tmpl",
		"~/.rkc/backups/templates/base/" + operation + "/" + name + ".tmpl",
	}

	for _, path := range paths {
		path, expandErr := expandPath(path)
		if expandErr != nil {
			logrus.Warnf("Cannot expand path: %s", path)
		}
		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}
		return ioutil.ReadFile(path)
	}

	return []byte(""), errors.New(fmt.Sprintf("template not found, looked in those paths: %s", strings.Join(paths, ",")))
}

func expandPath(path string) (string, error) {
	if len(path) == 0 || path[0] != '~' {
		return path, nil
	}

	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr.HomeDir, path[1:]), nil
}
