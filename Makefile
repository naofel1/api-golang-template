ifneq ("$(wildcard .env)","")
  include .env
  export
else
  $(shell cp .env.example .env)
  include .env
  export
endif

##################################################
#                  Golang                        #
##################################################

############
#  Linter  #
############

lint:
	$(shell go env GOPATH)/bin/golangci-lint run -c configs/.golangci.yml

lint.install:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.51.1

lint.docker:
	docker run --rm -v $(shell pwd):/app -w /app golangci/golangci-lint:v1.51.a golangci-lint run -v

lint.revive:
	$(shell go env GOPATH)/bin/revive --formatter stylish ./...

lint.revive.install:
	go install github.com/mgechev/revive@latest

lint.revive.docker:
	docker run --rm -v $(shell pwd):/var/api ghcr.io/mgechev/revive:1.2.3 --formatter stylish ./var/api/...

###################
#  Documentation  #
###################

docs:
	$(shell go env GOPATH)/bin/swag init --dir cmd/monolith -g main.go --parseDependency --parseInternal -o ./api/rest

docs.lint:
	$(shell go env GOPATH)/bin/swag fmt ./..

docs.install:
	go install github.com/swaggo/swag/cmd/swag@latest

############################
#  Certificate Generation  #
############################

certs: certs.student certs.admin

certs.student:
	openssl genrsa -out configs/certs/student.pem 2048 && \
	openssl rsa -in configs/certs/student.pem -outform PEM -pubout -out configs/certs/student_public.pem

certs.admin:
	openssl genrsa -out configs/certs/admin.pem 2048 && \
	openssl rsa -in configs/certs/admin.pem -outform PEM -pubout -out configs/certs/admin_public.pem

#############
#  Mocking  #
#############

mocks:
	$(shell go env GOPATH)/bin/mockery --all --inpackage  --dir=internal/service --output=internal/service

mocks.install:
	go install github.com/vektra/mockery/v2@latest

#############
#  Ent ORM  #
#############

ent:
	go generate ./internal/ent

migrate:
	go run -mod=mod internal/ent/migrate/main.go create_users

diagram:
	$(shell go env GOPATH)/bin/enter ./internal/ent/schema

#######################
#  Integration tests  #
#######################

test:
	go test ./...

test.v:
	go test ./... -v

###########
#  Local  #
###########

local:
	go run cmd/monolith/main.go

local.up:
	go run cmd/monolith/main.go

##################
#  Local Docker  #
##################

local.docker:
	docker-compose up --detach
local.docker.up:
	docker-compose up --detach
local.docker.up.build:
	docker-compose up --detach --build
local.docker.logs:
	docker-compose logs --follow --tail 1000
local.docker.stop:
	docker-compose stop
local.docker.down:
	docker-compose down
local.docker.ps:
	docker-compose ps
local.docker.clean:
	docker-compose down --rmi all --volumes

.PHONY: docs mocks test ent migrate diagram
