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
	notes, err := listNotes()
	if err != nil {
		return err
	}

	for _, note := range notes {
		fmt.Println(note)
	}

	return nil
}

func listNotes() ([]string, error) {
	entries, err := os.ReadDir(notesDir)
	if err != nil {
		return nil, err
	}

	var notes []string
	for _, entry := range entries {
		notes = append(notes, entry.Name())
	}

	return notes, nil
}
