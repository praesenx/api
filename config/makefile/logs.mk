.PHONY: logs\:fresh logs\:bin\:fresh

logs\:fresh:
	find $(STORAGE_PATH)/logs -maxdepth 1 -type f -not -name ".gitkeep" -delete


