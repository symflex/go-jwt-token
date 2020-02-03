package jwt

import (
	"encoding/base64"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var sourceHeaderData = map[string]string{"key": "value"}

func TestHeader_Data(t *testing.T) {
	jsonHeaderData, err := json.Marshal(sourceHeaderData)
	assert.NoError(t, err)

	encodedHeaderData := base64.URLEncoding.EncodeToString([]byte(jsonHeaderData))
	h := Header{rawData: sourceHeaderData, encodedData: encodedHeaderData}
	assert.Equal(t, *h.Data(), sourceHeaderData)
}

func TestHeader_Encoded(t *testing.T) {
	jsonHeaderData, err := json.Marshal(sourceHeaderData)
	assert.NoError(t, err)

	encodedHeaderData := base64.URLEncoding.EncodeToString([]byte(jsonHeaderData))
	h := Header{rawData: sourceHeaderData, encodedData: encodedHeaderData}
	assert.Equal(t, *h.Encoded(), encodedHeaderData)
}
