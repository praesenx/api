.PHONY: fresh audit watch format

format:
	gofmt -w -s .

fresh:
	rm -rf $(DB_INFRA_DATA_PATH) && \
	docker compose down --remove-orphans && \
	docker container prune -f && \
	docker image prune -f && \
	docker volume prune -f && \
	docker network prune -f && \
	docker ps

audit:
	$(call external_deps,'.')
	$(call external_deps,'./bin/...')
	$(call external_deps,'./app/...')
	$(call external_deps,'./database/...')
	$(call external_deps,'./docs/...')

watch:
	# --- Works with (air).
	# https://github.com/air-verse/air
	cd $(APP_PATH) && air


define external_deps
	@echo '-- $(1)';  go list -f '{{join .Deps "\n"}}' $(1) | grep -v github.com/$(REPO_OWNER)/blog | xargs go list -f '{{if not .Standard}}{{.ImportPath}}{{end}}'
endef
