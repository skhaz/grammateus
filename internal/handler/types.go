package handler

import (
	"skhaz.dev/streamopinion/internal/openai"
	"skhaz.dev/streamopinion/internal/twitch"
)

type Handler struct {
	Twitch *twitch.Client
	OpenAI *openai.Client
}

type Error struct {
	Error string `json:"error"`
}

type Response struct {
	Result string `json:"result"`
}
