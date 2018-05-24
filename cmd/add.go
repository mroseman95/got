package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// TODO check that the current directory is a .got repository

var addCmd = &cobra.Command{
	Use:   "add [path to file to add]",
	Short: "Add a file to be tracked by .got",
	Long:  "Compresses a file and adds it to .got/objects to be tracked",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := add(args[0]); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

// add takes a path to a file, and compresses that file adding it to
// .got/objects
func add(filePath string) error {
	// TODO check that working directory is a got repository

	// get absolute path to current working directory
	workingDirPath, err := os.Getwd()
	if err != nil {
		return err
	}

	// get absolute filepath for given filePath
	fileAbsPath, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}

	// check that filePath is within the current directory or subdirectories
	if !strings.HasPrefix(fileAbsPath, workingDirPath) {
		return errors.New("file is not in current got repository")
	}

	// TODO check if file exists
	// TODO zlib compress file and SHA-1 hash it
	// TODO store compressed file in .got/objects directory with dir/filenam matching hash

	return nil
}
