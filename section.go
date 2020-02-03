package jwt

type (
	Section interface {
		Data() *map[string]string
		Encoded() *string
	}
)
