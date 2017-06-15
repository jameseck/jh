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

type Conn struct {
	sslkeypair tls.Certificate
	sslca      string
	host       string
}

type FactFilters struct {
	Filters []FactFilter
}

type FactFilter struct {
	Name     string
	Value    string
	Operator string
}

type Fact struct {
	Value    string `json:"value"`
	Name     string `json:"name"`
	Certname string `json:"certname"`
}

type Facts []struct {
	Fact
}

func New(certfile string, keyfile string, cafile string, host string) *Conn {
	sslkeypair, err := tls.LoadX509KeyPair(certfile, keyfile)
	if err != nil {
		log.Fatalf("ERROR client certificate: %s", err)
	}
	return &Conn{
		sslkeypair: sslkeypair,
		sslca:      cafile,
		host:       host,
	}
}

func queryBuilder(HostRegex string, ff FactFilters, factAndOr string) (out string) {

	out = "[ \"and\"\n"
	out += ", [ \"or\"\n"
	out += fmt.Sprintf(",   [ \"~\", \"certname\", \"%s\" ]\n", HostRegex)
	out += "  ]\n"
	out += ", [ \"or\"\n"

	for i := range ff.Filters {
		out += fmt.Sprintf(",   [ \"=\", \"name\", \"%s\" ]\n", ff.Filters[i].Name)
	}
	out += "  ]\n"

	out += ", [ \"and\"\n"
	for i := range ff.Filters {
		out += ",   [ \"in\", \"certname\"\n"
		out += ",     [ \"extract\", \"certname\", [ \"select-facts\"\n"
		out += ",       [\"and\"\n"
		out += fmt.Sprintf(",         [ \"=\", \"name\", \"%s\" ]\n", ff.Filters[i].Name)
		out += fmt.Sprintf(",         [ \"%s\", \"value\", \"%s\" ]\n", ff.Filters[i].Operator, ff.Filters[i].Value)
		out += "        ]\n"
		out += "      ]]\n"
		out += "    ]\n"
	}
	//	out += "          ]\n"
	out += "  ]\n"
	out += "]\n"

	return out
}

func (puppetdb *Conn) Get(HostRegex string, ff FactFilters, factAndOr string) (response string) {
	fmt.Println("Running puppetdb.Get")

	query := queryBuilder(HostRegex, ff, factAndOr)
	fmt.Printf(query)

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

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Printf(bodyString)

	decoder := json.NewDecoder(resp.Body)
	var facts Facts
	err = decoder.Decode(&facts)

	var out string

	for _, f := range facts {
		out += fmt.Sprintf("Certname: %s Name: %s Value: %s\n", f.Certname, f.Name, f.Value)
	}

	return out
}
