VERSION := $(shell git describe --abbrev=8 --dirty --always --tags)
BINARY := escapepod
BUILDSTR := ${VERSION} ($(shell date -u +"%Y-%m-%dT%H:%M:%SZ"))
#PKGS := $(shell go list ./...)
STATIC := escapepod.toml.sample frontend-vue/dist:/

.PHONY: build
build: 
	GOFLAGS=-mod=vendor GO111MODULE=on go build -o ${BINARY} -ldflags="-w -X 'main.buildString=${BUILDSTR}'" ./cmd/escapepod

.PHONY: test
test:
	go test $(PKGS)

.PHONY: frontend-deps
frontend-deps:
	cd frontend-vue && yarn install
	
.PHONY: run-frontend
run-frontend:
	cd frontend-vue && yarn start

.PHONY: clean
clean:
	rm -f ${BINARY}

.PHONY: build-frontend
build-frontend:
	cd frontend-vue && yarn build

.PHONY: dist
dist: build
	stuffbin -a stuff -in ${BINARY} -out ${BINARY} ${STATIC}

.PHONY: run
run: build
	./${BINARY}

.PHONY: pack-releases
pack-releases:
	$(foreach var,$(RELEASE_BUILDS),stuffbin -a stuff -in ${var} -out ${var} ${STATIC} $(var);)

.PHONY: release
release:
	goreleaser release --rm-dist

.PHONY: verify
verify: install-tools
	golangci-lint run
	gosec -severity high --confidence medium -exclude G204 -quiet ./...


.PHONY: install-tools
install-tools:
	cd $(shell mktemp -d) && GO111MODULE=on go get github.com/securego/gosec/cmd/gosec@4b59c948083cd711b6a8aac8f32721b164899f57
	cd $(shell mktemp -d) && GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.17.1
