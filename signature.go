package jwt

type (
	Signature struct {
		rawData     string
		encodedData string
	}
)

func (s *Signature) Encoded() *string {
	return &s.encodedData
}

func (s *Signature) Data() *string {
	return &s.rawData
}
