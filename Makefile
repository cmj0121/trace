SRC := $(wildcard *.go)

.PHONY: all clean test run build upgrade help

all: build 		# default action
	@pre-commit install --install-hooks >/dev/null
	@git config commit.template .git-commit-template

clean:			# clean-up environment

build:			# build the binary/library
	gofmt -w -s $(SRC)
	go test -bench=. -cover -failfast -timeout 2s ./...

upgrade:		# upgrade all the necessary packages
	pre-commit autoupdate

help:			# show this message
	@printf "Usage: make [OPTION]\n"
	@printf "\n"
	@perl -nle 'print $$& if m{^[\w-]+:.*?#.*$$}' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?#"} {printf "    %-18s %s\n", $$1, $$2}'
