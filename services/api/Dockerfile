FROM golang:1.13.0 AS gobuild

ENV GO111_MODULE=1
WORKDIR src/github.com/layer-devops/livechat-example
COPY go.mod go.sum ./
RUN go mod download
COPY pkg ./pkg
COPY main.go ./main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o /api

FROM ubuntu:18.04
RUN apt-get update && \
    rm -rf /var/lib/apt/lists/*

COPY --from=gobuild /api /api
ENV PORT=80
EXPOSE 80
CMD ["/api"]