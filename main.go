package main

import (
	"github.com/itsHabib/forkutil/search"
	"github.com/spf13/cobra"
)

/// Root Command used to work with other sub commands
var rootCmd = &cobra.Command{}

func main() {
	rootCmd.Execute()
}

func init() {
	rootCmd = &cobra.Command{
		Use:   "forkutil",
		Short: "Project Forking tool For Github",
	}
	rootCmd.AddCommand(search.SearchCmd)
}
