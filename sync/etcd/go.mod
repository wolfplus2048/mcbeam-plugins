module github.com/wolfplus2048/mcbeam-plugins/sync/etcd/v3

go 1.15

require (
	github.com/micro/micro/v3 v3.0.0
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	go.uber.org/zap v1.16.0 // indirect
)
replace github.com/micro/micro/v3 => github.com/wolfplus2048/micro/v3 ccce045e
