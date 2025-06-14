package add

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yancarlodev/vault/infra"
	"os"
)

var (
	title     string
	content   string
	isPrivate bool
)

func init() {
	AddCmd.Flags().StringVarP(&title, "title", "t", "", "title of the note (required)")
	AddCmd.Flags().StringVarP(&content, "content", "c", "", "content of the note")
	AddCmd.Flags().BoolVarP(&isPrivate, "private", "p", false, "set the visibility of the note to private")

	AddCmd.MarkFlagRequired("title")
}

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new note",
	Long:  "Create a new private or public note",
	Run:   run,
}

func run(_ *cobra.Command, _ []string) {
	dataFolder := infra.Dirs.DataHome()

	titleTrimmed, titleNormalized := infra.NormalizeInput(title)

	fullFilePath := fmt.Sprintf("%s/%s.md", dataFolder, titleNormalized)

	if file, _ := os.Stat(fullFilePath); file != nil {
		cobra.CheckErr("A note with the same name already exists")
	}

	if content == "" {
		infra.OpenDefaultApp(fullFilePath)
	} else {
		if err := os.WriteFile(fullFilePath, []byte(content), 0644); err != nil {
			cobra.CheckErr(err)
		}
	}

	if _, err := os.Stat(fullFilePath); os.IsNotExist(err) {
		fmt.Printf("Note \"%s\" was not created.", titleTrimmed)

		return
	}

	fmt.Printf("Note \"%s\" created.", titleTrimmed)
}
