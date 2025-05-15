FROM golang:1.22-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o user-master-service ./cmd/main.go

FROM scratch
COPY --from=builder /app/user-master-service /user-master-service
ENTRYPOINT ["/user-master-service"]
