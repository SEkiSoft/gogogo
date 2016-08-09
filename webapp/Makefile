run: install
	./node_modules/.bin/webpack

run-fullmap: install
	./node_modules/.bin/webpack -d

install: package.json
	npm install

check-style:
	eslint --ext ".jsx" --ignore-pattern node_modules --quiet .

clean:
	@echo Cleaning Client
	rm -rf node_modules
	rm public/*
