package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

type office_proprties_type struct {
	Forecast string
}

type office_type struct {
	Properties office_proprties_type
}

type forecast_period_type struct {
	Number          int
	StartTime       string
	EndTime         string
	Temperature     int
	TemperatureUnit string
	WindSpeed       string
	WindDirection   string
	ShortForecast   string
}

type forecast_properties_type struct {
	Periods []forecast_period_type
}

type forecast_type struct {
	Properties forecast_properties_type
}

var lookupCmd = &cobra.Command{
	Use:   "lookup",
	Short: "lookup weather information by longitude and latitude",
	Long:  "lookup weather information by longitude and latitude",
	Run: func(cmd *cobra.Command, args []string) {

		lat, _ := cmd.Flags().GetFloat32("lat")
		lon, _ := cmd.Flags().GetFloat32("lon")
		limit, _ := cmd.Flags().GetInt("limit")

		url := fmt.Sprintf("https://api.weather.gov/points/%v,%v", lat, lon)
		office_resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}

		defer office_resp.Body.Close()
		body, err := io.ReadAll(office_resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		office := office_type{}
		err = json.Unmarshal(body, &office)
		if err != nil {
			log.Fatalln(err)
		}

		url = strings.TrimSpace(office.Properties.Forecast)
		if url == "" {
			log.Fatalln("Unable to get the office to query the weather")
		}

		weather_resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}

		defer weather_resp.Body.Close()
		body, err = io.ReadAll(weather_resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		forecast := forecast_type{}
		err = json.Unmarshal(body, &forecast)
		if err != nil {
			log.Fatalln(err)
		}

		periods := forecast.Properties.Periods
		count := len(periods)
		if count == 0 {
			log.Fatalln("No forecast found")
		}

		if limit == 0 || limit > count {
			limit = count
		}

		result, err := json.MarshalIndent(periods[0:limit], "", " ")
		if err != nil {
			log.Fatalln(err)
		}

		cmd.Println(string(result))

	},
}

func init() {

	lookupCmd.Flags().Float32("lon", 0, "The longitude")
	lookupCmd.Flags().Float32("lat", 0, "The latitude")
	lookupCmd.Flags().IntP("limit", "l", 0, "The max number of records to display")

	lookupCmd.MarkFlagRequired("lon")
	lookupCmd.MarkFlagRequired("lat")
}
