.SILENT:

.PHONY: fmt
fmt:
	npx prettier --write --ignore-unknown --ignore-path .gitignore --print-width=120 pkg/handler/public/index.html
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test ./...
