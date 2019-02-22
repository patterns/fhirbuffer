// See https://ericchiang.github.io/post/go-tls/
// See https://diogomonica.com/2017/01/11/hitless-tls-certificate-rotation-in-go/
//
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"log"
	"math/big"
	"os"
	"time"

	////"github.com/patterns/fhirbuffer/insecure"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func enableTLSConfig() []grpc.ServerOption {
	if !*tlsflag {
		// Vanilla TCP
		return []grpc.ServerOption{}
	}
	if *certFile != "" {
		// Using development env CRT, KEY files
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		return []grpc.ServerOption{grpc.Creds(creds)}
	}

	// Heroku hosting
	return test1()
}

func test1() []grpc.ServerOption {

	common := os.Getenv("FHIRBUFFER_COMMON")

	tlsCert, err := genNewCert(common)
	if err != nil {
		log.Fatalln("Failed to create cert:", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{*tlsCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    CertPool,
	}

	creds := credentials.NewTLS(tlsConfig)

	return []grpc.ServerOption{grpc.Creds(creds)}
}

func genNewCert(username string) (*tls.Certificate, error) {
	svckey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	svctemplate, err := certTemplate(username, time.Hour*12)
	if err != nil {
		return nil, err
	}
	svctemplate.KeyUsage = x509.KeyUsageDigitalSignature
	svctemplate.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth}

	cert, err := x509.CreateCertificate(rand.Reader, svctemplate, Cert.Leaf, svckey.Public(), Cert.PrivateKey)
	if err != nil {
		return nil, err
	}

	return &tls.Certificate{
		Certificate: [][]byte{cert},
		PrivateKey:  svckey,
	}, nil
}

func certTemplate(common string, expire time.Duration) (*x509.Certificate, error) {
	randOrg := make([]byte, 32)
	_, err := rand.Read(randOrg)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, err
	}

	return &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{base64.URLEncoding.EncodeToString(randOrg)},
			CommonName:   common,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(expire),
		BasicConstraintsValid: true,
	}, nil
}
