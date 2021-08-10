module github.com/wolfplus2048/mcbeam-plugins/store/minio/v3

go 1.13

require (
	github.com/micro/micro/v3 v3.3.0
	github.com/minio/minio-go/v7 v7.0.12
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
)

replace github.com/micro/micro/v3 => github.com/wolfplus2048/micro/v3 v3.2.0-mcbeam.0.20210804071852-fbef4a5fc3eb
