#
# Makefile to compile opmlsort for Mac OS X, Linux, Windows 7
# as well as Raspberry Pi Zero, 1, 2, and 3.
#

PROJECT = opml

MOTTO = "An OPML parser package plus opml cat and sort utilities"

VERSION = $(shell grep -m 1 'Version = `' opml.go | cut -d\` -f 2)

BRANCH = $(shell git branch | grep '* ' | cut -d\  -f 2)

PKGASSETS = $(shell which pkgassets)

OS = $(shell uname)

EXT = 
ifeq ($(OS), Windows)
	EXT = .exe
endif


CLI_NAMES = opmlsort opmlcat opml2json

build: $(CLI_NAMES)

opmlsort: bin/opmlsort$(EXT)

opmlcat: bin/opmlcat$(EXT)

opml2json: bin/opml2json$(EXT)

bin/opmlsort$(EXT): opml.go cmd/opmlsort/opmlsort.go
	env go build -o bin/opmlsort$(EXT) cmd/opmlsort/opmlsort.go


bin/opmlcat$(EXT): opml.go cmd/opmlcat/opmlcat.go
	env go build -o bin/opmlcat$(EXT) cmd/opmlcat/opmlcat.go

bin/opml2json$(EXT): opml.go cmd/opml2json/opml2json.go
	env go build -o bin/opml2json$(EXT) cmd/opml2json/opml2json.go

test:
	go test

man: build
	mkdir -p man/man1
	bin/opmlsort -generate-manpage | nroff -Tutf8 -man > man/man1/opmlsort.1
	bin/opmlcat -generate-manpage | nroff -Tutf8 -man > man/man1/opmlcat.1
	bin/opml2json -generate-manpage | nroff -Tutf8 -man > man/man1/opml2json.1


install:
	env CGO_ENABLED=0 GOBIN=$(HOME)/bin go install cmd/opmlsort/opmlsort.go
	env CGO_ENABLED=0 GOBIN=$(HOME)/bin go install cmd/opmlcat/opmlcat.go
	env CGO_ENABLED=0 GOBIN=$(HOME)/bin go install cmd/opml2json/opml2json.go

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -d man ]; then rm -fR man; fi

website:
	bash mk-website.bash $(PROJECT) $(MOTTO) $(VERSION)

dist/linux-amd64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/opmlsort cmd/opmlsort/opmlsort.go
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/opmlcat cmd/opmlcat/opmlcat.go
	env GOOS=linux GOARCH=amd64 go build -o dist/bin/opml2json cmd/opml2json/opml2json.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTALL.md docs/* bin/*
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/opmlsort.exe cmd/opmlsort/opmlsort.go
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/opmlcat.exe cmd/opmlcat/opmlcat.go
	env GOOS=windows GOARCH=amd64 go build -o dist/bin/opml2json.exe cmd/opml2json/opml2json.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTALL.md docs/* bin/*
	rm -fR dist/bin

dist/macosx-amd64:
	mkdir -p dist/bin
	env GOOS=darwin GOARCH=amd64 go build -o dist/bin/opmlsort cmd/opmlsort/opmlsort.go
	env GOOS=darwin GOARCH=amd64 go build -o dist/bin/opmlcat cmd/opmlcat/opmlcat.go
	env GOOS=darwin GOARCH=amd64 go build -o dist/bin/opml2json cmd/opml2json/opml2json.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macosx-amd64.zip README.md LICENSE INSTALL.md docs/* bin/*
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/opmlsort cmd/opmlsort/opmlsort.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/opmlcat cmd/opmlcat/opmlcat.go
	env GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/opml2json cmd/opml2json/opml2json.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTALL.md docs/* bin/*
	rm -fR dist/bin

dist/raspbian-arm6:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm GOARM=6 go build -o dist/bin/opmlsort cmd/opmlsort/opmlsort.go
	env GOOS=linux GOARCH=arm GOARM=6 go build -o dist/bin/opmlcat cmd/opmlcat/opmlcat.go
	env GOOS=linux GOARCH=arm GOARM=6 go build -o dist/bin/opml2json cmd/opml2json/opml2json.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm6.zip README.md LICENSE INSTALL.md docs/* bin/*
	rm -fR dist/bin

dist/linux-arm64:
	mkdir -p dist/bin
	env GOOS=linux GOARCH=arm64 GOARM=6 go build -o dist/bin/opmlsort cmd/opmlsort/opmlsort.go
	env GOOS=linux GOARCH=arm64 GOARM=6 go build -o dist/bin/opmlcat cmd/opmlcat/opmlcat.go
	env GOOS=linux GOARCH=arm64 GOARM=6 go build -o dist/bin/opml2json cmd/opml2json/opml2json.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-arm64.zip README.md LICENSE INSTSALL.md docs/* bin/*
	rm -fR dist/bin

generate_usage_pages: opmlsort opmlcat opml2json
	bash gen-usage-pages.bash

distribute_docs:
	mkdir -p dist/docs
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	bash gen-usage-pages.bash
	cp -v docs/opmlsort.md dist/docs/
	cp -v docs/opmlcat.md dist/docs/
	cp -v docs/opml2json.md dist/docs/

release: generate_usage_pages distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7 dist/raspbian-arm6 dist/linux-arm64

publish:
	bash mk-website.bash
	bash publish.bash

