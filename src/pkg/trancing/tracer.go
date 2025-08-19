package tracing

import (
        "context"
        "mini-douyin/src/config"

        "go.opentelemetry.io/otel"
        "go.opentelemetry.io/otel/exporters/jaeger"
        "go.opentelemetry.io/otel/sdk/resource"
        tracesdk "go.opentelemetry.io/otel/sdk/trace"
        semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

// InitTracer 初始化Jaeger链路追踪
func InitTracer(serviceName string) (*tracesdk.TracerProvider, error) {
        // 创建Jaeger exporter（基于HTTP协议）
        exp, err := jaeger.New(
                jaeger.WithCollectorEndpoint(
                        jaeger.WithEndpoint(config.JaegerEndpoint), // 从配置读取Jaeger地址
                ),
        )
        if err != nil {
                return nil, err
        }

        // 创建TracerProvider（配置采样率为100%，开发环境用；生产环境可改为0.5）
        tp := tracesdk.NewTracerProvider(
                tracesdk.WithBatcher(exp), // 批量导出追踪数据
                tracesdk.WithResource(resource.NewWithAttributes(
                        semconv.SchemaURL,
                        semconv.ServiceName(serviceName), // 服务名称（用于Jaeger识别）
                )),
                tracesdk.WithSampler(tracesdk.AlwaysSample()), // 100%采样
        )

        // 设置全局Tracer（后续代码可通过otel.Tracer()获取）
        otel.SetTracerProvider(tp)
        return tp, nil
}

// GetTracer 获取Tracer实例（业务代码中用）
func GetTracer(serviceName string) tracesdk.Tracer {
        return otel.Tracer(serviceName)
}