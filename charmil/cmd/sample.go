package cmd

import (
	handler "mysdk/sdk"
)

func init() {
	rootCmd.AddCommand(handler.SampleCmd)
}
