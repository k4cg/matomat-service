# Matomat - ReWrite = Matomat as a Service = MaaS

This project reimplements the functionality of the current standalone
[Matomat](https://github.com/k4cg/matomat) application in the form of
a RESTful service. This service then can be consumed by any client.

## Project contents
- `maas-client-tty` (will/should) contain a client implementation that can be
  run on tty
- `maas-server` contains the service server implementation
- `maas.yml` contains the service API definition in OpenAPI 3.0 format

## API definition
See `maas.yml`.

## Easy testing of the API
Use a tool like [Postman](https://www.getpostman.com/) and import the API
definition file. Another alternative would be to use [Swagger Code
Editor](https://editor.swagger.io//#/), import and then use the
"Authorize/Execute/Test" functionality.

## Authentication concept
A client authenticates to the service by passing a username and password (as
well as a validity time in seconds) to the server. The server then returns
a [JWT](https://en.wikipedia.org/wiki/JSON_Web_Token) token that is valid for
the requested amount of seconds (or it this was not given, for the configured
server default). This [JWT](https://en.wikipedia.org/wiki/JSON_Web_Token)
token has to be passed using an `Authorization` header alongside any other
service request (Using the "Bearer token" concept).

### What about RFID/NFC chip based "login"
This functionality has to be implemented by the client. The client would have
to offer a enroll functionality during which the user enters her credentials.
Then a very high token validity time is chosen by the client application
(months?) and a login operation - resulting in a JWT token - is performed
against the service. The received token is then stored on the RFID/NFC chip.
For subsequent "logins" the client only needs to read the token from the
RFID/NFC chip and use it to authenticate its calls to the service.

## Authorization concept
Currently there are only normal users and users with admin rights (admins).
Certain API operations can only be executed by admins (E.g. create user,
delete item, ...).

