.PHONY: test mock api

APP_NAME=mm2disc
CGO_ENABLED=0
GOARCH=amd64

build:
	GOOS=linux go build -a -installsuffix cgo -o ${APP_NAME} .

build-windows:
	go build -a -installsuffix cgo -o ${APP_NAME} .

run-windows: build-windows
	./${APP_NAME}

run: build
	./${APP_NAME}

docker:
	docker-compose up --build

test:
	go test ./... -v --bench --benchmem -coverprofile=coverage.out
	go tool cover -html=coverage.out