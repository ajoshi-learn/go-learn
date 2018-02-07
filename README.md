# Golang cheat sheet

## Installation

### System variables

`GOROOT` - go binary path

`GOPATH` - workspace path

`go env GOROOT` - check current `GOROOT`
`go env GOPATH` - check current `GOPATH`


```
export GOPATH=$HOME/workspace
export PATH=$PATH:/usr/local/go/bin
```

### Workspace structure

```
bin/
    hello                 # application binary
pkg/
    linux_amd64/          #
        github.com/user/
            stringutil.a  # object binary file
src/
    github.com/user/
        hello/
            hello.go      # app source code
        stringutil/
            reverse.go    # lib source code
```

## Go commands

```
build       compile packages and dependencies
clean       remove object files
env         print Go environment information
fix         run go tool fix on packages
fmt         run gofmt on package sources
get         download and install packages and dependencies
install     compile and install packages and dependencies
run         compile and run Go program
test        test packages
tool        run specified go tool
version     print Go version
```