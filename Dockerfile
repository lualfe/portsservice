FROM golang:1.19-alpine AS builder

WORKDIR /go/src/portsservice

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/app cmd/portsservice/main.go

FROM alpine:3.17.0

RUN apk add ca-certificates

WORKDIR /portsservice

COPY --from=builder /go/src/portsservice/out/app .
COPY --from=builder /go/src/portsservice/internal/usecase/testdata/ports.json .

CMD ["./app"]
