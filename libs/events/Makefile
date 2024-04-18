.PHONY: docs
REPO:=github.com/KYVENetwork/cometbft/v38/libs/events

docs:
	@go get github.com/davecheney/godoc2md
	godoc2md $(REPO) > README.md

test:
	go test -v ./...
