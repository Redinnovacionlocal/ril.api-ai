FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
     go build -mod=mod -o app ./cmd/api

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /app/app /app/app

ENV PORT=8080
EXPOSE 8080

USER nonroot:nonroot
CMD ["/app/app"]
