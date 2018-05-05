ARG ARCH
FROM multiarch/debian-debootstrap:${ARCH}-stretch

RUN apk add --no-cache ca-certificates

COPY ./main /opt/edunav/main

CMD /opt/edunav/main