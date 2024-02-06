package main

import (
	"crypto/rand"
	"embed"
	"encoding/hex"
	"fmt"
	"html/template"
	"math"
	"math/big"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"golang.org/x/crypto/argon2"
)

//go:embed static/*
var staticFS embed.FS

//go:embed html/*
var htmlFS embed.FS

const (
	HTTPContentType           = "Content-Type"
	HTTPCacheControl          = "Cache-Control"
	HTTPContentSecurityPolicy = "Content-Security-Policy"
)

// https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#argon2id
const (
	argon2Iterations  = uint32(2)
	argon2Memory      = uint32(15 * 1024) // ~15 MB
	argon2Parallelism = uint8(1)
	argon2KeyLength   = uint32(16)
)

func main() {
	htmlTemplates := template.Must(template.ParseFS(htmlFS, "html/*.gohtml"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		isNotUIRoute, _ := regexp.MatchString("^(static|api)/*", r.URL.Path)

		if isNotUIRoute {
			return
		}

		nonce := pseudoRandomNonceArgon2()

		w.Header().Set(HTTPContentType, "text/html; charset=utf-8")
		w.Header().Set(HTTPCacheControl, "max-age=0, private, must-revalidate")
		w.Header().Set(HTTPContentSecurityPolicy, fmt.Sprintf("default-src 'none'; style-src 'nonce-%s'; font-src https://fonts.gstatic.com", nonce))

		type DiscoverData struct {
			Description string
			Title       string
			Nonce       string
			Message     string
		}

		err := htmlTemplates.ExecuteTemplate(w, "discover.gohtml", DiscoverData{
			Description: "",
			Title:       "",
			Nonce:       nonce,
			Message:     "Hello, World!",
		})

		if err != nil {
			fmt.Println(err)
		}
	})

	http.Handle("/static/", staticHandlerMiddleware(http.FileServer(http.FS(staticFS))))

	server := &http.Server{
		Addr:         ":3000",
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
		IdleTimeout:  time.Minute,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}

// PseudoRandomNonceArgon2 generates a pseudorandom hash for use as a nonce.
func pseudoRandomNonceArgon2() string {
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

func staticHandlerMiddleware(assetsFS http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		assetsFS.ServeHTTP(w, r)
	}
}
