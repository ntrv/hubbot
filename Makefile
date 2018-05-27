BINPATH = bin
APPNAME = hubbot

.PHONY: default
default: build

.PHONY: generate
generate:
	go generate ./...

.PHONY: build-linux
build-linux: generate
	GOOS=linux GOARCH=amd64 go build -ldflags '-d -s -w' -o ${BINPATH}/${APPNAME}_linux .

.PHONY: build
build: generate
	go build -o ${BINPATH}/${APPNAME}

.PHONY: run
run:
	go run hubbot.go

.PHONY: archive
archive: build-linux
	zip -r ${BINPATH}/${APPNAME}_linux.zip ${BINPATH}/${APPNAME}_linux

.PHONY: test
test: generate
	go test -v ./...

.PHONY: integration-test
integration-test: 
	curl -XPOST -H "Content-Type: application/json" -H 'X-GitHub-Event: push' http://localhost:${PORT} -d @example/push.json

.PHONY: clean
clean:
	rm -f ${BINPATH}/*
