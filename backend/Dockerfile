FROM golang:1.26.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM golang:1.26.2 AS test

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

FROM gcr.io/distroless/static:nonroot

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

ENTRYPOINT ["/app/server"]
