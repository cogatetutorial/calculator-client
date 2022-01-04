FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/example/calculator-client/
COPY . .

RUN go get .

RUN CGO_ENABLED=0 go build -o /go/bin/calculator-client

FROM scratch
COPY --from=builder /go/bin/calculator-client /go/bin/calculator-client

ENTRYPOINT ["/go/bin/calculator-client"]
