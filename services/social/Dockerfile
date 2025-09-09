FROM golang:1.22 as builder
WORKDIR /app
COPY . /app
RUN go build -o social .
FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/social /social
EXPOSE 50053
ENTRYPOINT ["/social"]
