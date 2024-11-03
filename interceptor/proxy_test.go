package interceptor_test

import (
	"errors"
	"github.com/dbtedman/accretion/interceptor"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProxyHandler(t *testing.T) {
	toUrl, _ := url.Parse("https://example.com")

	proxy := interceptor.NewProxyHandler(*toUrl)

	assert.NotNil(t, proxy)
}

func TestNewProxyHandlerModifyResponse(t *testing.T) {
	toUrl, _ := url.Parse("https://example.com")
	proxy := interceptor.NewProxyHandler(*toUrl)
	response := &http.Response{Header: http.Header{}}

	err := proxy.ModifyResponse(response)

	assert.Nil(t, err)
}

func TestListenHTTPWithProxy(t *testing.T) {
	errorCh := make(chan error)
	proxy := interceptor.TestProxy{}
	proxyAddress, _ := url.Parse("https://example.com")

	go func() {
		interceptor.ListenHTTPWithProxy(&proxy, *proxyAddress, ":3001", &errorCh)
	}()
	err := <-errorCh

	assert.Nil(t, err)
	assert.Equal(t, ":3001", proxy.Addr)
}

func TestListenHTTPWithProxyHandlesListenAndServeError(t *testing.T) {
	errorCh := make(chan error)
	proxy := alwaysErrorProxy{}
	proxyAddress, _ := url.Parse("https://example.com")

	go func() {
		interceptor.ListenHTTPWithProxy(&proxy, *proxyAddress, ":3001", &errorCh)
	}()
	err := <-errorCh

	assert.ErrorContains(t, err, alwaysErrorMessage)
}

const alwaysErrorMessage = "always showing error"

type alwaysErrorProxy struct {
}

var _ interceptor.Proxy = &alwaysErrorProxy{}

func (t *alwaysErrorProxy) ListenAndServe(addr string, handler http.HandlerFunc) error {
	return errors.New(alwaysErrorMessage)
}
