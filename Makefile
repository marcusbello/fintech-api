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

test: docker_build_test
	docker-compose up -d
	docker-compose exec -T http go test ./...
	docker-compose down

unit_test:
	#go test `go list ./... | grep -v _test` -v
	go test ./.../http_test -v
docker_build:
	docker build . -t fintech-service:latest

docker_build_test:
	docker build . -t service_test --target=test

docker_run:
	docker run --publish 3030:8080 fintech-service