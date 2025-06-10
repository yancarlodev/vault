package list

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yancarlodev/vault/infra"
	"os"
	"strings"
	"time"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all notes",
	Long:  "List all notes",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	notes, err := os.ReadDir(infra.Dirs.DataHome())

	if err != nil {
		cobra.CheckErr(err)
	}

	for _, note := range notes {
		noteTitleUnderScoreReplaced := strings.ReplaceAll(note.Name(), "_", " ")
		noteTitleWithoutExtension, _ := strings.CutSuffix(noteTitleUnderScoreReplaced, ".md")

		fileInfo, err := note.Info()

		if err != nil {
			cobra.CheckErr(err)
		}

		fmt.Printf("[%s] %s\n", fileInfo.ModTime().Format(time.RFC822), noteTitleWithoutExtension)
	}
}
