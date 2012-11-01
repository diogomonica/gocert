package main

import (
	"crypto/tls"
	"net"
	"log"
	"time"
	"fmt"
	"crypto/x509"
	//"encoding/pem"
)

func handle_cert(cert *x509.Certificate) {
	// block := pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw}
	//pemdata := string(pem.EncodeToMemory(&block))
	fmt.Println(cert.Subject.CommonName)
	fmt.Println(cert.Version)
	fmt.Println(cert.SerialNumber)
	fmt.Println(cert.Subject)
	fmt.Println(cert.NotBefore, cert.NotAfter)

	//fmt.Println(pemdata)
}

func main() {

	config := tls.Config{InsecureSkipVerify: true}

	conn, err := net.DialTimeout("tcp", "squareup.com:443", 2*time.Second)

	if err != nil {
		log.Fatal(err)
	}

	client := tls.Client(conn, &config)

	err = client.Handshake()

	if err != nil {
		log.Fatal(err)
	}

	state := client.ConnectionState()

	for _, cert := range state.PeerCertificates {
		handle_cert(cert)
	}
	client.Close()
}