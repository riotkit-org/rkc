package generate

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// NewBackupsGenerateCommand creates the new command
func NewBackupsGenerateCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "generate",
		Short: "Generates backup & restore procedures",
		Run: func(command *cobra.Command, args []string) {
			err := command.Help()
			if err != nil {
				logrus.Errorf(err.Error())
			}
		},
	}

	command.AddCommand(NewBackupCommand())
	command.AddCommand(NewRestoreCommand())

	return command
}

func NewBackupCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "backup",
		Short: "Generates a backup procedure",
		Run: func(command *cobra.Command, args []string) {
			err := command.Help()
			if err != nil {
				logrus.Errorf(err.Error())
			}
		},
	}

	return command
}

func NewRestoreCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "restore",
		Short: "Generates a restore procedure",
		Run: func(command *cobra.Command, args []string) {
			err := command.Help()
			if err != nil {
				logrus.Errorf(err.Error())
			}
		},
	}

	return command
}
