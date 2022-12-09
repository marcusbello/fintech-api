gen_doc:
	swag fmt
	swag init -g pkg/delivery/http/handler.go
start_docker:
	docker compose up -d
stop_docker:
	docker compose down
build:start_docker
	cd cmd && go build -o calculator
run:build
	cd cmd && ./calculator