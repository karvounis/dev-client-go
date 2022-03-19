# Dev-Client-Go

**dev-client-go** is a client library for the Forem (dev.to) [developer api](https://developers.forem.com/api) written in Go. It provides fully typed methods for every operation you can carry out with the current api (beta)(0.9.7).

## Installation

> Go version >= 1.13

```sh
go get github.com/karvounis/dev-client-go
```

### Usage

Import the package and initialize a new client with your auth token(api-key).
To get a token, see the authentication [docs](https://developers.forem.com/api#section/Authentication).
