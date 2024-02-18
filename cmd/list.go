package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all notes",
	RunE:  runListCmd,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func runListCmd(cmd *cobra.Command, args []string) error {
	entries, err := os.ReadDir(notesDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fmt.Println(entry.Name())
	}

	return nil
}
