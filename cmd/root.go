package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// defining a flag
var shout bool

var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A Simple Todo CLI",
	Args:  cobra.MinimumNArgs(1), // only allow 1 or more arguments
	Run: func(cmd *cobra.Command, args []string) {
		greeting := fmt.Sprintf("Hello, %s\n", args[0])

		if shout {
			greeting = strings.ToUpper(greeting)
		}

		fmt.Println(greeting)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//init runs before the command is executed
	rootCmd.Flags().BoolVarP(&shout, "shout", "s", false, "Uppercase")
}
