package jwt

type (
	Header struct {
		rawData     map[string]string
		encodedData string
	}
)

func (h *Header) Data() *map[string]string {
	return &h.rawData
}

func (h *Header) Encoded() *string {
	return &h.encodedData
}
