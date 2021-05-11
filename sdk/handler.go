package sdk

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SampleCmd = &cobra.Command{
	Use:   "Sample",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sample command called")
	},
}
