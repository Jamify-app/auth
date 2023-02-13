PROG = auth
COMPOSE = docker compose
COMPOSE_PLATFORM = docker-compose-platform.yaml
COMPOSE_SERVICE = docker-compose.yaml
PORT = 8121
MONGODB_USERNAME = root
MONGODB_PASSWORD = password
MONGODB_PORT = 27017
ENV = PORT=$(PORT) \
	  MONGODB_USERNAME=$(MONGODB_USERNAME) \
	  MONGODB_PASSWORD=$(MONGODB_PASSWORD) \
	  MONGODB_PORT=$(MONGODB_PORT) \

.PHONY: init
init:
	go mod tidy

.PHONY: build
build:
	$(COMPOSE) build
	go build -o ./$(PROG)

.PHONY: start
start:
	$(COMPOSE) -f $(COMPOSE_SERVICE) up --build -d
	./$(PROG)

.PHONY: platform
platform:
	$(COMPOSE) -f $(COMPOSE_PLATFORM) up --build -d

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

.PHONY: test
test:
	$(ENV) go test ./...