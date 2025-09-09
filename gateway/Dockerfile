FROM golang:1.22 as builder
WORKDIR /app
COPY . /app
RUN go build -o gateway .
FROM gcr.io/distroless/base-debian12
WORKDIR /
COPY --from=builder /app/gateway /gateway
EXPOSE 8080
ENTRYPOINT ["/gateway"]
