package main

import (
	"workplaces/client/internals"

	"github.com/spf13/cobra"
)

func updateUser() *cobra.Command {
	var runCmd = &cobra.Command{
		Use:   "update",
		Short: "update user where flags id pc mac_adress user",
		Run: func(cmd *cobra.Command, args []string) {
			internals.Update(args)
		},
	}

	return runCmd
}
