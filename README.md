# Jwt token
[![Build Status](https://travis-ci.org/symflex/go-jwt-token.svg?branch=master)](https://travis-ci.org/symflex/go-jwt-token)
[![Coverage Status](https://coveralls.io/repos/github/symflex/go-jwt-token/badge.svg?branch=master)](https://coveralls.io/github/symflex/go-jwt-token?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/symflex/go-jwt-token)](https://goreportcard.com/report/github.com/symflex/go-jwt-token)

```
var header = map[string]string{"key": "value"}
var payload = map[string]string{"key": "value"}
var signature = "signature string"

token := CreateToken(header, payload, signature)
```