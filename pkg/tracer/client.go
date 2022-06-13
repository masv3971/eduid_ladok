package tracer

import (
	"context"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"os"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

// Trace type
type Trace struct {
	logger *logger.Logger
	cfg    *model.Cfg
	*tracesdk.TracerProvider
	fileHandler *os.File
}

// New returns an OpenTelemetry Trace struct
func New(cfg *model.Cfg, logger *logger.Logger) (*Trace, error) {
	trace := &Trace{
		logger: logger,
		cfg:    cfg,
	}

	var (
		exp tracesdk.SpanExporter
		err error
	)

	switch cfg.Tracing.Kind {
	case "file":
		exp, err = trace.newFileExporter(cfg.Tracing.Endpoint)
		if err != nil {
			return nil, err
		}
	case "jaeger":
		exp, err = trace.newJaegerExporter(cfg.Tracing.Endpoint)
		if err != nil {
			return nil, err
		}
	}

	res, err := trace.newResource()
	if err != nil {
		return nil, err
	}

	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(res),
	)

	trace.TracerProvider = tp

	return trace, nil
}

func (t *Trace) newFileExporter(path string) (tracesdk.SpanExporter, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	t.fileHandler = f

	exp, err := stdouttrace.New(
		stdouttrace.WithWriter(t.fileHandler),
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithoutTimestamps(),
	)

	return exp, nil
}

func (t *Trace) newJaegerExporter(url string) (tracesdk.SpanExporter, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(
		jaeger.WithEndpoint(url),
	))
	if err != nil {
		return nil, err
	}
	return exp, nil
}

func (t *Trace) newResource() (*resource.Resource, error) {
	env := "development"
	if t.cfg.Production {
		env = "production"
	}

	r, err := resource.Merge(
		resource.Default(),
		resource.NewSchemaless(
			semconv.ServiceNameKey.String("eduid_ladok"),
			semconv.TelemetrySDKLanguageGo,
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", env),
		),
	)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Close close trace
func (t *Trace) Close(ctx context.Context) {
	t.logger.Info("Quit")
	if err := t.Shutdown(ctx); err != nil {
		panic(err)
	}

	if t.fileHandler != nil {
		t.fileHandler.Close()
	}
}
