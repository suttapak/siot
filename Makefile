build:
	docker-compose build

up:
	docker-compose up -d mosquitto db frontend

backend:
	docker compose up -d backend    

run: build up backend

start: up backend
