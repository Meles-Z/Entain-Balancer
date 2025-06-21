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
