package Caching_Proxy

import (
	"net/http"

	"github.com/gregjones/httpcache"
)

type CacheAwareTransport struct {
	Transport http.RoundTripper
}

func (t *CacheAwareTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := t.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// Check if the response was served from cache
	if resp.Header.Get(httpcache.XFromCache) != "" {
		resp.Header.Set("X-Cache", "HIT")
	} else {
		resp.Header.Set("X-Cache", "MISS")
	}

	return resp, nil
}

type HeaderInjectorTransport struct {
	Transport http.RoundTripper
}

func (t *HeaderInjectorTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	transport := t.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	resp, err := transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// If no cache control header is present, add one to allow caching
	// This helps when the origin server doesn't send cache headers
	if resp.Header.Get("Cache-Control") == "" {
		resp.Header.Set("Cache-Control", "public, max-age=3600")
	}
	return resp, nil
}
