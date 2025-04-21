logs\:fresh:
	find $(STORAGE_PATH)/logs -maxdepth 1 -type f -not -name ".gitkeep" -delete

logs\:bin\:fresh:
	@rm -rf "$(BIN_LOGS_PATH)"
	@mkdir -m 777 $(BIN_LOGS_PATH)
	@touch $(BIN_LOGS_PATH)/.gitkeep
