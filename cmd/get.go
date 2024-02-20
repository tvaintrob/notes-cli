package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get <note>",
	Short:   "Get contents of a note",
	RunE:    runGetCmd,
	Args:    cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Aliases: []string{"g", "show", "read"},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		notes, err := listNotes()
		if err != nil {
			return nil, cobra.ShellCompDirectiveError
		}
		return notes, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func runGetCmd(cmd *cobra.Command, args []string) error {
	noteKey := args[0]
	notePath := path.Join(notesDir, noteKey)
	noteContent, err := os.ReadFile(notePath)
	if err != nil {
		return err
	}

	if isOutputPiped() {
		fmt.Println(string(noteContent))
	} else {
		out, err := glamour.Render(string(noteContent), "dark")
		fmt.Println(out)
		return err
	}

	return nil
}

func isOutputPiped() bool {
	fi, _ := os.Stdout.Stat()
	return (fi.Mode() & os.ModeCharDevice) == 0
}
