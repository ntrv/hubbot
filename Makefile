BINPATH = bin
APPNAME = hubbot

.PHONY: default
default: archive

.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -ldflags '-d -s -w' -o ${BINPATH}/${APPNAME} .

.PHONY: archive
archive: build
	zip -r ${BINPATH}/${APPNAME}.zip ${BINPATH}/${APPNAME}

.PHONY: clean
clean:
	rm -f ${BINPATH}/*
