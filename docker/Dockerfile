FROM golang:tip-alpine3.22 AS build

COPY go.mod main.go ./

COPY ip ./ip

COPY cloudflare ./cloudflare

COPY configuration ./configuration

RUN go build

FROM alpine:3.22.0

WORKDIR /home

ENV CDU_CONFIGURATION_FILE=/home/config.json

ENV CDU_EXEC_INTERVAL=60

COPY --from=build /go/cloudflare-dns-updater .

COPY docker/entrypoint.sh .

RUN chmod +x entrypoint.sh cloudflare-dns-updater

ENTRYPOINT ["sh", "-c"]

CMD ["/home/entrypoint.sh"]
