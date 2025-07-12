package edit

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yancarlodev/vault/infra"
)

var (
	title     string
	content   string
	isPrivate bool
	isPublic  bool
)

func init() {
	EditCmd.Flags().StringVarP(&title, "title", "t", "", "title of the note (required)")
	EditCmd.Flags().StringVarP(&content, "content", "c", "", "content of the note")
	EditCmd.Flags().BoolVar(&isPrivate, "private", false, "set the visibility of the note to private")
	EditCmd.Flags().BoolVar(&isPublic, "public", false, "set the visibility of the note to public")
}

var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a existing note",
	Long:  "Edit a existing note",
	Run:   run,
}

func run(_ *cobra.Command, args []string) {
	noteToEdit := args[0]

	noteToEditPath := infra.GetDataResourcePath(noteToEdit)

	if file, _ := os.Stat(noteToEditPath); file == nil {
		cobra.CheckErr("Note not found.")
	}

	titleTrimmed, titleNormalized := infra.NormalizeInput(title)

	notePath := infra.GetDataResourcePath(titleNormalized)

	if file, _ := os.Stat(notePath); file != nil {
		cobra.CheckErr("A note with the same name already exists")
	}

	if title == "" && content == "" {
		infra.OpenDefaultApp(noteToEditPath)
	} else {
		if title != "" {
			if err := os.Rename(noteToEditPath, notePath); err != nil {
				cobra.CheckErr(err)
			}
		}

		if content != "" {
			if err := os.WriteFile(notePath, []byte(content), 0644); err != nil {
				cobra.CheckErr(err)
			}
		}

		if _, err := os.Stat(notePath); os.IsNotExist(err) {
			fmt.Print(err)
			fmt.Printf("Note \"%s\" was not updated.", titleTrimmed)

			return
		}
	}

	fmt.Printf("Note \"%s\" updated.", noteToEdit)
}
