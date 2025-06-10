package rm

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yancarlodev/vault/infra"
	"os"
)

var RmCmd = &cobra.Command{
	Use: "rm",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		return nil
	},
	Short: "Remove a note. Can be chained with whitespaces to delete multiples notes.",
	Long:  "Remove a note. Can be chained with whitespaces to delete multiples notes.",
	Run:   run,
}

func run(_ *cobra.Command, notesTitle []string) {
	for _, title := range notesTitle {
		titleTrimmed, titleNormalized := infra.NormalizeInput(title)

		fullFilePath := fmt.Sprintf("%s/%s.md", infra.Dirs.DataHome(), titleNormalized)

		err := os.Remove(fullFilePath)

		if os.IsNotExist(err) {
			fmt.Println("[WARN]", err)

			continue
		}

		if err != nil {
			cobra.CheckErr(err)
		}

		fmt.Printf("Note \"%s\" removed", titleTrimmed)
	}
}
