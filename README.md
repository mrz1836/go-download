# go-download
**go-download** takes a JSON configuration of urls, paths and file names and downloads all the files locally

| | | | | | | |
|-|-|-|-|-|-|-|
| ![License](https://img.shields.io/github/license/mrz1836/go-download.svg?style=flat) | [![Report](https://goreportcard.com/badge/github.com/mrz1836/go-download?style=flat)](https://goreportcard.com/report/github.com/mrz1836/go-download)  | [![Codacy Badge](https://api.codacy.com/project/badge/Grade/6d55b5c2f8494371932f4ceb2173934f)](https://www.codacy.com/app/mrz1818/go-download?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=mrz1836/go-download&amp;utm_campaign=Badge_Grade) |  [![Build Status](https://travis-ci.com/mrz1836/go-download.svg?branch=master)](https://travis-ci.com/mrz1836/go-download)   |  [![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat)](https://github.com/RichardLitt/standard-readme) | [![Release](https://img.shields.io/github/release-pre/mrz1836/go-download.svg?style=flat)](https://github.com/mrz1836/go-download/releases) | [![GoDoc](https://godoc.org/github.com/mrz1836/go-download?status.svg&style=flat)](https://godoc.org/github.com/mrz1836/go-download) |

## Table of Contents
- [Installation](#installation)
- [Documentation](#documentation)
- [Examples & Tests](#examples--tests)
- [Benchmarks](#benchmarks)
- [Code Standards](#code-standards)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

## Installation

**go-download** requires a [supported release of Go](https://golang.org/doc/devel/release.html#policy).
```bash
$ go get -u github.com/mrz1836/go-download
```

## Documentation
You can view the generated [documentation here](https://godoc.org/github.com/mrz1836/go-download).

## Examples & Tests
All unit tests run via [Travis CI](https://travis-ci.com/mrz1836/go-download) and uses [Go version 1.12.x](https://golang.org/doc/go1.12). View the [deployment configuration file](.travis.yml).
```bash
$ cd ../go-download
$ go test ./... -v
```

## Benchmarks
Run the Go [benchmarks](download_test.go):
```bash
$ cd ../go-download
$ go test -bench . -benchmem
```

## Code Standards
Read more about this Go project's [code standards](CODE_STANDARDS.md).

## Usage
- View the [benchmarks](download_test.go)
- View the [tests](download_test.go)

Basic implementation (using JSON File):
1) store a local JSON file like example.JSON
2) change the location in download.go
3) run download package:
```bash
$ cd ../go-download
$ go run download.go
```

## Maintainers

[@MrZ1836](https://github.com/mrz1836)

## Contributing

View the [contributing guidelines](CONTRIBUTING.md) and follow the [code of conduct](CODE_OF_CONDUCT.md).

Support the development of this project 🙏

[![Donate](https://img.shields.io/badge/donate-bitcoin-brightgreen.svg)](https://mrz1818.com/?tab=tips&af=go-download)

## License

![License](https://img.shields.io/github/license/mrz1836/go-download.svg?style=flat)
