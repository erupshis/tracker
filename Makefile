
.PHONY: build-container docker_up docker_down
build-container:
	docker build -t tracker -f build/Dockerfile .

docker_down:
	docker compose -f ./build/docker-compose.yaml stop && docker compose -f ./build/docker-compose.yaml down

docker_up:
	docker compose -f ./build/docker-compose.yaml up

