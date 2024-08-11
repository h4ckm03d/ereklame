FROM golang:1.22-alpine as builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main cmd/api/main.go
FROM alpine:3
RUN apk --no-cache add ca-certificates
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]