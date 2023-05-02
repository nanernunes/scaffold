.PHONY:scaffold

TAG:=$(shell git describe --tags)

MAKE:=$(shell which make)
DOCKER-COMPOSE:=$(shell which docker-compose)

GOCMD:=$(shell which go)

%:
	@:

test:
	@$(GOCMD) test -v ./...

coverage:
	@$(GOCMD) test -v ./... -coverprofile=coverage.txt -covermode=atomic
	@$(GOCMD) tool cover -func=coverage.txt
	@$(GOCMD) tool cover -html=coverage.txt

deps:
	@$(GOCMD) mod download

cli:
	@$(GOCMD) run cmd/scaffold/scaffold.go $(filter-out $@,$(MAKECMDGOALS))

build:
	@sed -Eir "s@(/scaffold) .*@\1 $(TAG)@" pkg/templates/go.mod.tmpl
	@$(GOCMD) build -o scaffold cmd/scaffold/scaffold.go
