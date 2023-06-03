package handler

import (
	"skhaz.dev/streamopinion/pkg/openai"
	"skhaz.dev/streamopinion/pkg/twitch"
)

func NewHandler(twitch *twitch.Client, openai *openai.Client) *Handler {
	return &Handler{
		Twitch: twitch,
		OpenAI: openai,
	}
}
