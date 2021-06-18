module github.com/wolfplus2048/mcbeam-plugins/store/redis/v3

go 1.13

require (
	github.com/go-redis/redis/v8 v8.8.2 // indirect
	github.com/micro/micro/v3 v3.2.0
)

replace github.com/micro/micro/v3 => github.com/wolfplus2048/micro/v3 v3.2.0-mcbeam.0.20210421085145-e980dbeea9d6
