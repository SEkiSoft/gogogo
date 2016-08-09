BUILD_WEBAPP_DIR = ./webapp

check-style:
	@echo Checking style

	gofmt -w .

	cd $(BUILD_WEBAPP_DIR) && $(MAKE) check-style

build:
	@echo Building

	go build

	cd $(BUILD_WEBAPP_DIR) && $(MAKE) run

run-server:
	@echo Running Server

	go run *.go &

run-client:
	@echo Running Client

	cd $(BUILD_WEBAPP_DIR) && $(MAKE) run

run-client-fullmap:
	@echo Running Client with Full Source Map

	cd $(BUILD_WEBAPP_DIR) && $(MAKE) run-fullmap

run: run-server run-client

run-fullmap: run-server run-client-fullmap

clean:
	@echo Cleaning

	rm gogogo

	cd $(BUILD_WEBAPP_DIR) && $(MAKE) clean
