package gosecrets

import (
	cryand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func RandBelow(n int) int {
	rand.Seed(time.Now().UnixNano())
	if n <= 0 {
		fmt.Println("Upper bound must be positive.")
		os.Exit(1)
	}
	return rand.Intn(n)
}
func TokeBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := cryand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func TokenHex(n int) (string, error) {
	b := make([]byte, n)
	_, err := cryand.Read(b)
	if err != nil {
		return "", err
	}
	s := hex.EncodeToString(b)
	return s, nil
}
func TokenUrlSafe(n int) (string, error) {
	data, err := TokeBytes(n)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}
