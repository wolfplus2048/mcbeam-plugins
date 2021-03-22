package opentracing

import (
	"fmt"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
	"io"
)

func New(service string, agentAddr string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		ServiceName:         service,
		Sampler:             &config.SamplerConfig{
			Type: "const",
			Param: 1,
		},
		Reporter:            &config.ReporterConfig{
			LogSpans: true,
			LocalAgentHostPort: agentAddr,
		},
	}
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("Error: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}
