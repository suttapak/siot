# Stage 1: Build the Go binary
FROM golang:1.19.4-alpine as builder

# Set the working directory
WORKDIR /app

RUN mkdir /app/public/asset/images -p 
# Copy the source files
COPY . .

# Build the Go binary
RUN go build -o main .

# Stage 2: Create the final image
FROM alpine:latest

# Copy the Go binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/config.yaml .


# Set the command to run when the container starts
CMD ["./main"]
