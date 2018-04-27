# MaaS - Server Implementation
Implementation of Matomat As A Service using Golang.
Work in progress, feel free to contribute.

## Biggest TODOs
- Properly set MQTT message format
- Implement service stats

## Prerequisites
- golang version 1.7+
- sqlite3

## Installation
0. Pull all go dependencies using `go get ./...`
1. Create a sqlite3 db at the desired location using `sqlite3 matomat.db < schema.sql; sqlite3 matomat.db < initial_data.sql`
2. Create SSL certificate data (public and private cert parts)
3. Create a `config.yml` from `config.yml.template` and fill will real values
4. Run the server :-)

## Initial user info
Use `admin` with password `admin` for initial login. Create new user and delete user `admin` afterwards.
(Hash for the password in `initial_data.sql` was created with bcrypt using 10 rounds)

## Reference Links
[awesome-go](https://github.com/avelino/awesome-go)
[config](https://github.com/olebedev/config)
[mqtt](https://eclipse.org/paho/clients/golang/)
[sqlite](https://github.com/mattn/go-sqlite3)
[webservice in go](https://auth0.com/blog/authentication-in-golang/)
The first basic service stub was originally generated using [Swagger Codegen](https://github.com/swagger-api/swagger-codegen.git), but has moved far from it.
