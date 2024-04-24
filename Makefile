PARAMS := $(wordlist 2,100,${MAKECMDGOALS})

.PHONY: start-postgres
start-postgres:
	docker-compose up -d migrate ${PARAMS}

.PHONY: start-minio
start-minio:
	docker-compose up -d minio ${PARAMS}

.PHONY: start-db
start-db:
	make start-postgres && make start-minio

.PHONY: drop
drop:
	docker-compose rm --stop --force

.PHONY: start-app
start-app:
	docker-compose up -d gamelib ${PARAMS}

.PHONY: runserver
runserver:
	go run ./cmd/server/main.go -env="dev"

.PHONY: build
build:
	go build ./cmd/server/main.go