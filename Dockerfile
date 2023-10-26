FROM golang:1.18

WORKDIR /ci

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN  go build -a -o cpr .

RUN chmod +x /ci/cpr

CMD ["/ci/cpr"]
