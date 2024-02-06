package security

import (
	"crypto/rand"
	"encoding/hex"
	"math"
	"math/big"
	"strconv"
	"time"

	"golang.org/x/crypto/argon2"
)

// https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#argon2id
const (
	argon2Iterations  = uint32(2)
	argon2Memory      = uint32(15 * 1024) // ~15 MB
	argon2Parallelism = uint8(1)
	argon2KeyLength   = uint32(16)
)

// PseudoRandomNonceArgon2 generates a pseudorandom hash for use as a nonce.
func PseudoRandomNonceArgon2() string {
	return hashArgon2(time.Now().Format("2006-01-02 15:04:05.000000000"), pseudoRandomSalt())
}

func hashArgon2(plainText string, salt string) string {
	return hex.EncodeToString(argon2.IDKey(
		[]byte(plainText),
		[]byte(salt),
		argon2Iterations,
		argon2Memory,
		argon2Parallelism,
		argon2KeyLength,
	))
}

func pseudoRandomSalt() string {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	return strconv.FormatUint(nBig.Uint64(), 10)
}
