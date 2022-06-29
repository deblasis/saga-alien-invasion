# Alex's Saga Alien Invasion

[![Build and Tests](https://github.com/deblasis/saga-alien-invasion/actions/workflows/tests.yml/badge.svg)](https://github.com/deblasis/saga-alien-invasion/actions/workflows/tests.yml)

This repository contains the source-code of the `saga-alien-invasion` CLI.

## Demo

https://user-images.githubusercontent.com/29378614/176482487-7c736fbd-0042-42fb-b998-a8c31bb99629.mp4


## Usage

```
saga-alien-invasion [number of aliens invading] [flags]
```

### Options

```
  -h, --help                     help for saga-alien-invasion
      --mapfile string           the file containing the map of the 🌍, [relative path] (default "map.txt")
      --maxTurns int             the number of turns before the mothership calls the aliens back home and the program ends (default 10000)
      --numAliensForBattle int   the number aliens required to start a battle that will end up destroying a city (default 2)
      --verbose                  if True, sets the loglevel to show DEBUG messages
```

### Documentation

In this repository we leverage GoDoc and Docker, in order to view a self-hosted GoDoc site, run:

```sh
make godoc
```
and navigate to http://localhost:6060/pkg/github.com/deblasis/saga-alien-invasion/

### Run the application

You have a few options, Dockerized executable or compiled locally

#### Dockerized

You just have to build the container image with

```sh
make docker/build
```
and then execute the command

```sh
docker run deblasis-saga-alien-invasion saga-alien-invasion [number of aliens invading] [flags]
```

If you want to pass in a different map that's on your filesystem you can do that using a volume, something like:

```sh
docker run -v $(pwd)/anothermap.txt:/go/src/github.com/deblasis/saga-alien-invasion/map.txt deblasis-saga-alien-invasion saga-alien-invasion [number of aliens invading] [flags]
```

if you know your way around Docker there are other ways too! Feel free to PR a documentation change if this step can be simplified/improved in your opinion ;)

#### Compiled

```sh
make build
./saga-alien-invasion [number of aliens invading] [flags]
```

#### Run from source with logging for debug/dev purposes

```sh
go run ./... X --verbose
```

where **X** is the number of aliens you want to spawn

#### Tests

```sh
make test
```

and

```sh
make test/nocache
```

to make sure that there is no smoke and mirrors with the random stuff

##### Coverage

```sh
make test/cover
```

and then navigate to http://localhost:8080/
