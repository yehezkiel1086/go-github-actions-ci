composeup:
	docker compose up -d

composedown:
	docker compose down

run:
	go run main.go

.PHONY: composeup composedown run
