package handler

import (
	"streamopinion.fun/internal/openai"
	"streamopinion.fun/internal/twitch"
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
