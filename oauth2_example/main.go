package main

import (
	"log"
	"os"

	"song.com/go_get_started/oauth2_example/auth"
)

func main() {

	what := "auth"
	args := os.Args
	if len(args) > 1 {
		what = args[1]
	}

	switch what {
	case "auth":
		auth.AuthByMicrosoft()
	case "verify":
		auth.VerifyByMicrosoft()
	default:
		log.Println("Please tell me what to do")
	}

}
