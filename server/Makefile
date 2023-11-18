build:
	go build -o bin/vsoutlook-server
	ls -lh bin/vsoutlook-server

schema:
	dropdb vsoutlook_development
	createdb vsoutlook_development
	go run main.go db migrate

migrate:
	go run main.go db migrate

watch:
	air

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/linux-vsoutlook-server -ldflags '-s -w' \
		-gcflags=-trimpath=/Users/cui \
    -asmflags=-trimpath=/Users/cui
	upx bin/linux-vsoutlook-server

run:
	go run main.go

temp:
	go run main.go temp

deploy-prod:
	sh scripts/deploy-prod.sh
