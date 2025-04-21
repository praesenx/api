build\:fresh:
	make build:app && make build:run

build\:app:
	cp $(ROOT_PATH)/.env.production $(BIN_PATH)/.env && \
	make logs:bin:fresh && \
	rm -f $(ROOT_PATH)/bin/app && \
	CGO_ENABLED=0 go build -a -ldflags='-X main.Version=$(VERSION)' -o "$(ROOT_PATH)/bin/app" -tags '$(DATABASE) $(SOURCE)' $(APP_PATH)

build\:app\:linux:
	make logs:bin:fresh && \
	cd $(APP_PATH) && \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o "$(ROOT_PATH)/bin/app_linux" -ldflags='-X main.Version=$(VERSION) -extldflags "-static"' -tags '$(DATABASE) $(SOURCE)' $(APP_PATH)

build\:run:
	cd $(BIN_PATH) && ./app

build\:release:
	git tag v$(V)
	@read -p "Press enter to confirm and push to origin ..." && git push origin v$(V)
