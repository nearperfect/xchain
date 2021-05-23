.PHONY: xchain test clean

GOBIN = $(shell pwd)/build/bin
GO = 1.13.4

xchain:
	go run build/ci.go install ./cmd/xchain

	@echo "Done building."
	@echo "Run \"$(GOBIN)/xchain\" to launch xchain."

abigen:
	@echo "Done building."
	@echo "Run \"$(GOBIN)/xchain\" to launch xchain."
	go run build/ci.go install ./cmd/abigen


bootnode:
	go build -o ./build/bin/bootnode cmd/bootnode/main.go


test: all
	go run build/mci.go test

clean:
	go clean --cache
	rm -fr $(GOBIN)/xchain
