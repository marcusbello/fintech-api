# Set the version number
VERSION = 0.0.1

print-version:
	@echo "Version: $(VERSION)"

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
	docker-compose up -d fintech-test
	docker-compose exec -T fintech-test go test ./...
	docker-compose down

unit_test:
	#go test `go list ./... | grep -v _test` -v
	go test ./...

docker_build:
	docker build -t fintech-app:$(VERSION) .

docker_build_test:
	docker build --no-cache -t fintech-app-test:test --target=test .

docker_run:
	docker run --publish 3030:3030 fintech-app:$latest

docker_tag:
	docker tag fintech-app:$(VERSION) marcusbello/fintech-app:latest
	docker tag fintech-app:$(VERSION) marcusbello/fintech-app:$(VERSION)

docker_push:docker_build docker_tag
	docker push -a marcusbello/fintech-app