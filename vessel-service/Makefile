build:
	protoc -I. --go_out=. --micro_out=. proto/vessel/vessel.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t vessel-service .
run:
	docker run -p 50052:50051 \
		-e MICRO_SERVER_ADDRESS=:50051 \
		-e MICRO_REGISTRY=etcd \
		-e MICRO_REGISTRY_ADDRESS=172.17.0.2:2379 \
		vessel-service