FROM golang:1.24-alpine3.2.1 AS modules
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

FROM golang:1.24-alpine3.2.1 AS builder
COPY --from=modules /go/pkg /go/pkg
COPY ../database /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags migrate -o /bin/app ./app/main

FROM scratch

EXPOSE 3000

#COPY --from=builder /database/migrations /database/migrations
#COPY --from=builder /bin/app /app
#CMD ["/app"]