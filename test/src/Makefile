.PHONY: all
all: test

.PHONY: init-tfenv
init-tfenv:
	tfenv install
	tfenv use

.PHONY: init-tgenv
init-tgenv:
	tgenv install latest:0
	tgenv use latest:0

.PHONY: init
init: init-tfenv init-tgenv

.PHONY: test
test: init
	go mod download
	go test -v -timeout 60m

.PHONY: clean
clean:
	rm -rf ../../examples/complete/*.tfstate*
