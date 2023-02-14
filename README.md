# TP Examples

This repository has some examples about how to code with third-party go packages.

## Topics

- JSON Web Tokens -> [golang-jwt](https://github.com/golang-jwt/jwt), check the section *jwt-tool*.

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