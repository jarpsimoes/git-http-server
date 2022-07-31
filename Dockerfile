FROM golang:1.18-bullseye as builder

ENV APP_HOME /go/src/server

WORKDIR "$APP_HOME"

COPY ./src .

RUN go mod download
RUN go mod verify
RUN go build -o server

FROM golang:1.18-bullseye

ENV APP_HOME /go/src/server
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY --from=builder "$APP_HOME"/server $APP_HOME

EXPOSE 8081

CMD ["./server"]