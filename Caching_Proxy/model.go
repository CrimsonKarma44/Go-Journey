package Caching_Proxy

import (
	"github.com/gregjones/httpcache"
	"net/http"
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
