package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var yesFlag bool

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <note>",
	Short: "Delete a note",
	RunE:  runDeleteCmd,
	Args:  cobra.MinimumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		notes, err := listNotes()
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		return notes, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolVarP(&yesFlag, "yes", "y", false, "Skip confirmation prompt")
}

func runDeleteCmd(cmd *cobra.Command, args []string) error {
	for _, noteKey := range args {
		notePath := path.Join(notesDir, noteKey)
		if !yesFlag {
			fmt.Printf("Are you sure you want to delete %s? (y/N): ", noteKey)
			var response string
			fmt.Scanln(&response)
			if response != "y" {
				return nil
			}
		}
		if err := os.Remove(notePath); err != nil {
			return err
		}
	}

	return nil
}
