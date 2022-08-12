package weather

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var echoCmd = &cobra.Command{
	Use:   "echo text",
	Short: "Echo the input text after an HTTP POST",
	Long:  "Echo the input text after an HTTP POST",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		text := args[0]

		url := "http://API.SHOUTCLOUD.IO/V1/SHOUT"
		payload := []byte(fmt.Sprintf(`{"INPUT": "%s"}`, text))

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		if err != nil {
			log.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		cmd.Println("Response received:")
		cmd.Println(string(body))
	},
}

func init() {}
