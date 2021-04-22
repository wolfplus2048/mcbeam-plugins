module github.com/wolfplus2048/mcbeam-plugins/trace/opentracing/v3

go 1.15

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
	github.com/micro/micro/v3 v3.0.4
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
)

replace github.com/micro/micro/v3 => github.com/wolfplus2048/micro/v3 v3.2.0-mcbeam
