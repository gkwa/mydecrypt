package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

var test2Cmd = &cobra.Command{
	Use:   "test2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test2 called")
		fmt.Println("Absolute path of 'data' directory:", DataDirAbsPath)
	},
}

func init() {
	rootCmd.AddCommand(test2Cmd)
}

const DataDir = "data"

var DataDirAbsPath string

func init() {
	var err error
	DataDirAbsPath, err = filepath.Abs(DataDir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
