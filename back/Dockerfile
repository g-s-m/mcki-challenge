
FROM golang:1.13 AS builder
WORKDIR /build
COPY . /build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"'

FROM alpine
WORKDIR /mcki-back
COPY --from=builder /build/back /mcki-back/back
COPY --from=builder /build/reports.json /mcki-back/reports.json
ENTRYPOINT ["/mcki-back/back"]

