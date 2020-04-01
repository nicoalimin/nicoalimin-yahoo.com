ifneq (,)
	This makefile requires GNU Make.
endif

default: dev

dev:
	@go run .

test:
	@go test -covermode=count ./...