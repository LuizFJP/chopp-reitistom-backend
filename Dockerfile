FROM golang:1.24.6 AS builder

ARG MAIN_PACKAGE=./init
ARG VERSION=dev
ARG COMMIT=none
ARG BUILD_TIME=unknown

WORKDIR /

RUN apt-get update && apt-get install -y --no-install-recommends \
    && rm -rf /var/lib/apt/lists/*

ENV GOBIN=/usr/local/bin

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://proxy.golang.org,direct \
    GOSUMDB=sum.golang.org

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -trimpath \
    -ldflags "-s -w -X main.version=${VERSION} -X main.commit=${COMMIT} -X main.buildTime=${BUILD_TIME}" \
    -o /app ${MAIN_PACKAGE}

FROM gcr.io/distroless/static:nonroot

WORKDIR /
COPY --from=builder /app /app

EXPOSE 50051 2112
USER nonroot:nonroot
ENTRYPOINT ["/app"]