# Stage 1: Build stage
from golang:1.24-alpine as builder

RUN mkdir -p /app
WORKDIR /app

COPY . .

RUN go mod tidy && go mod download && go mod vendor

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

# Stage 2: Final stage
FROM alpine:edge


# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/myapp .

# Set the timezone and install CA certificates
RUN apk --no-cache add ca-certificates tzdata

# Set the entrypoint command
ENTRYPOINT ["/app/myapp"]