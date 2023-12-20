FROM golang:1.18 as builder

WORKDIR /ci

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" go build -o cpr .

FROM ubuntu:22.04

RUN mkdir /app
WORKDIR /app

COPY --from=builder /ci/cpr /app/cpr

ENTRYPOINT ["/app/cpr"]
