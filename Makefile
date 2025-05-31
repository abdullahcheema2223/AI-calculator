# Makefile for Intelligent Calculator

.PHONY: all backend frontend start-backend start-frontend install-backend install-frontend

all: install-backend install-frontend

backend:
	cd backend && go build -o calculator

frontend:
	cd frontend && npm run build

start-backend:
	cd backend && go run main.go

start-frontend:
	cd frontend && npm start

install-backend:
	cd backend && go mod tidy

install-frontend:
	cd frontend && npm install 