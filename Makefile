.PHONY: up down build restart logs

up:
	docker-compose up 

down:
	docker-compose down

build:
	docker-compose up --build

restart:
	docker-compose down && docker-compose up --build

logs:
	docker-compose logs -f

populate_db:
	@echo "Populating database..."
	@go run scripts/populate_db/populate.go

wipe_db:
	@echo "Populating database..."
	@go run scripts/wipe_db/wipe.go