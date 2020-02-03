package jwt

type (
	Payload struct {
		rawData     map[string]string
		encodedData string
	}
)

func (p *Payload) Data() *map[string]string {
	return &p.rawData
}

func (p *Payload) Encoded() *string {
	return &p.encodedData
}
