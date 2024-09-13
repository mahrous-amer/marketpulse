FROM golang:alpine AS builder

WORKDIR /marketpulse

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o marketpulse

FROM scratch

COPY --from=builder /marketpulse /marketpulse

ENTRYPOINT ["/marketpulse"]
