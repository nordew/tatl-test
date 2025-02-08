.PHONY: env start

env:
	@if [ ! -f .env ]; then \
		mv .env.example .env; \
	fi

start: env
	docker-compose up