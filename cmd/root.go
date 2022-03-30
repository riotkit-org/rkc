package cmd

import (
	"github.com/riotkit-org/rkc/cmd/backups"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Main creates the new command
func Main() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rkc",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()
			if err != nil {
				logrus.Errorf(err.Error())
			}
		},
	}
	cmd.AddCommand(backups.NewBackupsCommand())

	return cmd
}
