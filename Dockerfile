FROM golang:1.18.1
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/cespare/reflex@latest
CMD reflex -g '*.go' go run . --start-service

