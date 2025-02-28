# --- Main Configuration
ROOT_NETWORK=gocanto
ROOT_PATH=$(shell pwd)
ROOT_ENV_FILE="$(ROOT_PATH)/.env"
ROOT_EXAMPLE_ENV_FILE="$(ROOT_PATH)/.env.example"

# --- Database Configuration
DB_DOCKER_SERVICE_NAME="postgres"
DB_DOCKER_CONTAINER_NAME="gocanto-db"
DB_ROOT_PATH="$(ROOT_PATH)/database"
DB_SSL_PATH="$(DB_ROOT_PATH)/ssl"
DB_DATA_PATH="$(DB_ROOT_PATH)/data"
DB_SERVER_CRT="$(DB_SSL_PATH)/server.crt"
DB_SERVER_CSR="$(DB_SSL_PATH)/server.csr"
DB_SERVER_KEY="$(DB_SSL_PATH)/server.key"

env\:new:
	rm -f $(ROOT_ENV_FILE)
	cp $(ROOT_EXAMPLE_ENV_FILE) $(ROOT_ENV_FILE)

db\:up:
	docker-compose up $(DB_DOCKER_SERVICE_NAME) -d

db\:ping:
	docker port $(DB_DOCKER_CONTAINER_NAME)

db\:bash:
	docker exec -it $(DB_DOCKER_CONTAINER_NAME) bash

db\:fresh:
	docker-compose down $(DB_DOCKER_SERVICE_NAME) && \
	make db:prune && \
	rm -rf $(DB_DATA_PATH) && \
	make db:up

db\:prune:
	docker container prune -f
	docker image prune -f
	docker volume prune -f
	docker network prune -f

db\:logs:
	docker logs -f $(DB_DOCKER_CONTAINER_NAME)

db\:secure:
	make prune && \
	rm -rf $(DB_SERVER_CRT) && rm -rf $(DB_SERVER_CSR) && rm -rf $(DB_SERVER_KEY) && make prune && \
	openssl genpkey -algorithm RSA -out $(DB_SSL_PATH)/server.key && \
    openssl req -new -key $(DB_SERVER_KEY) -out $(DB_SERVER_CSR) && \
    openssl x509 -req -days 365 -in $(DB_SERVER_CSR) -signkey $(DB_SERVER_KEY) -out $(DB_SERVER_CRT) && \
    chmod 600 $(DB_SERVER_KEY) && chmod 600 $(DB_SERVER_CRT)

prune:
	rm -rf $(DB_DATA_PATH) && \
	docker compose down --remove-orphans && \
	docker container prune -f && \
	docker image prune -f && \
	docker volume prune -f && \
	docker network prune -f
