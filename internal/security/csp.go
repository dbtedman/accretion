package security

import "fmt"

func GenerateContentSecurityPolicy(nonce string) string {
	defaultSrc := "default-src 'self'"
	styleSrc := fmt.Sprintf("style-src 'nonce-%s'", nonce)
	fontSrc := "font-src https://fonts.gstatic.com"

	// https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP
	return defaultSrc + "; " + styleSrc + "; " + fontSrc + ";"
}
