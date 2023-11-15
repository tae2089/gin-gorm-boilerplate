
IMG_NAME = tae2089/gin-bolier
IMG_TAG = latest

image-build:
	docker build -t ${IMG_NAME}:${IMG_TAG} -f build/ci/Dockerfile . --build-arg GOPATH=${GOPATH}

docker-compose-run:
	docker compose -f build/cd/docker-compose.yml --env-file .env up -d

docker-compose-restart:
	docker compose -f build/cd/docker-compose.yml --env-file .env restart

docker-compose-down:
	docker compose -f build/cd/docker-compose.yml --env-file .env down

postgresql-run:
	docker compose -f build/cd/postgres-compose.yml --env-file .env up -d

postgresql-down:
	docker compose -f build/cd/postgres-compose.yml --env-file .env down