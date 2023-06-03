package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"streamopinion.fun/internal/handler"
	"streamopinion.fun/internal/openai"
	"streamopinion.fun/internal/twitch"
)

func main() {
	var (
		twitch  = twitch.NewClient(os.Getenv("TWITCH_USER"), os.Getenv("TWITCH_TOKEN"))
		openai  = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
		handler = handler.NewHandler(twitch, openai)
	)

	twitch.Start()

	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/summary", handler.Summary)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}
