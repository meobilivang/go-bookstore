# BookStore RESTful API

## Tech Stack

- `Go 1.18`
- `MYSQL`

## Set Up Environment

### Project

- Initialize `go.mod`:
```bash
$ go mod init github.com/meobilivang/go-bookstore
```

- Install following packages
```bash
$ go get github.com/jinzhu/gorm
---
$ go get github.com/jinzhu/gorm/dialects/mysql
---
$ go get github.com/gorilla/mux
```

### Database

- Replace `<hostname>` under `app.go` with the hostname of running MYSQL instance.

## Run Project

- Navigate to `main.go`:

```bash
$ cd cmd
```

- Start service:

```bash
$ go run main.go
```