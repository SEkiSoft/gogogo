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
	go test -run=. ./store
	go test -run=. ./utils

run:
	@echo Running server
	go run *.go &

clean:
	@echo Cleaning
	rm gogogo
