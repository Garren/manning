package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

func main() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}
	// set the configuration for our cert
	template := x509.Certificate{
		SerialNumber: serialNumber, // generally set up by a CA
		Subject:      subject,
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365*24 + time.Hour),
		// an ssl cert is an x509 certificate with the key extended usage set
		// to server authentication
		KeyUsage:    x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")},
	}
	// generate a private key.
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	// a der is a variation of the ITU-T "Basic Encoding Rules" standard. It
	// specifies a self describing/delimiting format for encoding ASN.1 (abstract
	// syntax notation) structures
	derBytes, _ := x509.CreateCertificate(
		rand.Reader,
		&template,
		&template,
		&pk.PublicKey,
		pk)

	// a PEM is a base 64 encoded DER
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}
