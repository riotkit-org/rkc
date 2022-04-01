package generate

type BackupSnippetGenerationCommand struct {
	Template       string
	DefinitionFile string
	IsKubernetes   bool
	KeyPath        string
}

func (c *BackupSnippetGenerationCommand) backupCommandMain() error {
	t := Templating{}
	variables, loadErr := t.LoadVariables(c.DefinitionFile)
	if loadErr != nil {
		return loadErr
	}

	// todo: JSON schema
	// todo: variables concatenation
	// todo: GPG key support
	// todo: Kubernetes support

	rendered, err := t.RenderTemplate(c.Template, "backup", variables)
	if err != nil {
		return err
	}

	println(rendered)
	return nil
}
