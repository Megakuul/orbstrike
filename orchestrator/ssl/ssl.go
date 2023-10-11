package ssl

import (
	"fmt"
	
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"

	"google.golang.org/grpc/credentials"
)

func GetAuthCredentials(cert string, key string, ca string) (credentials.TransportCredentials, error) {
	if cert==""||key==""||ca=="" {
		return nil, nil
	}
	
	decCert, _ := base64.StdEncoding.DecodeString(cert)
	decKey, _ := base64.StdEncoding.DecodeString(key)
	decCa, _ := base64.StdEncoding.DecodeString(ca)
	
	certPair, err := tls.X509KeyPair(decCert, decKey)
	if err!=nil {
		return nil, err
	}

	caPool := x509.NewCertPool()
	if !caPool.AppendCertsFromPEM(decCa) {
		return nil, fmt.Errorf("Failed to add CA to certificate pool!")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certPair},
		ClientAuth: tls.NoClientCert,
		ClientCAs: caPool,
	})
	
	return creds, nil
}

func GetTLSClientMutual(cert string, key string, ca string) (*tls.Config, error) {
	if cert==""||key==""||ca=="" {
		return nil, nil
	}

	decCert, _ := base64.StdEncoding.DecodeString(cert)
	decKey, _ := base64.StdEncoding.DecodeString(key)
	decCa, _ := base64.StdEncoding.DecodeString(ca)

	certPair, err := tls.X509KeyPair(decCert, decKey)
	if err!=nil {
		return nil, err
	}

	caPool := x509.NewCertPool()
	if !caPool.AppendCertsFromPEM(decCa) {
		return nil, fmt.Errorf("Failed to add CA to certificate pool!")
	}

	return &tls.Config{
		MinVersion: tls.VersionTLS12,
		Certificates: []tls.Certificate{certPair},
		RootCAs: caPool,
	}, nil
}

func GetTLSClientCA(ca string) (*tls.Config, error) {
	if ca=="" {
		return nil, nil
	}

	decCa, _ := base64.StdEncoding.DecodeString(ca)

	caPool := x509.NewCertPool()
	if !caPool.AppendCertsFromPEM(decCa) {
		return nil, fmt.Errorf("Failed to add CA to certificate pool!")
	}

	return &tls.Config{
		RootCAs: caPool,
	}, nil
}
