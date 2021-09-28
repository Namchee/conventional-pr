FROM golang:1.16

WORKDIR /ci

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN  go build -a -o ethos .

RUN chmod +x /ci/ethos

CMD ["/ci/ethos"]
