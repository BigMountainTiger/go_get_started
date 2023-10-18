package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"song.com/go_get_started/launch_chromium/browser"
)

func main() {
	browser_channel := make(chan error)
	sigint_channel := make(chan os.Signal, 1)
	signal.Notify(sigint_channel, os.Interrupt)
	defer signal.Stop(sigint_channel)

	fmt.Println("Please wait for the borwser to launch ...")

	chromium := &browser.Chromium{
		Embeded:        embeded[:],
		App_built_time: build_time,
	}

	err := chromium.Launch("http://www.yahoo.com")
	if err != nil {
		log.Fatalln("failed to start the browser", err)
	}
	defer chromium.Close()

	go func() {
		browser_channel <- chromium.Wait_for_terminate()
	}()

	var comment string
	select {
	case err := <-browser_channel:
		if err != nil {
			fmt.Println(err)
		}
		comment = "browser terminated"
	case <-sigint_channel:
		comment = "ctl+c interrupt"
	case <-time.After(30 * time.Second):
		comment = "timeout"
	}

	fmt.Println(comment)

}
