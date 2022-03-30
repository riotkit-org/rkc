package backups

import (
	"github.com/riotkit-org/rkc/cmd/backups/generate"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewBackupsCommand creates the new command
func NewBackupsCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "backups",
		Short: "Generates backup & restore procedures",
		Run: func(command *cobra.Command, args []string) {
			err := command.Help()
			if err != nil {
				logrus.Errorf(err.Error())
			}
		},
	}

	command.AddCommand(generate.NewBackupsGenerateCommand())

	return command
}
