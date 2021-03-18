.PHONY: all
# TODO: add env to bin name
BINARY=projects

LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

install:
	@echo "Installing air" 
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
	make dependencies
	make lint-prepare

integration-test: 
	docker-compose -f docker-compose.yml -f docker-compose.test.yml up --abort-on-container-exit
	docker-compose -f docker-compose.yml -f docker-compose.test.yml down

run-integration-test: 
	go test --tags="integration" -v ./...

run-unit-test: migrate-up unit-test migrate-down
unit-test:
	go test --tags="unit" -v ./...

test: 
	go test --tags="unit integration" -cover -covermode=atomic ./...

build: dependencies build-projects-cmd

# TODO: use tags
build-projects-cmd:
	go build -tags release -o ./build/${BINARY} ./cmd/projects/main.go

dependencies:
	go mod download

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

# send binary name to dockerfile?
docker:
	docker build -t projects .

start:
	docker-compose up -d

stop:
	docker-compose down

serve:
	air -c .air.conf

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s latest

# read more about linting in: https://github.com/golangci/golangci-lint#config-file - maybe use config file?
lint:
	./bin/golangci-lint run


# Note: temporary until we move migrations to file
dbname := $(or $(POSTGRES_DB),postgres)
dbuser := $(or $(POSTGRES_USER),postgres)
dbpassword := $(or $(POSTGRES_PASSWORD),)
dbhost :=  $(or $(POSTGRES_HOST), localhost)
dbport :=  $(or $(POSTGRES_PORT), 5432)

migrate-up:
	migrate -path migrations -database "postgresql://${dbuser}:${dbpassword}@${dbhost}:${dbport}/${dbname}?sslmode=disable" -verbose up

migrate-down:
	migrate -path migrations -database "postgresql://${dbuser}:${dbpassword}@${dbhost}:${dbport}/${dbname}?sslmode=disable" -verbose down

