# Build prover-server
FROM golang:1.17-alpine as base

WORKDIR /build

RUN apk add --no-cache --update git

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./pkg ./pkg

RUN CGO_ENABLED=0 go build -o ./prover ./cmd/prover/prover.go


# Main image
FROM node:16

ENV APP_USER=app
ENV APP_UID=1001

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

RUN adduser -u $APP_UID $APP_USER --disabled-password --gecos "First Last,RoomNumber,WorkPhone,HomePhone"

ENV NPM_CONFIG_PREFIX=/home/app/node/.npm-global
RUN npm install -g snarkjs@latest

ENV PATH=${PATH}:/home/app/node/.npm-global/bin

COPY --from=base /build/prover /home/app/prover
#COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY ./configs   /home/app/configs
COPY ./circuits  /home/app/circuits
COPY ./js        /home/app/js

RUN chown -R $APP_USER:$APP_USER /home/app

USER app:app
WORKDIR /home/app

# Command to run
ENTRYPOINT ["/home/app/prover"]

EXPOSE 8002
