# Scaffold

[![Go](https://github.com/nanernunes/scaffold/actions/workflows/go.yml/badge.svg)](https://github.com/nanernunes/scaffold/actions/workflows/go.yml)
[![version](https://img.shields.io/github/tag/nanernunes/scaffold.svg)](https://github.com/nanernunes/scaffold/releases/latest)
[![GoDoc](https://godoc.org/github.com/scaffold?status.png)](https://godoc.org/github.com/nanernunes/scaffold)
[![license](https://img.shields.io/github/license/nanernunes/scaffold.svg)](../LICENSE.md)
[![LoC](https://tokei.rs/b1/github/nanernunes/scaffold?category=lines)](https://github.com/nanernunes/scaffold)
[![codecov](https://codecov.io/gh/nanernunes/scaffold/branch/master/graph/badge.svg)](https://codecov.io/gh/nanernunes/scaffold)

Scaffold is a Golang framework that automates the creation of Swaggered APIs.

## Requirements

- [gorilla/mux](https://github.com/gorilla/mux)
- [swaggo/swag](https://github.com/swaggo/swag)
- [go-orm/gorm](https://github.com/go-gorm/gorm)

## Installation

To install Scaffold, run the following command:

`go get github.com/nanernunes/scaffold`

## Getting Started with Scaffold

```go
scaffold new blog
scaffold generate User name:string age:int
```

## Folder Structure

```bash
.
├── Makefile
├── api
│   ├── api.go
│   └── controllers
│       ├── index.go
│       └── users.go
├── app
│   ├── models
│   │   └── user.go
│   └── services
│       └── user.go
├── cmd
│   └── api.go
├── config
│   ├── api.go
│   ├── application.go
│   └── config.go
├── go.mod
└── go.sum

7 directories, 12 files
```

## Running the Server

```sh
make dev
```

## Contributing

Contributions to Scaffold are welcome and appreciated. If you would like to contribute, please open an issue or submit a pull request.

## License

Scaffold is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
