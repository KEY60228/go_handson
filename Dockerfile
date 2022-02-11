FROM golang:1.17-alpine3.14

RUN apk add --update --no-cache alpine-sdk build-base bash gcc make git curl postgresql-client gzip tar \
    libx11-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev linux-headers mesa-dev
RUN go install golang.org/x/tools/cmd/goimports@latest \
    && go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest \
    && go install github.com/ramya-rao-a/go-outline@latest \
    && go install github.com/cweill/gotests/gotests@latest \
    && go install github.com/fatih/gomodifytags@latest \
    && go install github.com/josharian/impl@latest \
    && go install github.com/haya14busa/goplay/cmd/goplay@latest \
    && go install honnef.co/go/tools/cmd/staticcheck@latest \
    && go install golang.org/x/tools/gopls@latest
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.1/migrate.linux-amd64.tar.gz | tar xvz
RUN mv ./migrate /usr/bin/migrate

WORKDIR /go/go_handson
