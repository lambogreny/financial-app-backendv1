FROM golang:1.13.4 as builder
ENV DATA_DIRETORY /go/src/financial-app-backend
WORKDIR $DATA_DIRETORY
ARG APP_VERSION
ARG CGO_ENABLED=0
COPY . . 
RUN go build -ldflags="-X financial-app-backend/internal/config.Version=$APP_VERSION" financial-app-backend/cmd/server

FROM alpine:3.10
ENV DATA_DIRETORY /go/src/financial-app-backend/
RUN apk add --update --no-cache \
    ca-certificates
COPY internal/database/migrations ${DATA_DIRETORY}internal/database/migrations
COPY --from=builder ${DATA_DIRETORY}server /financial-app-backend

ENTRYPOINT [ "/financial-app-backend" ]