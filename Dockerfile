FROM golang:1.18

WORKDIR /usr/src/gqlgen-todos

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY gqlgen.yml server.go tools.go ./
COPY graph ./graph

# CMD [ "go", "run", "server.go" ]