##########################
## builder image
##########################

FROM golang:1.21.5 as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

ENV GOOS linux
ENV GOARCH ${GOARCH:-amd64}
ENV CGO_ENABLED=0

RUN go build -v -o zpe-api cmd/zpe-api/main.go


##########################
## user creator
##########################
FROM alpine:latest as user

ENV APP_HOME /build
ENV APP_USER zpe
ENV APP_GROUP zpe

RUN addgroup -S ${APP_GROUP} && adduser -S ${APP_USER} -G ${APP_GROUP}  --no-create-home
RUN apk --no-cache add ca-certificates \
    && update-ca-certificates

COPY --from=builder ${APP_HOME}/zpe-api ${APP_HOME}/zpe-api

RUN chown ${APP_USER}:${APP_GROUP} ${APP_HOME}/zpe-api \
    /etc/ssl/certs/ca-certificates.crt

RUN sed 's/ash/nologin/g' /etc/passwd


################################
## generate clean, final image
################################ 
FROM alpine:3.14

ENV APP_HOME /build
ENV APP_USER evoluinfo

ARG VERSION
ENV APP_VERSION=$VERSION
ENV GIN_MODE=release

COPY --from=user ${APP_HOME}/zpe-api ${APP_HOME}/zpe-api
COPY --from=user /etc/passwd /etc/passwd
COPY --from=user /etc/group /etc/group
COPY --from=user /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=builder /build/zpe-api /usr/local/bin/zpe-api

USER $APP_USER
WORKDIR $APP_HOME

EXPOSE 8080

ENTRYPOINT [ "zpe-api" ]
