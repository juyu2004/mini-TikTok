FROM golang:1.22 as builder
WORKDIR /app
COPY . /app
RUN go build -o video .
FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/video /video
EXPOSE 50052
ENTRYPOINT ["/video"]
