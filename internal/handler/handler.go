package handler

import (
	"streamopinion.fun/internal/openai"
	"streamopinion.fun/internal/twitch"
)

func NewHandler(twitch *twitch.Client, openai *openai.Client) *Handler {
	return &Handler{
		Twitch: twitch,
		OpenAI: openai,
	}
}
