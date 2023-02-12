PROG = jamify-service-template
COMPOSE = docker compose

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
	go run .

.PHONY: clean
clean:
	$(COMPOSE) down
	docker system prune -f
	rm ./$(PROG)

.PHONY: stop
stop:
	$(COMPOSE) down