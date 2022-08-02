package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"song.com/go_get_started/cli_cobra_example/utils"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logout from the weather app",
	Long:  "Logout from the weather app",
	Run: func(cmd *cobra.Command, args []string) {

		file_name := "login"

		data, err := utils.Read_file_from_home(file_name)
		if err != nil {
			log.Fatal("You are not logged in at this time.")
		}

		err = utils.Remove_file_from_home(file_name)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Bye", string(data))
		fmt.Println("You are logged out")
	},
}
