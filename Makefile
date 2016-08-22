build:
	@echo Building server and client
	go build

check-style:
	@echo Checking style
	gofmt -w .

test:
	@echo Running unit tests
	go test -run=. ./api
	go test -run=. ./model
	go test -run=. ./utils

run:
	@echo Running server
	go run *.go &

stop:
	@echo Stopping server

	@for PID in $$(ps -ef | grep "[g]o run" | awk '{ print $$2 }'); do \
		kill $$PID; \
	done

	@for PID in $$(ps -ef | grep "[g]o-build" | awk '{ print $$2 }'); do \
		kill $$PID; \
	done

clean:
	@echo Cleaning
	rm gogogo
