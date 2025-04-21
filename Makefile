SHELL := /bin/bash # <--- Add this line at the top

RAND    := $(shell echo $$RANDOM)
GREEN   := \033[0;32m
BLUE    := \033[0;34m
RED     := \033[0;31m
YELLOW  := \033[1;33m
NC      := \033[0m

.PHONY: fresh audit watch format

.PHONY: env\:init env\:check

.PHONY: build\:app build\:app\:linux build\:release build\:run build\:fresh

.PHONY: db\:local db\:up db\:ping db\:bash db\:fresh db\:logs
.PHONY: db\:delete db\:secure db\:secure\:show db\:chmod db\:seed
.PHONY: db\:migrate db\:withdraw db\:migrate\:create db\:migrate\:force

.PHONY: logs\:fresh logs\:bin\:fresh

# -------------------------------------------------------------------------------------------------------------------- #
# ------------------------------------------------ App Configuration ------------------------------------------------- #
# -------------------------------------------------------------------------------------------------------------------- #
SOURCE ?= go_bindata
DATABASE ?= postgres
VERSION ?= $(shell git describe --tags 2>/dev/null | cut -c 2-)
REPO_OWNER ?= $(shell cd .. && basename "$$(pwd)")
ROOT_NETWORK ?= gocanto
ROOT_PATH ?= $(shell pwd)
ROOT_ENV_FILE ?= $(ROOT_PATH)/.env
ROOT_EXAMPLE_ENV_FILE := $(ROOT_PATH)/.env.example
STORAGE_PATH ?= $(ROOT_PATH)/storage
BIN_PATH ?= $(ROOT_PATH)/bin
BIN_LOGS_PATH ?= $(ROOT_PATH)/bin/storage/logs
APP_PATH ?= $(ROOT_PATH)/
# -------------------------------------------------------------------------------------------------------------------- #
# -------------------------------------------------------------------------------------------------------------------- #

include ./config/makefile/env.mk
include ./config/makefile/app.mk
include ./config/makefile/db.mk
include ./config/makefile/build.mk
include ./config/makefile/logs.mk
