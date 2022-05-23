.PHONY: all lint gen-go gen-ts

all: clean lint gen-go gen-ts

TEMPDIR:= $(shell mktemp -d)
gen-go-cmds:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
	go install github.com/twitchtv/twirp/protoc-gen-twirp@v8.1.2

gen-go: clean-go
	buf generate --template buf.gen.go.yaml protos
	cp go/* gen/go

update-go-module: gen-go
	rm -rf gen/go/go.mod gen/go/go.sum
	ln -s ../../go/go.mod gen/go
	ln -s ../../go/go.sum gen/go
	cd gen/go && go mod tidy

gen-ts-cmds:
	yarn

gen-ts: clean-ts
	buf generate --template buf.gen.ts.yaml protos

lint:
	buf lint protos

clean: clean-go clean-ts

clean-go:
	rm -rf gen/go

clean-ts:
	rm -rf gen/ts
