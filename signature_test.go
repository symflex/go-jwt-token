package jwt

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"testing"
)

var rawSignatureData = "test string"

func TestSignature_Data(t *testing.T) {
	encodedSignatureData := base64.URLEncoding.EncodeToString([]byte(rawSignatureData))
	signature := Signature{rawData: rawSignatureData, encodedData: encodedSignatureData}
	assert.Equal(t, *signature.Data(), rawSignatureData)
}

func TestSignature_Encoded(t *testing.T) {
	encodedSignatureData := base64.URLEncoding.EncodeToString([]byte(rawSignatureData))
	signature := Signature{rawData: rawSignatureData, encodedData: encodedSignatureData}
	assert.Equal(t, *signature.Encoded(), encodedSignatureData)
}
