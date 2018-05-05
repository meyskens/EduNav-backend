ARG ARCH
FROM multiarch/debia-deboostrap:${ARCH}-stretch

RUN apk add --no-cache ca-certificates

COPY ./main /opt/edunav/main

CMD /opt/edunav/main