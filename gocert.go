package main

import (
    "crypto/tls"
    "net"
    "log"
    "time"
    "fmt"
    "crypto/x509"
    "crypto/x509/pkix"
    "math"
    "os"
    "flag"
    "encoding/pem"
    "io/ioutil"

)

func Bold(str string) string {
      return "\033[1m" + str + "\033[0m"
}

func BoldPrintf(str string) {
    fmt.Printf(Bold(str))
}

func printPkix(pkixName pkix.Name) {
    fmt.Printf("%s - %s\n", pkixName.CommonName, pkixName.Organization)
}

func print_cert(cert *x509.Certificate) {
    BoldPrintf(Bold("CommonName: "))
    printPkix(cert.Subject)
    BoldPrintf("Issuer: ")
    printPkix(cert.Issuer)
    BoldPrintf("Alt Names: ")
    fmt.Printf("%s\n", cert.DNSNames)
    timeDifference := cert.NotAfter.Sub(time.Now())
    BoldPrintf("Expires in: ")
    fmt.Printf("%v days\n", math.Floor(timeDifference.Hours()/24))
    fmt.Printf("Version: %v ; Serial: %s\n", cert.Version, cert.SerialNumber)
    fmt.Printf("BasicConstraintsValid: %v ; IsCA: %v ; MaxPathLen: %v\n",
                cert.BasicConstraintsValid, cert.IsCA, cert.MaxPathLen)
}

func save_cert(cert *x509.Certificate) {
    block := pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw}
    pemdata := string(pem.EncodeToMemory(&block))

    err := ioutil.WriteFile(cert.Subject.CommonName + ".pem", []byte(pemdata), 0644)
    if err != nil { panic(err) }
}

func usage() {
    fmt.Fprintf(os.Stderr, "usage: %s [-save] [-chain] hostname:port\n", os.Args[0])
    flag.PrintDefaults()
    os.Exit(2)
}

func main() {

    var chain = flag.Bool("chain", false, "Print details about all the certificates on the chain")
    var save = flag.Bool("save", false, "Save the PEM to the local disk in the format of #{CN}.pem")
    flag.Parse()

    args := flag.Args()
    if len(args) < 1 {
        usage()
    }

    config := tls.Config{InsecureSkipVerify: true}
    conn, err := net.DialTimeout("tcp", args[0], 2*time.Second)

    if err != nil {
        log.Fatal(err)
    }

    client := tls.Client(conn, &config)
    err = client.Handshake()

    if err != nil {
        log.Fatal(err)
    }

    state := client.ConnectionState()

    certificates := []*x509.Certificate{}
    if *chain {
        certificates = state.PeerCertificates
    } else {
        certificates = []*x509.Certificate{state.PeerCertificates[0]}
    }

    for _, cert := range certificates {
        print_cert(cert)
        if *save {
            save_cert(cert)
        }
        fmt.Println("--------------------")
    }
    client.Close()
}