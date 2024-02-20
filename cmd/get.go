package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var noClip bool

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
	getCmd.Flags().BoolVarP(&noClip, "no-clip", "C", false, "Disable clipboard copy")
}

func runGetCmd(cmd *cobra.Command, args []string) error {
	noteKey := args[0]
	notePath := path.Join(notesDir, noteKey)
	noteContent, err := os.ReadFile(notePath)
	if err != nil {
		return err
	}

	trimmed := strings.TrimSpace(string(noteContent))
	if err := clipboard.Init(); err != nil {
		return err
	}

	if !noClip {
		clipboard.Write(clipboard.FmtText, []byte(trimmed))
	}

	if isOutputPiped() {
		fmt.Println(trimmed)
		return nil
	}

	out, err := glamour.RenderWithEnvironmentConfig(trimmed)
	fmt.Println(out)
	return err
}

func isOutputPiped() bool {
	fi, _ := os.Stdout.Stat()
	return (fi.Mode() & os.ModeCharDevice) == 0
}
