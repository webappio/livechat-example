FROM golang:1.13.0 AS gobuild

LABEL maintainer="Colin Chartier <colin@layerci.com>"
ENV GO111MODULE=on
ENV GODEBUG="mlock=0"

ENV TAG="v4.5.0"

WORKDIR /go/src/github.com/golang-migrate/migrate
RUN git clone https://github.com/golang-migrate/migrate.git . && git checkout ${TAG}
RUN go build -o /migrate -ldflags="-s -w -X main.version=${TAG}" -tags "postgres file" ./cmd/migrate

FROM ubuntu:18.04
COPY --from=gobuild /migrate /migrate
COPY migrations /migrations

# Run migrate (CMD) and then sleep forever (so that we can put this in a deployment & use apply)
# (bash parses first arg as $0, so "dummy" lets $@ be replaced by arg)
ENTRYPOINT [ \
    "bash", "-c", \
    "/migrate \
        -source file:///migrations \
        -database postgres://postgresql:postgresql@postgresql:5432/livechat-example?sslmode=disable \
        $@ && while true; do sleep 9999; done", \
    "dummy" \
]
