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

func randBelow(n int) int {
	rand.Seed(time.Now().UnixNano())
	if n <= 0 {
		fmt.Println("Upper bound must be positive.")
		os.Exit(1)
	}
	return rand.Intn(n)
}
func tokeBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := cryand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func tokenHex(n int) (string, error) {
	b := make([]byte, n)
	_, err := cryand.Read(b)
	if err != nil {
		return "", err
	}
	s := hex.EncodeToString(b)
	return s, nil
}
func tokenUrlSafe(n int) (string, error) {
	data, err := tokeBytes(n)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}
