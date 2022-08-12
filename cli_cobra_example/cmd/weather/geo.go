package weather

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"song.com/go_get_started/cli_cobra_example/utils"
)

var appid string

var geoCmd = &cobra.Command{
	Use:   "geo city-name",
	Short: "lookup geo location information by city, state, and county",
	Long:  "lookup geo location information by city, state, and county",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		city := args[0]
		state, _ := cmd.Flags().GetString("state")
		county, _ := cmd.Flags().GetString("county")
		limit, _ := cmd.Flags().GetInt("limit")

		template := "http://api.openweathermap.org/geo/1.0/direct?q=%s,%s,%s&limit=%v&appid=%s"
		url := fmt.Sprintf(template, city, state, county, limit, appid)

		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		r, err := utils.PrettyString(string(body))
		if err != nil {
			log.Fatalln(err)
		}

		cmd.Println(r)

	},
}

func init() {
	geoCmd.Flags().StringP("state", "s", "", "The state code (optional). For example, MA, PA, TX, etc.")
	geoCmd.Flags().String("county", "us", "The country code")
	geoCmd.Flags().IntP("limit", "l", 5, "Max number of search result to return")

	geoCmd.MarkFlagRequired("city")
}
