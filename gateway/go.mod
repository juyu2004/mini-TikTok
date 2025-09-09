module github.com/example/mini-tiktok/gateway

go 1.22

require (
	github.com/gin-gonic/gin v1.10.0
	github.com/golang-jwt/jwt/v5 v5.2.0
	google.golang.org/grpc v1.64.0
	github.com/hashicorp/consul/api v1.26.1
	go.opentelemetry.io/otel v1.27.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.27.0
	go.opentelemetry.io/otel/sdk v1.27.0
)

replace github.com/example/mini-tiktok/proto => ../
