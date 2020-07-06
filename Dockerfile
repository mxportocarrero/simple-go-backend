FROM golang:1.14.4 as builder
ENV DATA_DIRECTORY /go/src/simple-go-backend
WORKDIR $DATA_DIRECTORY
ARG APP_VERSION
ARG CGO_ENABLED=0
COPY . .
RUN go build -ldflags="-X simple-go-backend/internal/config.Version=$APP_VERSION" simple-go-backend/cmd/server

FROM alpine:3.10
ENV DATA_DIRECTORY /go/src/simple-go-backend
RUN apk add --update --no-cache \
  ca-certificates
COPY --from=builder ${DATA_DIRECTORY}/server /simple-go-backend

ENTRYPOINT ["/simple-go-backend"]