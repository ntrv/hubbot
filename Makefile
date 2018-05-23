BINPATH = bin
APPNAME = hubbot

.PHONY: default
default: build

.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags '-d -s -w' -o ${BINPATH}/${APPNAME}_linux .

.PHONY: build
build:
	go build -o ${BINPATH}/${APPNAME}

.PHONY: run
run:
	go run hubbot.go

.PHONY: archive
archive: build-linux
	zip -r ${BINPATH}/${APPNAME}_linux.zip ${BINPATH}/${APPNAME}_linux

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	rm -f ${BINPATH}/*
