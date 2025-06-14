package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

func createCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create [project-type] [project-path]",
		Short: "Create a new Go project",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			projectType := args[0]
			projectPath := args[1]

			if projectType == "webapp" {
				if err := os.MkdirAll(projectPath, 0755); err != nil {
					fmt.Printf("Error creating project path %s: %v\n", projectPath, err)
					os.Exit(1)
				}

				dirs := []string{
					filepath.Join(projectPath, "cmd"),
					filepath.Join(projectPath, "internal"),
					filepath.Join(projectPath, "pkg"),
				}

				for _, dir := range dirs {
					err := os.MkdirAll(dir, 0755)

					log.Printf("⌛ Attempting to create directory %s...\n", dir)

					if err != nil {
						log.Fatalf("❌ Error creating directory %s: %v\n", dir, err)
					} else {
						log.Printf("✅ Successfully created directory %s\n", dir)
					}
				}

				log.Printf("✅ Created Go webapp project structure at: %s\n", projectPath)
			}
		},
	}
}

func init() {
	rootCmd.AddCommand(createCmd())
}
