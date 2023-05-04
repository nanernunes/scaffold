# Scaffold

[![Go](https://github.com/nanernunes/scaffold/actions/workflows/go.yml/badge.svg)](https://github.com/nanernunes/scaffold/actions/workflows/go.yml)
[![version](https://img.shields.io/github/tag/nanernunes/scaffold.svg)](https://github.com/nanernunes/scaffold/releases/latest)
[![GoDoc](https://godoc.org/github.com/scaffold?status.png)](https://godoc.org/github.com/nanernunes/scaffold)
[![license](https://img.shields.io/github/license/nanernunes/scaffold.svg)](../LICENSE.md)
[![LoC](https://tokei.rs/b1/github/nanernunes/scaffold?category=lines)](https://github.com/nanernunes/scaffold)
[![codecov](https://codecov.io/gh/nanernunes/scaffold/branch/master/graph/badge.svg)](https://codecov.io/gh/nanernunes/scaffold)

Scaffold is a Golang framework that includes everything needed to create database-backed API applications.

## Requirements

- make (Makefile)
- [SQLite](https://sqlite.org/)
- [Go](https://go.dev/)

## Frameworks and libraries

Scaffold comes with:

- [gorilla/mux](https://github.com/gorilla/mux), a powerful HTTP router and URL matcher for building Go web servers
- [swaggo/swag](https://github.com/swaggo/swag), automatically generate RESTful API documentation with Swagger
- [go-orm/gorm](https://github.com/go-gorm/gorm), the fantastic ORM library for Golang, aims to be developer friendly

## Getting Started

1. Install Scaffold at the command prompt:

```bash
sudo curl -L "https://github.com/nanernunes/scaffold/releases/latest/download/scaffold-$(uname -s)-$(uname -m)" -o /usr/local/bin/scaffold
sudo chmod +x /usr/local/bin/scaffold
```

2. At the command prompt, create a new Scaffold application: (where `myapp` is the application name)

```bash
scaffold new myapp
```

3. Change directory to `myapp` and start the web server:

```bash
cd myapp
make dev
```

4. Go to http://localhost:4000 and you'll see the Swagger documentation.

5. Under http://localhost:4000/index there is the default Index (controller/view) home page.

6. Follow the guidelines to start developing your application:

   - [Getting Started with Scaffold](docs/getting-started-with-scaffold.md)

## Contributing

Contributions to Scaffold are welcome and appreciated. If you would like to contribute, please open an issue or submit a pull request.

## License

Scaffold is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
