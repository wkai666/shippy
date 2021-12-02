build:
	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-cli .
run:
	docker run -e MICRO_REGISTRY=etcd -e MICRO_REGISTRY_ADDRESS=172.17.0.1:2379 consignment-cli
