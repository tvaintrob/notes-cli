package cmd

import (
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"
)

const defaultEditor = "vim"

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:     "edit",
	Short:   "Edit contents of a note",
	Args:    cobra.RangeArgs(1, 2),
	RunE:    runEditCmd,
	Aliases: []string{"e", "set", "update"},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		notes, err := listNotes()
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		return notes, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}

func runEditCmd(cmd *cobra.Command, args []string) error {
	var noteKey, noteValue string
	noteKey = args[0]
	notePath := path.Join(notesDir, noteKey)

	// Create the notes directory if it doesn't exist
	if _, err := os.Stat(notesDir); os.IsNotExist(err) {
		err := os.Mkdir(notesDir, 0755)
		if err != nil {
			return err
		}
	}

	if len(args) == 2 {
		noteValue = args[1]
		return os.WriteFile(notePath, []byte(noteValue), 0644)
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = defaultEditor
	}

	return editInEditor(editor, notePath)
}

func editInEditor(editor, notePath string) error {
	cmd := exec.Command(editor, notePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
