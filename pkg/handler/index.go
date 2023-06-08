package handler

import (
	"embed"
	"net/http"
)

//go:embed public/index.html
var index embed.FS

func (h *Handler) Index(writer http.ResponseWriter, request *http.Request) {
	bytes, _ := index.ReadFile("public/index.html")

	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.Write(bytes)
}

func (h *Handler) Icon(writer http.ResponseWriter, request *http.Request) {
	bytes := make([]byte, 0)

	writer.Header().Set("Content-Type", "image/x-icon")
	writer.Write(bytes)
}
