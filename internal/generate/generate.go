package generate

import (
	"crypto/rand"
	"fmt"
	"math/big"
	mathrand "math/rand"
	"time"
)

const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func init() {
	mathrand.Seed(time.Now().UnixNano())
}

// CryptoRandom generates a cryptographically-safe random number
func CryptoRandom(n int) (string, error) {
	if n <= 0 {
		return "", nil
	}

	bytes := make([]byte, n)
	for i := range bytes {
		bigint, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphanum))))
		if err != nil {
			return "", fmt.Errorf("couldn't generate random string of len %d: %w", n, err)
		}

		bytes[i] = alphanum[bigint.Int64()]
	}

	return string(bytes), nil
}

// MathRandom generates a random string that does not need to be cryptographically secure
// This is preferred to CryptoRandom when you don't need the cryptographic security as it is
// not a drain on the entropy pool.
func MathRandom(n int) string {
	if n <= 0 {
		return ""
	}

	bytes := make([]byte, n)
	for i := range bytes {
		//nolint:gosec // We purposely use mathrand to avoid draining the entropy pool
		j := mathrand.Int31n(int32(len(alphanum)))
		bytes[i] = alphanum[j]
	}

	return string(bytes)
}
