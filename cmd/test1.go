package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var test1Cmd = &cobra.Command{
	Use:   "test1",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		absPath, err := filepath.Abs("data")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Absolute path of 'data' directory:", absPath)
	},
}

func init() {
	rootCmd.AddCommand(test1Cmd)
}
