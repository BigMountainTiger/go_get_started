package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"song.com/go_get_started/cli_cobra_example/utils"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the weather app",
	Long:  "Login to the weather app",
	Run: func(cmd *cobra.Command, args []string) {

		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")

		if password == "" {
			cmd.Println("Password is required")
		}

		err := utils.Save_file_to_home("login", []byte(user))
		if err != nil {
			log.Fatal(err)
		}

		cmd.Println("Welcome", user)
	},
}

func init() {

	loginCmd.Flags().StringP("user", "u", "", "The user name")
	loginCmd.Flags().StringP("password", "p", "", "The password")

	loginCmd.MarkFlagRequired("user")
	loginCmd.MarkFlagRequired("password")

}
