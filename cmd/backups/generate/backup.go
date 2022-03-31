package generate

type BackupSnippetGenerationCommand struct {
	DefinitionFile string
	IsKubernetes   bool
	KeyPath        string
}

func (c *BackupSnippetGenerationCommand) backupCommandMain() error {
	t := Templating{}
	rendered, err := t.RenderTemplate("postgres", "backup", map[string]string{})
	if err != nil {
		return err
	}

	println(rendered)
	return nil
}
