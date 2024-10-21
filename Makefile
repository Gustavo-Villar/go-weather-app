# Makefile for running Docker Compose

.PHONY: up up-detached down restart restart-detached

up:
	docker-compose up

up-detached:
	docker-compose up -d

down:
	docker-compose down

restart:
	docker-compose down
	docker-compose up

restart-detached:
	docker-compose down
	docker-compose up -d
