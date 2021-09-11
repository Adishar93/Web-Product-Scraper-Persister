# syntax=docker/dockerfile:1

FROM golang:1.17.1-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /go-amascr
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /go-amascr /go-amascr

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/go-amascr"]