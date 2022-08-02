package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"song.com/go_get_started/cli_cobra_example/cmd/weather"
)

var RootCmd = &cobra.Command{
	Use:   "poc",
	Short: "A command-line utility for geo location and weather lookup",
	Long:  "A command-line utility for geo location and weather lookup",
}

func Execute() {
	args := os.Args
	if len(args) == 2 {
		arg0 := args[1]
		if arg0 == "-v" || arg0 == "--version" {
			fmt.Println("poc cli version 1.0.0")
			os.Exit(0)
		}
	}

	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	RootCmd.Flags().BoolP("version", "v", false, "Show the poc version")
	// rootCmd.CompletionOptions.DisableDefaultCmd = true

	RootCmd.AddCommand(loginCmd)
	RootCmd.AddCommand(logoutCmd)
	RootCmd.AddCommand(weather.WeatherRootCmd)
}
