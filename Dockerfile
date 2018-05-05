ARG ARCH
FROM multiarch/alpine:${ARCH}-edge

RUN apk add --no-cache ca-certificates

COPY ./main /opt/edunav/main

CMD /opt/edunav/main