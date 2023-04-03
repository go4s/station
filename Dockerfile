FROM ghcr.io/zhangli2946/images:builder-go1.19.7-alpine3.17                                 AS Builder

WORKDIR /go/src
ADD . .

ARG MODULE_PATH
ARG MODULE_NAME
ENV MODULE_PATH=$MODULE_PATH
ENV MODULE_NAME=$MODULE_NAME
ENV GO111MODULE="on"
ENV CGO_ENABLE=1

RUN make

FROM alpine:3.17                                                                            AS Final
LABEL org.opencontainers.image.source="http://github.com/go4s/station"

WORKDIR /var/app

ARG MODULE_NAME
ENV MODULE_NAME=$MODULE_NAME

COPY --from=Builder                               /go/src/bootstrap.sh                      .
COPY --from=Builder                               /go/bin/$MODULEA_NAME                     .

ENTRYPOINT ["/var/app/bootstrap.sh"]