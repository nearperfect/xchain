# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: moac vnode all test clean

GOBIN = $(shell pwd)/build/bin
GO = 1.13.4

compile_chain3:
	cd internal/jsre/deps && go-bindata -nometadata -pkg deps -o bindata.go bignumber.js chain3.js && gofmt -w -s bindata.go;

#build/env.sh go run build/mci.go install ./cmd/moac
moac: #compile_chain3
	go run build/ci.go install ./cmd/moac

	@echo "Done building."
	@echo "Run \"$(GOBIN)/moac\" to launch moac."

all:
	go run build/mci.go install

android:
	go run build/mci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/moac.aar\" to use the library."

bootnode:
	go build -o ./build/bin/bootnode cmd/bootnode/main.go

ios:
	go run build/mci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Moac.framework\" to use the library."

test: all
	go run build/mci.go test

clean:
	go clean --cache
	rm -fr $(GOBIN)/moac

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/jteeuwen/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go install ./cmd/abigen
