FROM golang:alpine AS builder


# Move to working directory /build
WORKDIR /build

COPY main.go .
COPY fakeDataGenerator/ ./fakeDataGenerator/

# generates module
RUN go mod init main
RUN go mod tidy

# Build the application
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build

# Build a small image using distroless from google
FROM gcr.io/distroless/base-debian10

COPY --from=builder /build/main .

# Command to run

CMD ["/main","--generator-throttling-rate=0","--handler-throttling-rate=0","--channel-size=10000","--is-gzip=false"]
