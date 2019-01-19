// See https://jbrandhorst.com/post/grpc-auth/
// 
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
	"os"
	"math/big"
	"time"

	////"github.com/patterns/fhirbuffer/insecure"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func enableTLSConfig() []grpc.DialOption {
	if !*tlsflag {
		// Vanilla TCP
		return []grpc.DialOption {grpc.WithInsecure()}
	}
        if *caFile != "" {
		// Using development env CRT, KEY files 

                creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
                if err != nil {
                        log.Fatalf("Failed to create TLS credentials %v", err)
                }
                return []grpc.DialOption {grpc.WithTransportCredentials(creds)}
	}


	// Heroku hosting
	return test1();
}

func test1() []grpc.DialOption {

	common := os.Getenv("FHIRBUFFER_CLIENT")

	tlsCert, err := genNewCert(common)
	if err != nil {
		log.Fatalln("Failed to create cert:", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{*tlsCert},
		RootCAs: CertPool,
		ServerName: *serverHostOverride,
	}

	creds := credentials.NewTLS(tlsConfig)

	return []grpc.DialOption { grpc.WithTransportCredentials(creds) }
}

func genNewCert(username string) (*tls.Certificate, error) {
	prvkey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	template, err := certTemplate(username, time.Hour *2)
	if err != nil {
		return nil, err
	}
	template.KeyUsage = x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature
	template.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}

	cert, err := x509.CreateCertificate(rand.Reader, template, Cert.Leaf, prvkey.Public(), Cert.PrivateKey)
	if err != nil {
		return nil, err
	}

	return &tls.Certificate{
		Certificate: [][]byte{cert},
		PrivateKey:  prvkey,
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
