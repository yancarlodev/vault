package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	title     string
	content   string
	isPrivate bool
)

func init() {
	addCmd.Flags().StringVarP(&title, "title", "t", "", "title of the note (required)")
	addCmd.Flags().StringVarP(&content, "content", "c", "", "content of the note")
	addCmd.Flags().BoolVarP(&isPrivate, "private", "p", false, "set the visibility of the note to private")

	addCmd.MarkFlagRequired("title")
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new note",
	Long:  "Create a new private or public note",
	Run:   run,
}

func run(_ *cobra.Command, _ []string) {
	dataFolder := Dirs.DataHome()

	fullFilePath := fmt.Sprintf("%s/%s.md", dataFolder, title)

	if file, _ := os.Stat(fullFilePath); file != nil {
		cobra.CheckErr("A note with the same name already exists")
	}

	if err := os.WriteFile(fullFilePath, []byte(content), 0644); err != nil {
		cobra.CheckErr(err)
	}

	fmt.Printf("Note %s created.", title)
}
