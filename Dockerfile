FROM golang:1.18-bullseye as builder

ENV APP_HOME /go/src/server

WORKDIR ${APP_HOME}

COPY ./src .

RUN go mod download
RUN go mod verify
RUN go build -o server

FROM golang:1.18-bullseye

ENV APP_HOME "/go/src/server"
ENV PATH_CLONE "_clone"
ENV PATH_PULL "_pull"
ENV PATH_VERSION "_version"
ENV PATH_WEBHOOK "_hook"
ENV PATH_HEALTH "_health"
ENV REPO_BRANCH "main"
ENV REPO_TARGET_FOLDER "target-git"
ENV REPO_URL ""
ENV HTTP_PORT 8081

RUN mkdir -p ${APP_HOME}
WORKDIR ${APP_HOME}

COPY --from=builder "$APP_HOME"/server $APP_HOME

CMD ["./server"]