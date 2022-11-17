# fizz buzz example
REST api exposing fizz buzz endpoint

## What's fizz buzz

The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by ""fizz"", all multiples of 5 by ""buzz"", and all multiples of 15 by ""fizzbuzz"". 

## Endpoints :

```
GET /health
```
Return status = 200 if the server is up

```
POST /fizz-buzz
```
with body :

```
{
    "int1":2,
    "int2":3,
    "limit":20,
    "str1":"r",
    "str2":"buzz"
}
```

Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

## Dev : 

- To run the server

```sh
make serve
```
- to run tests

```sh
make test
```

- to run as docker container

```sh
make docker
```