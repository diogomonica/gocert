gocert
======

A small go program to print details on X509 certificates and safe them in .pem format.

# Why?

- Because I can never remember openssl's syntax. The s_client -connect is not intuitive. 
- Because I want a better way for displaying certificates, with emphasis on what matters.
- Because I want to be able to extract remote certificates .pem's easily

# Installation 

- Just build gocert.go and move it somewhere in your path. For example, if you are running OSX (this you already have the GO environment installed):

	`cd /tmp; curl -O https://raw.github.com/diogomonica/gocert/master/gocert.go;
	go build gocert.go; mv gocert /usr/local/bin`

# Usage

	usage: gocert [-save] [-chain] hostname:port
	  -chain=false: Print details about all the certificates on the chain
	  -save=false: Save the PEM to the local disk in the format of #{CN}.pem

# Examples

### Retrieving information about the certificate of diogomonica.com:

	# gocert diogomonica.com:443	
	CommonName: www.diogomonica.com - []
	Issuer: StartCom Class 1 Primary Intermediate Server CA - [StartCom Ltd.]
	Alt Names: [www.diogomonica.com diogomonica.com]
	Expires in: 359 days
	Version: 3 ; Serial: 516289
	BasicConstraintsValid: true ; IsCA: false ; MaxPathLen: -1
	--------------------

### Retrieving information about the full chain of diogomonica.com:

	# gocert -chain diogomonica.com:443
	CommonName: www.diogomonica.com - []
	Issuer: StartCom Class 1 Primary Intermediate Server CA - [StartCom Ltd.]
	Alt Names: [www.diogomonica.com diogomonica.com]
	Expires in: 359 days
	Version: 3 ; Serial: 516289
	BasicConstraintsValid: true ; IsCA: false ; MaxPathLen: -1
	--------------------
	CommonName: StartCom Class 1 Primary Intermediate Server CA - [StartCom Ltd.]
	Issuer: StartCom Certification Authority - [StartCom Ltd.]
	Alt Names: []
	Expires in: 1816 days
	Version: 3 ; Serial: 24
	BasicConstraintsValid: true ; IsCA: true ; MaxPathLen: -1
	--------------------
	CommonName: StartCom Certification Authority - [StartCom Ltd.]
	Issuer: StartCom Certification Authority - [StartCom Ltd.]
	Alt Names: []
	Expires in: 8719 days
	Version: 3 ; Serial: 1
	BasicConstraintsValid: true ; IsCA: true ; MaxPathLen: -1
	--------------------

### Retrieving and saving in PEM format the certificate of diogomonica.com:

	# gocert -save diogomonica.com:443
	CommonName: www.diogomonica.com - []
	Issuer: StartCom Class 1 Primary Intermediate Server CA - [StartCom Ltd.]
	Alt Names: [www.diogomonica.com diogomonica.com]
	Expires in: 359 days
	Version: 3 ; Serial: 516289
	BasicConstraintsValid: true ; IsCA: false ; MaxPathLen: -1
	--------------------
	# * ls -alh www.diogomonica.com.pem 
	-rw-r--r--  1 diogo  diogo   2.8K Nov  2 22:19 www.diogomonica.com.pem
