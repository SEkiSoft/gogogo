run: install
	./node_modules/.bin/webpack

run-fullmap: install
	./node_modules/.bin/webpack -d

install: package.json
	npm install

check-style:
	./node_modules/.bin/semistandard

clean:
	@echo Cleaning Client
	rm -rf node_modules
	rm public/*
