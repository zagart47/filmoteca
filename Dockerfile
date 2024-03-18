FROM golang:1.22.1-alpine as builder
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY .. .
RUN env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main cmd/filmoteca/main.go

FROM scratch
COPY internal/config/config.yaml internal/config/
COPY --from=builder main /main
EXPOSE 8080
ENTRYPOINT ["/main"]
