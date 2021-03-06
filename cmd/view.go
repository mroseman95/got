package cmd

import (
	"fmt"

	"github.com/mroseman95/got/object"
	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view [hash of got object]",
	Short: "View the contents of a got object",
	Long:  "Uncompresses a got object with the given hash, and prints the contents to stdout",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := view(args[0]); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}

func view(hash string) error {
	obj, err := object.LoadBlob(gotRootDir, hash)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", obj.Content)

	return nil
}
