package handler

import (
	"embed"
	"net/http"
)

//go:embed public/index.html
var index embed.FS

func (h *Handler) Index(writer http.ResponseWriter, request *http.Request) {
	file, _ := index.ReadFile("public/index.html")

	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.Write(file)
}
