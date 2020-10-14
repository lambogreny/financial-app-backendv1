VERSION=$(shell git rev-parse --short HEAD)


run:
	docker-compose build  --build-arg APP_VERSION=$(VERSION)
	docker-compose up server 

build-dev:
	docker build --build-arg APP_VERSION=$(VERSION) .
