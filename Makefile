up:
	docker compose up -d

test:
	go test ./...

stop:
	docker stop $(docker ps -aq)