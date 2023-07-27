# Angus: Love this. Great structure, well thought out. Impressive to see you thinking of this stuff.

ifneq ("$(wildcard .env)","")
	# Angus: fun fact: not every environment variable that can be expressed in a .env file can be included in a Makefile.
	# This include directive will fail if you have any multi-line environment variables (e.g. PEM strings). This makes
	# me very sad and I've wasted many hours trying to find a way around it. Be warned!
	#
	# Naofel: I conducted a small test to verify your claim since I had never used multiline variables
	# in my env file before, and you were right. Although you can use backslashes at the end of each line to create a
	# multiline variable, it introduces unwanted spaces when there are newlines. Consequently, I discovered that the
	# only effective approach is to run a bash script beforehand, although it does not directly use the env file.
	# This discovery is quite intriguing! 🧐 Thank you for the tips! 😀 https://gist.github.com/naofel1/bb9b497b11707be79db381a00c81e17a
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
	# Angus: The Go race detector is an incredible tool. It will crash the program and warn you if it detects a data
	# race in your code. This has helped me catch many concurrency errors I've made. ALWAYS enable it for tests and
	# local development. Avoid it in production and benchmarks, because it slows the program down significantly.
	#
	# Naofel: Ok I will always use it now for tests and local development. Thank you for the tips! 😀
	go test -race ./...

test.v:
	go test -race ./... -v

###########
#  Local  #
###########

local:
	go run -race cmd/monolith/main.go

local.up:
	go run -race cmd/monolith/main.go

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
