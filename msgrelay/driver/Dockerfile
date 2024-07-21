# syntax=docker/dockerfile:1
## deployment image
FROM gcr.io/distroless/base-debian11

WORKDIR /
COPY  simp ./simp

EXPOSE 7000
USER nonroot:nonroot
ENTRYPOINT ["/simp"]

