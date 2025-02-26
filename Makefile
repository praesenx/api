DB_NETWORK = gocanto
APP_PATH = $(shell pwd)
DB_MIGRATIONS_PATH = migrations

DB_SSL_PATH = $(shell pwd)/database/ssl

env\:generate:
	cp $(shell pwd)/.env.example $(shell pwd)/.env

db\:up:
	docker-compose up postgres -d

db\:ping:
	docker port gocanto-db

db\:bash:
	docker exec -it gocanto-db bash

db\:fresh:
	docker-compose down postgres && \
	make db:prune && \
	rm -rf $(shell pwd)/database/data && \
	make db:up

db\:prune:
	docker container prune -f
	docker image prune -f
	docker volume prune -f
	docker network prune -f

db\:logs:
	docker logs -f gocanto-db

db\:secure:
	rm -rf $(DB_SSL_PATH)/*.* && \
	make prune && \
	openssl genpkey -algorithm RSA -out $(DB_SSL_PATH)/server.key && \
    openssl req -new -key $(DB_SSL_PATH)/server.key -out $(DB_SSL_PATH)/server.csr && \
    openssl x509 -req -days 365 -in $(DB_SSL_PATH)/server.csr -signkey $(DB_SSL_PATH)/server.key -out $(DB_SSL_PATH)/server.crt && \
    chmod 600 $(DB_SSL_PATH)/server.key && \
    chmod 600 $(DB_SSL_PATH)/server.crt

prune:
	rm -rf $(shell pwd)/database/data && \
	docker compose down --remove-orphans && \
	docker container prune -f && \
	docker image prune -f && \
	docker volume prune -f && \
	docker network prune -f

