package helps

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// WithResponseHeaderTimeout returns a shallow-cloned client whose standard
// transport times out only while waiting for upstream response headers.
func WithResponseHeaderTimeout(client *http.Client, timeout time.Duration) *http.Client {
	if timeout <= 0 {
		return client
	}
	if client == nil {
		client = &http.Client{}
	}

	clone := *client
	switch transport := clone.Transport.(type) {
	case nil:
		if defaultTransport, ok := http.DefaultTransport.(*http.Transport); ok {
			transportClone := defaultTransport.Clone()
			transportClone.ResponseHeaderTimeout = timeout
			clone.Transport = transportClone
		}
	case *http.Transport:
		transportClone := transport.Clone()
		transportClone.ResponseHeaderTimeout = timeout
		clone.Transport = transportClone
	default:
		log.Debugf("response header timeout not applied for transport type %T", transport)
	}

	return &clone
}
