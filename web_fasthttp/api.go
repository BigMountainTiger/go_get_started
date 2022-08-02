package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/valyala/fasthttp"
)

func get_item(ctx *fasthttp.RequestCtx) {

	authorization := string(ctx.Request.Header.Peek("authorization"))
	_, err := authenticate(authorization)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusUnauthorized)
		return
	}

	entry, err := get_dynamo_entry()
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	result, err := json.Marshal(entry)
	if err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		return
	}

	ctx.SetStatusCode(200)
	ctx.SetBody([]byte(result))
}

func main() {

	ok := init_auth()
	if !ok {
		log.Fatal("Unable to init auth capability")
	}

	m := func(ctx *fasthttp.RequestCtx) {

		switch string(ctx.Path()) {
		case "/get-item":
			get_item(ctx)
		case "/profile/healthcheck":
			ctx.SetStatusCode(200)
		default:
			ctx.Error("not found", fasthttp.StatusNotFound)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	fasthttp.ListenAndServe(":"+port, m)
	log.Println("Server started at port", port)
}
