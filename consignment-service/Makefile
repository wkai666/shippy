build:
	protoc -I. --go_out=. --micro_out=. proto/consignment/consignment.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service .
run:
	docker run -p 50051:50051 \
        -e MICRO_SERVER_ADDRESS=:50051 \
        -e MICRO_REGISTRY=etcd \
        -e MICRO_REGISTRY_ADDRESS=172.17.0.1:2379 \
        consignment-service