FROM golang:1.19 as builder

RUN mkdir -p /build
WORKDIR /build

COPY . .

RUN go mod download

ENV CGO_ENABLED=0
# Compile static binary
RUN go build -ldflags "-s -w" ./cmd/api

##################################################
# Final image
FROM alpine

RUN mkdir -p /app/config
WORKDIR /app

ENV CONFIGOR_DIR /app/config

# Copy config
COPY --from=builder /build/config/config.yaml config/config.yaml

# Copy the binary from the previous stage
COPY --from=builder /build/api api

CMD ["/app/api"]
