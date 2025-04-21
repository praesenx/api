
env\:init:
	#echo "---> $(ROOT_ENV_FILE) <> $(ROOT_EXAMPLE_ENV_FILE) <----- | "
	rm -f $(ROOT_ENV_FILE) && cp $(ROOT_EXAMPLE_ENV_FILE) $(ROOT_ENV_FILE)

env\:check:
	@if [ ! -f .env ]; then \
	  printf "     $(RED)Error:$(NC) $(BLUE)[.env]$(NC) file not found in the current directory. Please run $(BLUE)[env:init]$(NC) first.\n"; \
	else \
	  printf "   $(BLUE)[.env]$(NC) file found.\n"; \
	fi

