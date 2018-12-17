// Package gosecrets generate cryptographically strong pseudo-random numbers suitable for
// managing secrets such as account authentication, tokens, and similar.
package gosecrets

import (
	cryand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math/rand"
	"time"
)

// RandBelow function return random int in the range [0, n]
func RandBelow(n int) (int, error) {
	rand.Seed(time.Now().UnixNano())
	if n <= 0 {
		return 0, errors.New("n int must be more than 0")
	}
	return rand.Intn(n) + n, nil
}

// TokeBytes return bytes of len(n) with error nil and nil of bytes and error if n less than 0
func TokeBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := cryand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// TokenHex return random string in hex for len(n) or empty string with error if n is less than 0
func TokenHex(n int) (string, error) {
	b := make([]byte, n)
	_, err := cryand.Read(b)
	if err != nil {
		return "", err
	}
	s := hex.EncodeToString(b)
	return s, nil
}

//TokenURLSafe eturn a random URL-safe text string, in Base64 encoding.
// The string has n random bytes.  If n is less or equal to 0
// return empty string and error
func TokenURLSafe(n int) (string, error) {
	data, err := TokeBytes(n)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}
