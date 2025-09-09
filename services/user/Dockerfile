FROM golang:1.22 as builder
WORKDIR /app
COPY . /app
RUN go build -o user .
FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/user /user
EXPOSE 50051
ENTRYPOINT ["/user"]
