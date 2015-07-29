package proxy

import (
    "net/http"
)

// CopyHeaders copy headers from source to destination.
// Nothing would be returned.
func CopyHeaders(dest, src http.Header) {
    for key, values := range src {
        for _, value := range values {
            dest.Add(key, value)
        }
    }
}

// ClearHeaders clear headers.
func ClearHeaders(headers http.Header) {
    for key, _ := range headers {
        headers.Del(key)
    }
}

// RmProxyHeaders remove Hop-by-hop headers.
func RmProxyHeaders(req *http.Request) {
    req.RequestURI = ""
    req.Header.Del("Proxy-Connection")
    req.Header.Del("Connection")
    req.Header.Del("Keep-Alive")
    req.Header.Del("Proxy-Authenticate")
    req.Header.Del("Proxy-Authorization")
    req.Header.Del("TE")
    req.Header.Del("Trailers")
    req.Header.Del("Transfer-Encoding")
    req.Header.Del("Upgrade")
}
