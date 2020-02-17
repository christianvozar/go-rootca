// Copyright © 2020 Christian R. Vozar <christian@rogueethic.com>
//
// Package rootca provides a http.Client with a certificate bundle.
package gorootca

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
)

// New returns a http.Client with a pre-configured root certificate bundle.
func New() *http.Client {
	cp := x509.NewCertPool()
	cp.AppendCertsFromPEM(rootCertificateBundle)

	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: cp}}}
}
