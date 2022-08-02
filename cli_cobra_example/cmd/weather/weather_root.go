package weather

import (
	"github.com/spf13/cobra"
)

var WeatherRootCmd = &cobra.Command{
	Use:   "weather",
	Short: "Look up the weather and geo location",
	Long:  "Look up the weather and geo location",
}

func init() {

	WeatherRootCmd.AddCommand(geoCmd)
	WeatherRootCmd.AddCommand(lookupCmd)
	WeatherRootCmd.AddCommand(echoCmd)
}
