package main

import (
	"workplaces/client/internals"

	"github.com/spf13/cobra"
)



func createUser() *cobra.Command {
	var runCmd = &cobra.Command{
		Use:   "create",
		Short: "create user",
		Run: func(cmd *cobra.Command, args []string) {
			internals.Create()
		},
	}

	return runCmd
}
