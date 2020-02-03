# Jwt token
[![Build Status](https://travis-ci.org/symflex/go-jwt.svg?branch=master)](https://travis-ci.org/symflex/go-jwt)
[![Coverage Status](https://coveralls.io/repos/github/symflex/go-jwt/badge.svg)](https://coveralls.io/github/symflex/go-jwt)

```
var header = map[string]string{"key": "value"}
var payload = map[string]string{"key": "value"}
var signature = "signature string"

token := CreateToken(header, payload, signature)
```