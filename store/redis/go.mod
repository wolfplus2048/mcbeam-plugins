module github.com/wolfplus2048/mcbeam-plugins/store/redis/v3

go 1.13

require (
	github.com/asim/go-micro/v3 v3.0.0-20210120135431-d94936f6c97c
	github.com/go-redis/redis/v7 v7.4.0
	github.com/micro/micro/v3 v3.0.4
)

replace github.com/micro/micro/v3 => github.com/wolfplus2048/micro/v3 v3.2.0-mcbeam
