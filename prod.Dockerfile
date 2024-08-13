FROM golang:1.22.3-alpine3.19 AS builder

COPY . /app
WORKDIR /app

RUN go mod download
RUN go mod tidy
RUN go build -o ./bin/recipe cmd/recipe/main.go

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/bin/recipe .
COPY config/config.yaml config.yaml
ENV YAML_CONFIG_FILE_PATH=config.yaml
COPY config/.env .env
ENV ENV_CONFIG_FILE_PATH=.env
COPY migrations migrations

CMD ["./recipe"]
