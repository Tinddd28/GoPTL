FROM golang:1.23-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc musl-dev

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# build
COPY . ./

RUN go build -o ./bin/app cmd/main.go

FROM alpine AS runner
RUN apk --no-cache add bash postgresql-client

COPY --from=builder /usr/local/src/bin/app /app/app
COPY --from=builder /go/bin/migrate /app/migrate

COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

COPY .env /.env
COPY config/dev.yaml /config/dev.yaml
COPY internal/migrations /app/migrations

RUN ls -la /app
ENTRYPOINT ["/app/entrypoint.sh"]