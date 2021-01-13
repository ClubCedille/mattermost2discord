.PHONY: test mock api

APP_NAME=mm2disc
CGO_ENABLED=0
GOOS=linux

build:
	go build -a -installsuffix cgo -o ${APP_NAME} .

run: build
	./${APP_NAME}

docker:
	docker-compose up --build

test:
	go test ./... -v --bench --benchmem -coverprofile=coverage.out
	go tool cover -html=coverage.out