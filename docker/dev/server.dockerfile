FROM golang:1.22.6-alpine

ENV PORT=3000

WORKDIR /app

COPY ./server/go.mod ./server/go.sum ./
RUN go mod download

COPY ./server ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /server

EXPOSE $PORT

CMD ["/server"]