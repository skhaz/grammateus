package twitch

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/textproto"
	"reflect"
	"strings"
	"time"
)

const Endpoint = "irc.chat.twitch.tv:6697"

func NewClient(user string, token string) *Client {
	var (
		callbacks = make(map[string][]*Callback)
	)

	return &Client{
		user:      user,
		token:     token,
		callbacks: callbacks,
	}
}

func (c *Client) Start() {
	var (
		err  error
		conf = &tls.Config{MinVersion: tls.VersionTLS12}
	)

	dialer := &net.Dialer{
		KeepAlive: time.Second * 30,
	}

	c.conn, err = tls.DialWithDialer(dialer, "tcp", Endpoint, conf)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(c.conn, "PASS %s\r\n", c.token)
	fmt.Fprintf(c.conn, "NICK %s\r\n", c.user)

	go c.polling()
}

func (c *Client) polling() {
	tp := textproto.NewReader(bufio.NewReader(c.conn))

	for {
		line, err := tp.ReadLine()
		if err != nil {
			panic(err)
		}

		fmt.Println(line)

		if strings.HasPrefix(line, "PING") {
			go handlePing(c.conn, line)
			continue
		}

		if !strings.Contains(line, "PRIVMSG") {
			continue
		}

		room, message := Parse(line)

		for _, callback := range c.callbacks[room] {
			go processMessage(message, *callback)
		}
	}
}

func (c *Client) Register(room string, callback *Callback) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.callbacks[room] = append(c.callbacks[room], callback)
	fmt.Fprintf(c.conn, "JOIN #%s\r\n", room)
}

func (c *Client) Unregister(room string, callback *Callback) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if callbacks, ok := c.callbacks[room]; ok {
		for i, inner := range callbacks {
			if inner == callback {
				c.callbacks[room] = append(callbacks[:i], callbacks[i+1:]...)
				break
			}
		}
	}

	if len(c.callbacks[room]) == 0 {
		fmt.Fprintf(c.conn, "PART #%s\r\n", room)
	}
}

func Parse(line string) (room, message string) {
	index := strings.Index(line, "#")
	if index != -1 {
		space := strings.Index(line[index:], " ")
		if space != -1 {
			room = line[index+1 : index+space]
		}
	}

	index = strings.Index(line, " :")
	if index != -1 {
		message = line[index+2:]
	}

	return room, message
}

func handlePing(writer io.Writer, line string) {
	fmt.Fprintf(writer, "PONG %s\r\n", line[5:])
}

func processMessage(message string, callback Callback) {
	if callback != nil && reflect.TypeOf(callback).Kind() == reflect.Func {
		callback(message)
	}
}
