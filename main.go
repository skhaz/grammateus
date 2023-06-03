package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"skhaz.dev/summarizer/internal/handler"
	"skhaz.dev/summarizer/internal/openai"
	"skhaz.dev/summarizer/internal/twitch"
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
