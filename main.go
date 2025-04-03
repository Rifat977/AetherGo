package main

import (
	"fmt"
	"os"

	"AetherGo/internal/app"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "aethergo",
		Short: "AetherGo CLI",
		Long:  "AetherGo is a lightweight Go web framework with a CLI tool for project management.",
	}

	var createCmd = &cobra.Command{
		Use:   "create [project-name]",
		Short: "Create a new AetherGo project",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			projectName := args[0]
			err := app.CreateNewProject(projectName)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			fmt.Println("Project created successfully:", projectName)
		},
	}

	rootCmd.AddCommand(createCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
