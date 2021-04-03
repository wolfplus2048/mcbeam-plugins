module github.com/wolfplus2048/mcbeam-plugins/sync/etcd/v3

go 1.15

require (
	github.com/micro/micro/v3 v3.0.4
	go.etcd.io/etcd v0.5.0-alpha.5.0.20200425165423-262c93980547
	go.uber.org/zap v1.16.0
)

replace github.com/micro/micro/v3 => github.com/wolfplus2048/micro/v3 v3.2.0-mcbeam.0.20210402072107-ccce045ec5fe
