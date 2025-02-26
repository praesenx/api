DB_NETWORK = gocanto
APP_PATH = $(shell pwd)
DB_MIGRATIONS_PATH = migrations

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



prune:
	docker compose down --remove-orphans
	docker container prune -f
	docker image prune -f
	docker volume prune -f
	docker network prune -f

