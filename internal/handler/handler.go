package handler

import (
	"skhaz.dev/summarizer/internal/openai"
	"skhaz.dev/summarizer/internal/twitch"
)

func NewHandler(twitch *twitch.Client, openai *openai.Client) *Handler {
	return &Handler{
		Twitch: twitch,
		OpenAI: openai,
	}
}
