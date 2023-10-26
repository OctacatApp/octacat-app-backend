# builder stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o build/main src/cmd/main.go

# app stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/build /app/build  
COPY --from=builder /app/etc /app/etc
CMD [ "./build/main", "-env", "prod" ]