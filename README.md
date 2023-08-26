# overlap-checker
This is a Go service to check if two given date ranges overlap or not.

This has one endpoint which takes two date ranges as input and returns true if they overlap else false.

Endpoint: `/check_overlap`

Example 1:
```
curl --location --request POST 'localhost:8080/check_overlap' \
--header 'Content-Type: application/json' \
--data-raw '{
  "range1": {
    "start": "2023-08-01T00:00:00Z",
    "end": "2023-08-15T00:00:00Z"
  },
  "range2": {
    "start": "2023-08-10T00:00:00Z",
    "end": "2023-08-20T00:00:00Z"
  }
}'
```
In this case the response will be:
```
{
    "overlap": true
}
```

Example 2:
```
curl --location --request POST 'localhost:8080/check_overlap' \
--header 'Content-Type: application/json' \
--data-raw '{
  "range1": {
    "start": "2023-08-01T00:00:00Z",
    "end": "2023-08-15T00:00:00Z"
  },
  "range2": {
    "start": "2023-08-15T00:00:00Z",
    "end": "2023-08-20T00:00:00Z"
  }
}'
```
In this case the response will be:
```
{
    "overlap": false
}
```

## How to run
The pre-requisite is to have `go` installed.
We can install `go` by following the instructions [here](https://golang.org/doc/install).
We can run this service using the below command.
```
make run
```

## How to run tests
We can run the tests using the below command.
```
make test
```

## How to build
We can build the binary using the below command.
```
make build
```