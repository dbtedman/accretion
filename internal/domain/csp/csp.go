package csp

import (
	"encoding/hex"
	"golang.org/x/crypto/argon2"
	"math/rand"
	"strconv"
	"time"
)

const argon2Iterations = uint32(2)
const argon2Memory = uint32(15 * 1024) // ~15 MB
const argon2Parallelism = uint8(1)
const argon2KeyLength = uint32(16)

func Generate(nonce string) string {
	defaultSrc := "default-src 'self'"
	styleSrc := "style-src 'self' 'nonce-" + nonce + "' https://fonts.googleapis.com"
	fontSrc := "font-src 'self' https://fonts.gstatic.com"

	// https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP
	return defaultSrc + "; " + styleSrc + "; " + fontSrc + ";"
}

func GenerateNonce() string {
	plainText := time.Now().Format("2006-01-02 15:04:05.000000000")
	salt := strconv.FormatUint(rand.Uint64(), 10)

	hash := argon2.IDKey(
		[]byte(plainText),
		[]byte(salt),
		argon2Iterations,
		argon2Memory,
		argon2Parallelism,
		argon2KeyLength,
	)

	return hex.EncodeToString(hash)
}
