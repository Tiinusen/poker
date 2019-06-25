package cmd

import (
	"crypto/tls"
	"crypto/x509"
)

var trustedCerts *x509.CertPool
var serverCert *tls.Certificate

//go:generate go-bindata -o assets.go -pkg $GOPACKAGE -prefix ../../../assets/certificates ../../../assets/certificates/...

func getTrustedCerts() *x509.CertPool {
	if trustedCerts != nil {
		return trustedCerts
	}
	trustedCerts, err := x509.SystemCertPool()
	if err != nil {
		trustedCerts = x509.NewCertPool()
	}
	rootCertBytes, err := rootPokerCrtBytes()
	if err != nil {
		panic(err)
	}
	trustedCerts.AppendCertsFromPEM(rootCertBytes)
	return trustedCerts
}

func getBuiltinCert() *tls.Certificate {
	if serverCert != nil {
		return serverCert
	}
	serverCertData, err := pokerCrtBytes()
	if err != nil {
		panic(err)
	}
	serverCertKey, err := pokerKeyBytes()
	if err != nil {
		panic(err)
	}
	cert, err := tls.X509KeyPair(serverCertData, serverCertKey)
	if err != nil {
		panic(err)
	}
	serverCert = &cert
	return serverCert
}
