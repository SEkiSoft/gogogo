# GoGoGo
[![Build Status](https://travis-ci.org/DavidLu1997/gogogo.svg?branch=master)](https://travis-ci.org/DavidLu1997/gogogo)

GoGoGo is an online multiplayer implementation of the ancient board game Go, written in Go with a machine learning AI player (inspired by AlphaGo).

## Building
Before you can build this, make sure you have the following dependencies installed:
* [Go compiler (the latest version)](https://golang.org/doc/install)
* [Git (at least version 1.7)](https://git-scm.com/downloads)

### Linux and Mac OS
1. Make sure you have your [GOPATH](https://golang.org/doc/code.html) environment variable set.
2. `git clone --recursive` this repository to `$GOPATH/src/github.com/SEkiSoft/gogogo`.
3. Run `make build` in that directory.

### Windows
1. Make sure you have your [GOPATH](https://golang.org/doc/code.html) environment variable set.
2. `git clone --recursive` this repository to `%GOPATH%\src\github.com\SEkiSoft\gogogo`.
3. Run `go build` in that directory.

### Common Build Errors:
* If there are multiple errors, ensure your `GOPATH` is set correctly, according to [Go documentation](https://golang.org/doc/code.html).
* This repository uses git submodules. If you did not clone the repository using `git clone --recursive`, make sure you initialize the submodules by running `git submodule update --init`.

## Running

### Linux and Mac OS
1. Install `mysql`
2. Create an user with username `gouser` and password `gotest`
3. Create a database called `gogogo`
4. Ensure the database is listening on `localhost:3306` via TCP
5. `make run`
6. The server is now running on `localhost:3030`

### Windows
1. Install [MySQL Server](http://dev.mysql.com/downloads/windows/installer/5.7.html)
    * MySQL Workbench is highly recommended for development.
2. Create an user with username `gouser` and password `gotest`
3. Create a database called `gogogo`
4. Ensure the database is listening on `localhost:3306` via TCP
5. `go run gogogo.go`
6. The server is now running on `localhost:3030`

All settings can be changed in `config/config.json`
