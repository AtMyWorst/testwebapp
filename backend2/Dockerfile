FROM golang:1.12-rc-alpine as build

RUN apk add --no-cache git=2.20.1-r0
RUN apk add --no-cache ca-certificates

WORKDIR /src
COPY go.mod go.sum .env *.go ./
RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -ldflags "-s -w"

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/webml /webml
COPY --from=build /src/.env /.env

ENTRYPOINT ["/webml"]
