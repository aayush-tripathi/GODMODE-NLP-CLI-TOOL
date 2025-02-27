package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// base command.
var rootCmd = &cobra.Command{
	Use:   "nlpcli",
	Short: "CLI tool for NLP-to-Bash command generation",
	Long:  `This CLI tool sends natural language queries to an NLP backend and returns suggested Bash commands.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
