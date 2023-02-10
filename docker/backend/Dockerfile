FROM golang:1.19-alpine

WORKDIR /srv/app

COPY ./backend/ /srv/app

RUN go mod download \
  && go get github.com/githubnemo/CompileDaemon \
  && go install github.com/githubnemo/CompileDaemon


ENTRYPOINT ["CompileDaemon"]
# enable src code watching for automatic builds with CompileDaemon
CMD [\
  "-exclude-dir=.git/,bin/",\
  "-build=go build -o ./bin/server ./cmd/server/main.go",\
  "-command=./bin/server",\
  "-polling=true"\
]