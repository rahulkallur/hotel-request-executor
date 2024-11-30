FROM golang:alpine AS builder

WORKDIR /app 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./request_executor

# Final stage
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/request_executor /usr/bin/

EXPOSE 8000

ENTRYPOINT ["request_executor"]
