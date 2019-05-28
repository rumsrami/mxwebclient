package rand

import (
	"crypto/rand"
	"encoding/base64"
)

const fixedSizeToken = 32

// Bytes help us generate n random bytes
func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

// string will generate a random byte slice of size nBytes and then
// return a string that is the base64 URL encoded version of
// that byte slice.
func stringify(nBytes int) (string, error){
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// Gen32ByteToken generates a random 32 byte string
func Gen32ByteToken() (string, error) {
	token, err := stringify(fixedSizeToken)
	if err != nil {
		return "", nil
	}
	return token, nil
}