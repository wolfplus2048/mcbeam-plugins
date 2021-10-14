module github.com/wolfplus2048/mcbeam-plugins/ws_session/v3

go 1.14

require (
	github.com/golang/protobuf v1.4.3
	github.com/micro/micro/v3 v3.3.0
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c
	google.golang.org/protobuf v1.25.0
)

replace github.com/micro/micro/v3 v3.3.0 => github.com/wolfplus2048/micro/v3 v3.2.0-mcbeam.0.20210804071852-fbef4a5fc3eb
