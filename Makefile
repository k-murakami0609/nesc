dev:
	air

test:
	go test ./...

lint:
	go lint ./...

push-image:
	sh docker.sh