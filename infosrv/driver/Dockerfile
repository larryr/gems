# syntax=docker/dockerfile:1
## deployment image
FROM gcr.io/distroless/base-debian11

WORKDIR /
COPY  infosrv ./infosrv

EXPOSE 9999
USER nonroot:nonroot
ENTRYPOINT ["/infosrv"]

