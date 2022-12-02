# Build prover-server
FROM golang:1.18.2-bullseye as base

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./pkg ./pkg

RUN go build -o ./prover ./cmd/prover/prover.go
RUN go build -tags="rapidsnark_noasm" -o ./prover_noasm ./cmd/prover/prover.go


# Main image
FROM alpine:3.16.0

RUN apk add --no-cache libstdc++ gcompat libgomp

COPY --from=base /build/prover /home/app/prover
COPY --from=base /build/prover_noasm /home/app/prover_noasm
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base "/go/pkg/mod/github.com/wasmerio/wasmer-go@v1.0.4/wasmer/packaged/lib/linux-amd64/libwasmer.so" \
"/go/pkg/mod/github.com/wasmerio/wasmer-go@v1.0.4/wasmer/packaged/lib/linux-amd64/libwasmer.so"
COPY docker-entrypoint.sh /usr/local/bin/

COPY ./configs   /home/app/configs
COPY ./circuits  /home/app/circuits

WORKDIR /home/app

# Command to run
ENTRYPOINT ["docker-entrypoint.sh"]

EXPOSE 8002
