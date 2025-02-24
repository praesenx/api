#APP_PATH = $(shell pwd)
#DB_MIGRATIONS_VOL = "$(APP_PATH)/migrations:/migrations"
#DB_MIGRATIONS_DIR = "$(APP_PATH)/migrations"

DB_NETWORK = gocanto
APP_PATH = $(shell pwd)
DB_MIGRATIONS_PATH = migrations

fresh:
	make prune && make migrate

run:
	docker-compose up --build

migrate-up:
	docker-compose run migrate up

migrate-down:
	docker-compose run migrate down

prune:
	docker compose down --remove-orphans
	docker container prune -f
	docker image prune -f
	docker volume prune -f
	docker network prune -f

ssh:
	docker run --rm -it migrate/migrate:latest sh

migrate:
	docker run --rm \
      -v $(APP_PATH)/migrations:/migrations \
      --network gocanto \
      migrate/migrate \
      -path=/migrations \
      -database "postgres://user:password@gocanto:5432/mydatabase?sslmode=disable" \
      up