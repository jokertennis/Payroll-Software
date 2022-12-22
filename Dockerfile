FROM golang:1.19

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change.
# go mod verify: Check that dependencies of the main module stored in the module cache have not been modified since they were downloaded.
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# reference document of "go build" : https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/01.3.html
# When is "/usr/local/bin/app" is declared?  The answer of this question is that /etc/profile file is executed when logging into the shell.
# https://zenn.dev/shinji/articles/0eaad9d8c0954f
COPY . .
RUN go build ./main.go

CMD ["go", "run", "main.go"]