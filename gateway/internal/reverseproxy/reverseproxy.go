package reverseproxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func New(cluster string) (http.HandlerFunc, error) {
	remote, err := url.Parse(cluster)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)

	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Host = remote.Host
		r.URL.Scheme = remote.Scheme
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
		r.Host = remote.Host

		proxy.ServeHTTP(w, r)
	}, nil
}
