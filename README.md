# GoGoGo
[![Build Status](https://travis-ci.org/DavidLu1997/gogogo.svg?branch=master)](https://travis-ci.org/DavidLu1997/gogogo)

GoGoGo is an online multiplayer implementation of the ancient board game Go, written in Go with a machine learning AI player (inspired by AlphaGo).

## Building
Before you can build this, make sure you have the following dependencies installed:
* [Go compiler (the latest version)](https://golang.org/doc/install)
* [Node.js with npm](https://nodejs.org/en/download/)

### Linux and Mac OS
1. Make sure you have your [GOPATH](https://golang.org/doc/code.html) environment variable set.
2. `git clone` this repository to `$GOPATH/src/davidlu1997/gogogo`.
3. Run `make build` in that directory.

### Windows
1. Ensure you have the following additional dependencies installed:
	* GNU Make
	* Git
2. Make sure you have your [GOPATH](https://golang.org/doc/code.html) environment variable set.
3. `git clone` this repository to `%GOPATH%\src\davidlu1997\gogogo`.
4. Run `make build` in that directory.


## Running
*To be completed.*


## Contributing
All contributions are welcome to this project! This is the workflow for this project:
* Fork this repository
* Create a topic branch to work under based on the `development` branch
* Ensure your code complies with the following styles:
	* [gofmt](https://golang.org/cmd/gofmt/) for all Go code
	* [semistandard](https://github.com/Flet/semistandard) for all JavaScript code
* Make sure your branch builds (Travis CI will ensure this!)
* Submit a pull request, when completed, against `development` branch
* Get approval from at least two core developers
