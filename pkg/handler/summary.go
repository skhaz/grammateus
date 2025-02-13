package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"skhaz.dev/streamopinion/pkg/openai"
	"skhaz.dev/streamopinion/pkg/twitch"
)

func (h *Handler) Summary(w http.ResponseWriter, r *http.Request) {
	var (
		callback twitch.Callback

		room     = r.URL.Query().Get("room")
		batch, _ = strconv.Atoi(r.URL.Query().Get("batch"))

		count    = 0
		messages = make([]string, 0, batch)
		channel  = make(chan string)

		e = json.NewEncoder(w)
	)

	w.Header().Set("Content-Type", "application/json")

	callback = func(message string) {
		count++
		messages = append(messages, strconv.Quote(message))
		channel <- message
	}

	h.Twitch.Register(room, &callback)

	defer h.Twitch.Unregister(room, &callback)

	for count < batch {
		<-channel
	}

	request := &openai.Request{
		Model: openai.ModelGPT4o,
		Messages: []*openai.Message{
			{Role: openai.RoleSystem, Content: "Take a deep breath"},
			{Role: openai.RoleSystem, Content: "You are a streamer assistant who resumes what their chat is about."},
			{Role: openai.RoleSystem, Content: "Each message will be surrounded by double quotation marks."},
			{Role: openai.RoleSystem, Content: "Your analysis should be as short as possible, without introductions, getting straight to the point."},
			{Role: openai.RoleSystem, Content: "Your analysis MUST be in the same language spoken in the chat. Not in always be in English."},
			{Role: openai.RoleUser, Content: strings.Join(messages, " ")},
		}}

	response, err := h.OpenAI.Do(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		e.Encode(Error{err.Error()})
		return
	}

	if response.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		e.Encode(Error{response.Error.Message})
		return
	}

	w.WriteHeader(http.StatusOK)
	e.Encode(Response{response.Choices[0].Message.Content})
}
