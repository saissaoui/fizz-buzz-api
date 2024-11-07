# fizz buzz example
REST api exposing fizz buzz endpoint

## What's fizz buzz

The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by ""fizz"", all multiples of 5 by ""buzz"", and all multiples of 15 by ""fizzbuzz"". 
Statistics of the requested data is stored in a redis database . 

## Endpoints :

- health :  Return status = 200 if the server is up
```
GET /health
```
- fizz-buzz : Returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.

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

- Statistics : Retruns an array of requests with the number of times each one was requested .

```
GET /statistics
```
returns : 
```
{
    "stats": [
        {
            "request": "{int1: 2, int2: 3, limit: 20, str1: fizz, str2 : buzz }",
            "count": 13
        },
        {
            "request": "{int1: 2, int2: 3, limit: 20, str1: fiee, str2 : burr }",
            "count": 2
        },
        {
            "request": "{int1: 2, int2: 3, limit: 20, str1: t, str2 : f }",
            "count": 4
        },
        {
            "request": "{int1: 2, int2: 10, limit: 10, str1: Two, str2 : Ten }",
            "count": 1
        }
    ]
}
```
## Dev : 

### Requirements

* Go 1.22;
* [MockGen CMD][1];
* Make;
* Docker

#### To run the server

```sh
make serve
```
#### to run tests

In some cases you will need to regenerate mock files for tests. This is done with the following command:

```sh
make generate-mocks
```

```sh
make tests
```

#### to install docker containers needed for dev env (redis)
```sh
make docker-dev
```

#### to run api as docker container 
```sh
make docker-run
```

[1]: https://github.com/golang/mock
