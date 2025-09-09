FROM golang:1.22 as builder
WORKDIR /app
COPY . /app
RUN go build -o recommend .
FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/recommend /recommend
EXPOSE 50054
ENTRYPOINT ["/recommend"]
