.PHONY: all
all: test

.PHONY: init
init:
	tfenv install
	tfenv use

.PHONY: test
test: init
	go mod download
	go test -v -timeout 60m

.PHONY: clean
clean:
	rm -rf ../../examples/complete/*.tfstate*
