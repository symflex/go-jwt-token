package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type (
	Token interface {
		Header() *Header
		Payload() *Payload
		Signature() *Signature
		ToString() string
	}

	JwtToken struct {
		headerSection    Header
		payloadSection   Payload
		signatureSection Signature
	}
)

func CreateToken(header map[string]string, payload map[string]string, signature string) (JwtToken, error) {

	jsonHeaderData, err := json.Marshal(header)
	jsonPayloadData, err := json.Marshal(payload)

	encodedHeader := base64.URLEncoding.EncodeToString([]byte(jsonHeaderData))
	encodedPayload := base64.URLEncoding.EncodeToString([]byte(jsonPayloadData))
	encodedSignature := base64.URLEncoding.EncodeToString([]byte(signature))

	Header := Header{rawData: header, encodedData: encodedHeader}
	Payload := Payload{rawData: payload, encodedData: encodedPayload}
	Signature := Signature{rawData: signature, encodedData: encodedSignature}

	return JwtToken{headerSection: Header, payloadSection: Payload, signatureSection: Signature}, err
}

func (t *JwtToken) Header() *Header {
	return &t.headerSection
}

func (t *JwtToken) Payload() *Payload {
	return &t.payloadSection
}

func (t *JwtToken) Signature() *Signature {
	return &t.signatureSection
}

func (t *JwtToken) ToString() string {
	return fmt.Sprintf("%s.%s.%s", *t.Header().Encoded(), *t.Payload().Encoded(), *t.Signature().Encoded())
}
