package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// root command
var rootCmd = &cobra.Command{Use: "hoge"}

func init() {
	// subcommand 1
	var detectCmd = &cobra.Command{
		Use:   "detect",
		Short: "detect charactor encoding",
		Run: func(cmd *cobra.Command, args []string) {
			Detect(cmd, args)
		},
	}

	// subcommand 2
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			Version(cmd, args)
		},
	}

	rootCmd.AddCommand(detectCmd, versionCmd)
}

// command execute
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
