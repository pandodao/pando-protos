# pando-protos

Protobuf definitions of Pando services.

## Introduction

We use this repository to save protobuf files of pando services, [buf](https://github.com/bufbuild/buf) is the very convenient tool to help us manage protobuf dependencies. Combined with GitHub Actions, when you have a new commit, it will generate and submit the Golang/TypeScript files for you, this allows you to focus on business logic rather than enviroment settings.

| Branch         | Generated Go Branch    | Generated TS Branch    |
| -------------- | ---------------------- | ---------------------- |
| main           | gen-go                 | gen-ts                 |
| ${BRANCH_NAME} | ${BRANCH_NAME}\_gen-go | ${BRANCH_NAME}\_gen-ts |

Finally, When a PR is closed, the generated branches will be cleaned up automatically, for the user, this will also force you to update the dependency.

## Use in project

Let's create a Go project as an example.

```
mkdir $TMPDIR/test && cd $TMPDIR/test
git init
go mod init test
```

Go supports [Private Modules](https://go.dev/ref/mod#private-modules), but it requires you to do extra configuration. Another choice is [git submodule](https://git-scm.com/book/en/v2/Git-Tools-Submodules), it's easy to use.

```
git submodule add -b gen-go git@github.com:fox-one/pando-protos.git
```

Add a `main.go` file.

```
package main

import (
        "fmt"

        "github.com/fox-one/pando-protos/pando/v1"
)

func main() {
        fmt.Printf("%+v\n", pando.Asset{})
}
```

Then when you execute `go mod tidy`, you will see the error.

```
go: finding module for package github.com/fox-one/pando-protos/pando/v1
test imports
        github.com/fox-one/pando-protos/pando/v1: cannot find module providing package github.com/fox-one/pando-protos/pando/v1: module github.com/fox-one/pando-protos/pando: git ls-remote -q origin in $HOME/go/pkg/mod/cache/vcs/d8c94b5ac6e1e6f73463b11d144b9ef66de4df3a3afc4805c5dbaaccf3d97e38: exit status 128:
        fatal: could not read Username for 'https://github.com': terminal prompts disabled
Confirm the import path was entered correctly.
If this is a private repository, see https://golang.org/doc/faq#git_https for additional information.
```

Update `go.mod`, add replace statement, and then it will work.

```
$ go mod edit -replace github.com/fox-one/pando-protos=./pando-protos
$ go mod tidy
$ go run main.go
{state:{NoUnkeyedLiterals:{} DoNotCompare:[] DoNotCopy:[] atomicMessageInfo:<nil>} sizeCache:0 unknownFields:[] AssetId: ChainId: Symbol: Name: Logo: Price:}
```

## Local generated

If you dislike that GitHub action runs too slowly, or the way disrupts your development, you can make protos locally.

### Preinstall

Usually, those tools are backwards compatibility, the installed version should be larger than version in the table to avoid problems. And if you encounter unexpected differences in the new version, roll back the version.

| Tool                                                | Version |
| --------------------------------------------------- | ------- |
| [protoc](https://grpc.io/docs/protoc-installation/) | 3.9.1   |
| [buf](https://docs.buf.build/installation)          | 1.4.0   |

### Make

For Backend(Go) Developer, require go installed.

```
# once
make gen-go-cmds

make lint && make gen-go
```

For Frontend(TS) Developer, require node installed.

```
# once
make gen-ts-cmds

make lint && make gen-ts
```

Both.

```
# once
make gen-go-cmds && make gen-ts-cmds

make
```

### Use your local protos in project

For Go, you can specify the path in `go.mod`.

```
go mod edit -replace github.com/fox-one/pando-protos=$YOUR_PROTOS_PATH/gen/go
```

For TS, you can create a soft link point to your local protos.

```
ln -s $YOUR_PROTOS_PATH/gen/ts pando-protos
```
