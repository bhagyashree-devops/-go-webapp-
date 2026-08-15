# ---------- Builder Stage ----------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy dependency files first (for caching)
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY main.go .

# Build the binary. CGO_ENABLED=0 makes it static so it can run in a distroless image.
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server .

# ---------- Runtime Stage ----------
# Using distroless means there is no shell, no package manager, nothing for hackers to exploit.
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .

# Run as a non-root user for security
USER nonroot:nonroot

EXPOSE 8080

ENTRYPOINT ["./server"]
