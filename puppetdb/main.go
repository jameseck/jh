package puppetdb

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Puppetdb struct {
	sslkeypair tls.Certificate
	sslca      string
	host       string
}

type Fact struct {
	Value    string `json:"value"`
	Name     string `json:"name"`
	Certname string `json:"certname"`
}

type Facts []struct {
	Fact
}

func New(certfile string, keyfile string, cafile string, host string) *Puppetdb {
	sslkeypair, err := tls.LoadX509KeyPair(certfile, keyfile)
	if err != nil {
		log.Fatalf("ERROR client certificate: %s", err)
	}
	return &Puppetdb{
		sslkeypair: sslkeypair,
		sslca:      cafile,
		host:       host,
	}
}

func (puppetdb *Puppetdb) Get(query string) (response Facts) {
	fmt.Println("vim-go")

	query = `[ "and",
                  [ "or",
                    [ "~", "certname", ".*" ]
                  ],
                  [ "or",
                    [ "=", "name", "ipaddress" ]
                  ]
                ]`

	// Load CA cert
	caCert, err := ioutil.ReadFile(puppetdb.sslca)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{puppetdb.sslkeypair},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	// Do GET something
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/v3/facts", puppetdb.host), nil)

	q := req.URL.Query()
	q.Add("query", query)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var facts Facts
	err = decoder.Decode(&facts)

	for _, f := range facts {
		fmt.Printf("Certname: %s Name: %s Value: %s\n", f.Certname, f.Name, f.Value)
	}

	return facts
}
