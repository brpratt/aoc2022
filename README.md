# aoc2022

Advent of Code 2022

## Testing

```
go test ./...
```

## Running

This repo builds all the solutions into a single binary. To compile and run:

```
go build
./aoc2022 <day>
```

Alternatively:

```
go run main.go <day>
```

## AWK?

Some of the days have a solution written in [AWK](https://en.wikipedia.org/wiki/AWK). To run the AWK solutions:

```
./dayXY/dayXY.awk < ./dayXY/input.txt
```

## Dev Container

Have Docker and VS Code? This repo includes a [dev container](https://code.visualstudio.com/docs/devcontainers/containers) configuration.

The dev container includes all the tools necessary to build and run the Go and AWK solutions.
