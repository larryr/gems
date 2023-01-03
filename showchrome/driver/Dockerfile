# syntax=docker/dockerfile:1
## deployment image
FROM gcr.io/distroless/base-debian11

WORKDIR /
COPY  showchrome ./showchrome

EXPOSE 7000
USER nonroot:nonroot
ENTRYPOINT ["/showchrome"]

