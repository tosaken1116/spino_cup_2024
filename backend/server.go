package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tosaken1116/spino_cup_2024/backend/internal/app"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	ctx := context.Background()

	exporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint("otel-collector.opentelemetry.svc.cluster.local:4317"),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		fmt.Println("Failed to create exporter:", err)
		return
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			"api.spino.kurichi.dev",
			attribute.String("service.name", "test-service"),
		)),
	)
	defer func() { _ = tp.Shutdown(ctx) }()

	otel.SetTracerProvider(tp)

	tracer := otel.Tracer("test-tracer")

	// シンプルなトレースを作成
	func() {
		_, span := tracer.Start(ctx, "test-span")
		defer span.End()

		span.SetAttributes(attribute.String("key1", "value1"))
		time.Sleep(500 * time.Millisecond)
	}()

	fmt.Println("Trace sent")

	app, err := app.New()
	if err != nil {
		log.Printf("failed to create server: %v\n", err)
		return
	}
	defer func() {
		if err := app.Close(); err != nil {
			log.Printf("failed to close server: %v\n", err)
		}
	}()

	if err := app.Migrate(context.Background()); err != nil {
		log.Fatalf("failed to migrate: %v\n", err)
	}

	if err := app.Start(); err != nil {
		log.Printf("failed to start server: %v\n", err)
		return
	}
}
