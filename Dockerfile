FROM golang:1.22.2-alpine as builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o glofox main.go

FROM alpine:3.8

EXPOSE 8080

COPY --from=builder /app/glofox .

ENTRYPOINT [ "/glofox" ]
