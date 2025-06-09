package rm

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RmCmd.MarkFlagRequired("title")
}

var RmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a note. Can be chained with whitespaces to delete multiples notes.",
	Long:  "Remove a note. Can be chained with whitespaces to delete multiples notes.",
	Run:   run,
}

func run(_ *cobra.Command, args []string) {
	fmt.Println(args)
}
