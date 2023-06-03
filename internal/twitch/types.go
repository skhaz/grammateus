package twitch

import (
	"crypto/tls"
	"sync"
)

type Callback func(string)

type Client struct {
	mu        sync.Mutex
	user      string
	token     string
	callbacks map[string][]*Callback
	conn      *tls.Conn
}
