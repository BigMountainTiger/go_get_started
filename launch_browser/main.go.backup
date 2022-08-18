package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"
)

var url string = "https://www.google.com"

func launch(command string) {

	cmd := exec.Command(command, url)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)
	fmt.Println(cmd.Process.Pid)

}

func main() {
	os := runtime.GOOS

	switch os {
	case "linux":
		launch("xdg-open")
	case "windows":
		launch("start")
	case "darwin":
		launch("open")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("It does not look to have a all ideal solution at this time")
}
