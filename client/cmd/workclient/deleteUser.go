package main

import (
	"workplaces/client/internals"

	"github.com/spf13/cobra"
)

func deleteUser() *cobra.Command {
	var runCmd = &cobra.Command{
		Use:   "delete",
		Short: "delete user where flags id",
		Run: func(cmd *cobra.Command, args []string) {
			internals.Delete(args)
		},
	}

	return runCmd
}
