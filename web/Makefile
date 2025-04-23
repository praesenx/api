SHELL := /bin/bash

.PHONY: fresh

fresh:
	rm -rf node_modules
	npm cache clean --force
	npm install
	npm run dev

format:
	prettier --write .
