package main

import (
	"workplaces/client/internals"

	"github.com/spf13/cobra"
)



func createUser() *cobra.Command {
	var runCmd = &cobra.Command{
		Use:   "create",
		Short: "create user where flags pc_name mac_adress user_name",
		Run: func(cmd *cobra.Command, args []string) {
			internals.Create(args)
		},
	}

	return runCmd
}
