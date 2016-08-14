# GoGoGo
[![Build Status](https://travis-ci.org/DavidLu1997/gogogo.svg?branch=master)](https://travis-ci.org/DavidLu1997/gogogo)

GoGoGo is an online multiplayer implementation of the ancient board game Go, written in Go with a machine learning AI player (inspired by AlphaGo).

## Building
Before you can build this, make sure you have the following dependencies installed:
* [Go compiler (the latest version)](https://golang.org/doc/install)
* [Node.js with npm](https://nodejs.org/en/download/)
* [Git (at least version 1.7)](https://git-scm.com/downloads)

### Linux and Mac OS
1. Make sure you have your [GOPATH](https://golang.org/doc/code.html) environment variable set.
2. `git clone --recursive` this repository to `$GOPATH/src/github.com/davidlu1997/gogogo`.
3. Run `make build` in that directory.

### Windows
1. Ensure you have the following additional dependencies installed:
	* GNU Make
2. Make sure you have your [GOPATH](https://golang.org/doc/code.html) environment variable set.
3. `git clone --recursive` this repository to `%GOPATH%\src\github.com\davidlu1997\gogogo`.
4. Run `make build` in that directory.

### Common Build Errors:
* If there are multiple errors, ensure your `GOPATH` is set correctly, according to [Go documentation](https://golang.org/doc/code.html).
* This repository uses git submodules. If you did not clone the repository using `git clone --recursive`, make sure you initialize the submodules by running `git submodule update --init`.

## Running
*To be completed.*
