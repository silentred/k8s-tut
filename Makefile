GITHASH := $(shell git rev-parse HEAD)

build:
	go build -o hello hello-app/main.go

docker:
	docker build -t silentred/hello-app:latest .
	docker push silentred/hello-app:latest
