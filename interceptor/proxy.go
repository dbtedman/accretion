package interceptor

import (
	"github.com/charmbracelet/log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func NewProxyHandler(toUrl url.URL) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(&toUrl)
	proxy.ModifyResponse = func(response *http.Response) error {
		return nil
	}

	return proxy
}

func ListenHTTPWithProxy(proxyServer Proxy, proxyAddressURL url.URL, listenAddress string, errorCh *chan error) {
	aProxy := NewProxyHandler(proxyAddressURL)
	aProxy.Transport = TransportLogger{}

	handler := func(writer http.ResponseWriter, request *http.Request) {
		request.Host = request.URL.Host
		aProxy.ServeHTTP(writer, request)
	}

	log.Infof("Listening to requests at %s as a proxy to %s.", listenAddress, proxyAddressURL.String())
	if err := proxyServer.ListenAndServe(listenAddress, handler); err != nil {
		*errorCh <- err
	} else {
		*errorCh <- nil
	}
}

type TransportLogger struct{}

func (TransportLogger) RoundTrip(req *http.Request) (*http.Response, error) {
	res, err := http.DefaultTransport.RoundTrip(req)

	log.Infof("%s %s [%s %s] %s %d", time.Now().Format(time.RFC3339), req.RemoteAddr, req.Method, req.RequestURI, res.Status, res.ContentLength)

	return res, err
}

type Proxy interface {
	ListenAndServe(addr string, handler http.HandlerFunc) error
}

type ServerProxy struct {
}

var _ Proxy = &ServerProxy{}

func (s *ServerProxy) ListenAndServe(addr string, handler http.HandlerFunc) error {
	server := &http.Server{
		Addr:         addr,
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
		IdleTimeout:  time.Minute,
		Handler:      handler,
	}

	return server.ListenAndServe()
}

type TestProxy struct {
	Addr    string
	Handler http.HandlerFunc
}

var _ Proxy = &TestProxy{}

func (t *TestProxy) ListenAndServe(addr string, handler http.HandlerFunc) error {
	t.Addr = addr
	t.Handler = handler
	return nil
}
