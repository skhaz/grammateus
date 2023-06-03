package twitch_test

import (
	"testing"

	"skhaz.dev/streamopinion/pkg/twitch"
)

func TestParse(t *testing.T) {
	line := "#room :message"
	expectedRoom := "room"
	expectedMessage := "message"

	room, message := twitch.Parse(line)

	if room != expectedRoom {
		t.Errorf("Expected room name is %s, but got %s", expectedRoom, room)
	}

	if message != expectedMessage {
		t.Errorf("Expected message is %s, but got %s", expectedMessage, message)
	}
}
