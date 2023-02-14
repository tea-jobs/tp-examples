# TP Examples

This repository has some examples about how to code with third-party go packages.

## Topics

- JSON Web Tokens -> [golang-jwt](https://github.com/golang-jwt/jwt), check the section *jwt-tool* or *http-server*.

### jwt-tool

To build the jwt-tool:
```
make build-jwt-tool
```
To generate a token string:
```
MODE=-g DATA=foo KEY=bar make run-jwt-tool
--->
Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoiZm9vIn0.t3xuLjTwvdgNaliBdCd_h04CH5mqc7dx4kObDJWTZN8
```
To parse a token string:
```
MODE=-p DATA=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoiZm9vIn0.t3xuLjTwvdgNaliBdCd_h04CH5mqc7dx4kObDJWTZN8 KEY=bar make run-jwt-tool
--->
Data: foo
```

### http-server

To build the http-server:
```
make build-http-server
```
To generate a token string:
```
curl -X POST -d "data=abc&key=foo" localhost:8080/generateToke
--->
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoiYWJjIn0.i-_8FgJvSCzS9B3l0-ELoNPQKSRXRmld27wM85j5CaE
```
To parse a token string:
```
curl -H "Content-Type: application/json" -X "POST" -d '{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoiYWJjIn0.i-_8FgJvSCzS9B3l0-ELoNPQKSRXRmld27wM85j5CaE","key":"foo"}' localhost:8080/parseToken
--->
Data: abc
```