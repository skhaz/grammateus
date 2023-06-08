package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"skhaz.dev/streamopinion/pkg/handler"
	"skhaz.dev/streamopinion/pkg/openai"
	"skhaz.dev/streamopinion/pkg/twitch"
)

func main() {
	var (
		twitch  = twitch.NewClient(os.Getenv("TWITCH_USER"), os.Getenv("TWITCH_TOKEN"))
		openai  = openai.NewClient(os.Getenv("OPENAI_API_KEY"))
		handler = handler.NewHandler(twitch, openai)
	)

	twitch.Start()

	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/favicon.ico", handler.Icon)
	http.HandleFunc("/summary", handler.Summary)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), nil))
}
