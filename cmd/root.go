package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "flamingo",
	Short: "Swift Configurations & Projects File Navigator",
	Long: `Switch smoothly through different configurations and projects 
        without ever needing to cd into each individual file location`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Flamingo")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
