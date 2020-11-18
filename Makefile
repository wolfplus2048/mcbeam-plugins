.PHONY: build clean  test proto
test: vet
	go test -v ./...

clean:
	rm -rf ./micro

proto:
	protoc --proto_path=. --micro_out=. --go_out=. session/proto/gate.proto

