ARG GOVER=1.13.15
ARG GOOS=linux
ARG GOARCH=amd64

FROM golang:${GOVER}

RUN mkdir /target
RUN mkdir /build
WORKDIR /build

COPY go.mod .
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

# COPY the source code as the last step
COPY . .

RUN chmod +x copy.sh

# Build the binary
RUN CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} \
    go build -o imds-debug \
             -ldflags "-X main.goVersion=$(go version | cut -d " " -f 3)" ./main.go

ENTRYPOINT ["/build/copy.sh"]


