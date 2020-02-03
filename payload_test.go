package jwt

import (
	"encoding/base64"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sourcePayloadData = map[string]string{"key": "value"}

func TestPayload_Data(t *testing.T) {
	jsonPayloadData, err := json.Marshal(sourcePayloadData)
	assert.NoError(t, err)

	encodedPayloadData := base64.URLEncoding.EncodeToString([]byte(jsonPayloadData))
	p := Payload{rawData: sourcePayloadData, encodedData: encodedPayloadData}
	assert.Equal(t, *p.Data(), sourcePayloadData)
}

func TestPayload_Encoded(t *testing.T) {
	jsonPayloadData, err := json.Marshal(sourcePayloadData)
	assert.NoError(t, err)

	encodedPayloadData := base64.URLEncoding.EncodeToString([]byte(jsonPayloadData))
	p := Payload{rawData: sourcePayloadData, encodedData: encodedPayloadData}
	assert.Equal(t, *p.Encoded(), encodedPayloadData)
}
