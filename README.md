# Clean Architecture In Golang

[![BCH compliance](https://bettercodehub.com/edge/badge/Err0r500/Clean-Architecture-in-Golang?branch=master)](https://bettercodehub.com/)

> The purpose of this architecture is to be as flexible as possible in order to develop a project as quickly as possible and to maintain this speed during the whole development lifespan, whatever its size and the changes that will have to be done, for whatever reason.

## Complete Go App using this architecture

[err0r500/go-realworld-clean](https://github.com/err0r500/go-realworld-clean)

## Detailed explanation of this repo

[Hands-On Clean Architecture](https://medium.com/@matthieujacquotdev/clean-architecture-in-golang-7ebafe4c70db)

## Demo

Run `make demo` and read the logs from `src/main_test.go` to see:

- how the execution flows from a layer to another while the dependency inversion rule is strictly observed
- how an Interface can be substituted to another allowing us to plug IN/OUT dependencies
