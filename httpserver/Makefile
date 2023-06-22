export tag=v1.0
root:
	export ROOT=github.com/yuewibing/golang

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t yuweibing/httpserver:${tag} .

push: release
	echo "pushing cncamp/httpserver"
	docker push yuweibing/httpserver:v1.0
