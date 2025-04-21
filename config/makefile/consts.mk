SOURCE ?= go_bindata
DATABASE ?= postgres
VERSION ?= $(shell git describe --tags 2>/dev/null | cut -c 2-)
REPO_OWNER ?= $(shell cd .. && basename "$$(pwd)")

# ------ App Configuration
ROOT_NETWORK ?= gocanto
ROOT_PATH ?= $(shell pwd)
ROOT_ENV_FILE ?= $(ROOT_PATH)/.env
ROOT_EXAMPLE_ENV_FILE? = $(ROOT_PATH)/.env.example
STORAGE_PATH ?= $(ROOT_PATH)/storage
BIN_PATH ?= $(ROOT_PATH)/bin
BIN_LOGS_PATH ?= $(ROOT_PATH)/bin/storage/logs
APP_PATH ?= $(ROOT_PATH)/

# ------ Database Configuration
# --- Docker
DB_DOCKER_SERVICE_NAME ?= postgres
DB_DOCKER_CONTAINER_NAME ?= gocanto-db
# --- Paths
DB_SEEDER_ROOT_PATH ?= $(ROOT_PATH)/database/seeder
DB_INFRA_ROOT_PATH ?= $(ROOT_PATH)/database/infra
DB_INFRA_SSL_PATH ?= $(DB_INFRA_ROOT_PATH)/ssl
DB_INFRA_DATA_PATH ?= $(DB_INFRA_ROOT_PATH)/data
# --- SSL
DB_INFRA_SERVER_CRT ?= $(DB_INFRA_SSL_PATH)/server.crt
DB_INFRA_SERVER_CSR ?= $(DB_INFRA_SSL_PATH)/server.csr
DB_INFRA_SERVER_KEY ?= $(DB_INFRA_SSL_PATH)/server.key
# --- Migrations
DB_MIGRATE_PATH ?= $(DB_INFRA_ROOT_PATH)/migrations
DB_MIGRATE_VOL_MAP ?= $(DB_MIGRATE_PATH):$(DB_MIGRATE_PATH)
