package generate

import (
	"log"
)

type BackupSnippetGenerationCommand struct {
	DefinitionFile string
	IsKubernetes   bool
	KeyPath        string
}

func (c *BackupSnippetGenerationCommand) backupCommandMain() error {
	t := Templating{}
	variables, loadErr := t.LoadVariables(c.DefinitionFile)
	if loadErr != nil {
		log.Panic(loadErr)
	}

	// todo: JSON schema
	// todo: variables concatenation
	// todo: GPG key support
	// todo: Kubernetes support

	rendered, err := t.RenderTemplate("postgres", "backup", variables)
	if err != nil {
		return err
	}

	println(rendered)
	return nil
}
