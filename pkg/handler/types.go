package handler

import (
	"skhaz.dev/streamopinion/pkg/openai"
	"skhaz.dev/streamopinion/pkg/twitch"
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
