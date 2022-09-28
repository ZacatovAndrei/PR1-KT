FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod .
COPY *.go ./
COPY menu.json .

RUN go build -o Kitchen

EXPOSE 8087

ENTRYPOINT ["/app/Kitchen"]
