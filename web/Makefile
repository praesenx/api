SHELL := /bin/bash

.PHONY: format env\:fresh lint\:check lint\:fix

format:
	npx prettier --write .

env\:fresh:
	rm -rf node_modules
	npm cache clean --force
	npm install

lint\:check:
	npx eslint . --ext .js,.jsx,.cjs,.mjs,.vue

lint\:fix:
	npx eslint . --ext .js,.jsx,.cjs,.mjs,.vue --fix
