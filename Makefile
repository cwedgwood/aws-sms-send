# by default we build this and only this
proj := $(shell basename $(shell pwd))

default: $(proj) # runtest

all: $(proj)

runtest: $(proj)
	time ./$(proj) -h

fmt:
	gofmt -s=true -w *.go

clean:
	rm -f *~ */*~ .*~ $(proj)
	go clean

$(proj): Makefile *.go
	go vet
	go build
	strip $(proj)

install: $(proj)
	install --compare $(proj) /usr/local/bin/

.PHONY: all runtest fmt clean install
