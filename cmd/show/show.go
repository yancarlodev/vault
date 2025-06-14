package show

import (
	"fmt"
	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/spf13/cobra"
	"github.com/yancarlodev/vault/infra"
	"os"
)

var ShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the content of a note",
	Long:  "Show the content of a note",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			cobra.CheckErr(err)
		}

		if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
			cobra.CheckErr(err)
		}

		return nil
	},
	Run: run,
}

func run(_ *cobra.Command, args []string) {
	title := args[0]

	_, titleNormalized := infra.NormalizeInput(title)

	notePath := infra.GetDataResourcePath(titleNormalized)

	source, err := os.ReadFile(notePath)

	if err != nil {
		cobra.CheckErr(err)
	}

	result := markdown.Render(string(source), 80, 3)

	fmt.Println(string(result))
}
