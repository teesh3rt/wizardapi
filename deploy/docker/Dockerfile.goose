FROM golang:1.25 AS builder

ENV CGO_ENABLED=0
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

FROM gcr.io/distroless/static AS runner

WORKDIR /migrations
COPY --from=builder /go/bin/goose /usr/local/bin/goose
COPY ./migrations .
ENTRYPOINT ["goose", "up"]