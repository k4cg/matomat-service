# MaaS - Server Implementation
Implementation of Matomat As A Service using Golang.
Work in progress, feel free to contribute.

## Biggest TODOs
- Properly set MQTT message format

(for not so big TODOs simply grep for TODO, the result will speak for itself ^^)

## Requirements
- golang version 1.7+
- sqlite3
- ~440MB RAM (measured using `memusg`=> see "Reference Links")

## Installation
0. Pull all go dependencies using `go get ./...`
1. Create a sqlite3 db at the desired location using `sqlite3 matomat.db < schema.sql; sqlite3 matomat.db < initial_data.sql`
2. Create SSL certificate data (public and private cert parts)
3. Create a `config.yml` from `config.yml.template` and fill will real values
4. Run the server :-)

## Initial user info
Use `admin` with password `admin` for initial login. Create new user and delete user `admin` afterwards.
(Hash for the password in `initial_data.sql` was created with bcrypt using 10 rounds)

## Shutting down the server
Send a SIGTERM or SIGINT, server will then shutdown after a grace period configured in `config.yml`.

## Local testing
Easy creation of a self signed certificate for local TESTING.
```
go run $GOROOT/src/crypto/tls/generate_cert.go --rsa-bits 4096 --host 127.0.0.1,::1,localhost --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
```
## Event Dispatcher

Example MQTT Messages

```
'matomat;item-consumed;1;2;Bier;100;1'
'matomat;item-consumed;1;1;Club Mate;100;1'
```

## Reference Links
* [awesome-go](https://github.com/avelino/awesome-go)
* [config](https://github.com/olebedev/config)
* [mqtt](https://eclipse.org/paho/clients/golang/)
* [sqlite](https://github.com/mattn/go-sqlite3)
* [webservice in go](https://auth0.com/blog/authentication-in-golang/)
* The first basic service stub was originally generated using [Swagger Codegen](https://github.com/swagger-api/swagger-codegen.git), but has moved far from it.
* [memusg - measure peak memory usage](https://github.com/jhclark/memusg)
