ARG ARCH
FROM multiarch/debian-debootstrap:${ARCH}-stretch

COPY ./main /opt/edunav/main

WORKDIR /opt/edunav/

CMD /opt/edunav/main