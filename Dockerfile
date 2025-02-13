FROM golang:1.23-bullseye
WORKDIR /opt
COPY . .
RUN CGO_ENABLED=0 go build -o app

FROM gcr.io/distroless/static-debian12
COPY --from=0 /opt/app /
CMD ["/app"]
