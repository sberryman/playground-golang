FROM golang:1.18.3-bullseye

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y \
      ca-certificates \
      gettext \
      curl \
      gettext-base \
      bash

ENV GONOSUMDB "go.bird.co/*"
ENV GOPROXY "https://artifactory.svc.bird.co/artifactory/api/go/golang"

# app dir
WORKDIR /app

# install 8 tools
RUN go install github.com/cweill/gotests/gotests@latest && \
    go install github.com/fatih/gomodifytags@latest && \
    go install github.com/josharian/impl@latest && \
    go install github.com/haya14busa/goplay/cmd/goplay@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install golang.org/x/tools/gopls@latest && \
    go install golang.org/x/tools/cmd/goimports@latest