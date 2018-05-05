ARG ARCH
FROM multiarch/debian-deboostrap:${ARCH}-stretch

RUN apk add --no-cache ca-certificates

COPY ./main /opt/edunav/main

CMD /opt/edunav/main