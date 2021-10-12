##
## Build
##
FROM golang:1.16-alpine as base

#ENV GOFLAGS="-mod=vendor"
#ENV CGO_ENABLED=1

WORKDIR /build

RUN apk add --no-cache --update git

COPY go.mod ./
COPY go.sum ./
RUN go mod download
#RUN go mod vendor

#COPY . .
COPY ./cmd ./cmd
COPY ./pkg ./pkg

RUN CGO_ENABLED=0 go build -o ./prover ./cmd/prover/prover.go


# Build an prover image
FROM node:14

ENV APP_USER=app
ENV APP_UID=1001
ENV DOCKER_GID=999

RUN apt-get update && apt-get install -y ca-certificates
RUN mkdir -p /home/$APP_USER
RUN adduser -u $APP_UID $APP_USER && chown -R $APP_USER:$APP_USER /home/$APP_USER
RUN addgroup --system --gid ${DOCKER_GID} docker
RUN addgroup ${APP_USER} docker
RUN mkdir -p /app && chown -R $APP_USER:$APP_USER /app
RUN rm rm -rf /var/lib/apt/lists/*

USER app:app

ENV NPM_CONFIG_PREFIX=/app/node/.npm-global

RUN node -v
RUN npm -v

RUN npm install -g circom@latest
RUN npm install -g snarkjs@latest

RUN export PATH=$PATH:/app/node/.npm-global/bin/

COPY ./configs      /app/configs
COPY ./circuits     /app/circuits
COPY --from=base /build/prover /app/prover
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app

# Command to run
ENTRYPOINT ["/app/prover"]