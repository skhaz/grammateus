package handler

import (
	"skhaz.dev/streamopinion/internal/openai"
	"skhaz.dev/streamopinion/internal/twitch"
)

func NewHandler(twitch *twitch.Client, openai *openai.Client) *Handler {
	return &Handler{
		Twitch: twitch,
		OpenAI: openai,
	}
}
