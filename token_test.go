package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var header = map[string]string{"key": "value"}
var payload = map[string]string{"key": "value"}
var signature = "signature string"
var tokenCompleted = "eyJrZXkiOiJ2YWx1ZSJ9.eyJrZXkiOiJ2YWx1ZSJ9.c2lnbmF0dXJlIHN0cmluZw=="

func TestCreateToken(t *testing.T) {
	token, err := CreateToken(header, payload, signature)
	assert.NoError(t, err)
	assert.ObjectsAreEqual(token, JwtToken{})
}

func TestJwtToken_Header(t *testing.T) {
	token, err := CreateToken(header, payload, signature)
	assert.NoError(t, err)
	assert.ObjectsAreEqual(token.Header(), Header{})
}

func TestJwtToken_Payload(t *testing.T) {
	token, err := CreateToken(header, payload, signature)
	assert.NoError(t, err)
	assert.ObjectsAreEqual(token.Payload(), Payload{})
}

func TestJwtToken_Signature(t *testing.T) {
	token, err := CreateToken(header, payload, signature)
	assert.NoError(t, err)
	assert.ObjectsAreEqual(token.Signature(), Signature{})
}

func TestJwtToken_String(t *testing.T) {
	token, err := CreateToken(header, payload, signature)
	assert.NoError(t, err)
	assert.Equal(t, token.ToString(), tokenCompleted)
}
