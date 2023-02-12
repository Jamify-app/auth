PROG = auth
COMPOSE = docker compose
PORT = 8121
MONGODB_USERNAME = root
MONGODB_PASSWORD = passoword
MONGODB_PORT = 27017
ENV = PORT=$(PORT) \
	  MONGODB_USERNAME=$(MONGODB_USERNAME) \
	  MONGODB_PASSWORD=$(MONGODB_PASSWORD) \
	  MONGODB_PORT=$(MONGODB_PORT)

.PHONY: init
init:
	go mod tidy

.PHONY: build
build:
	$(COMPOSE) build
	go build -o ./$(PROG)

.PHONY: start
start:
	$(COMPOSE) up -d
	./$(PROG)

.PHONY: run
run:
	$(ENV) go run .

.PHONY: clean
clean:
	$(COMPOSE) down
	docker system prune -f
	rm ./$(PROG)

.PHONY: stop
stop:
	$(COMPOSE) down