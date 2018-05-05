ARG ARCH
FROM multiarch/debian-debootstrap:${ARCH}-stretch

COPY ./main /opt/edunav/main

CMD /opt/edunav/main