package cmd

import (
	"io"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <note>",
	Short: "Get contents of a note",
	RunE:  runGetCmd,
	Args:  cobra.ExactArgs(1),
  Aliases: []string{"g", "show", "read"},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func runGetCmd(cmd *cobra.Command, args []string) error {
	noteKey := args[0]
	notePath := path.Join(notesDir, noteKey)

	f, err := os.Open(notePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(os.Stdout, f)
	return err
}
