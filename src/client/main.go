package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {

	var (
		caCertFile, resolve string
		dumpHeaders         bool
	)

	flag.BoolVar(&dumpHeaders, "i", false, "print headers")
	flag.StringVar(&caCertFile, "cacert", "", "ca certificate")
	flag.StringVar(&resolve, "resolve", "", "dns resolver")
	flag.Parse()

	target := flag.Arg(0)

	caCert, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		log.Fatal(err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	if resolve != "" {

		rsv := strings.SplitN(resolve, ":", 3)
		if len(rsv) != 3 {
			log.Fatalf(" -resolve format is [domain]:[port]:[ip]\n")
		}

		m := map[string]string{
			rsv[0] + ":" + rsv[1]: rsv[2] + ":" + rsv[1],
		}

		client.Transport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {

			ipAddr := m[addr]

			if ipAddr == "" {
				return nil, fmt.Errorf("target %s not match resolve %s", addr, resolve)
			}

			conn, err := net.Dial("tcp", ipAddr)
			return conn, err
		}
	}

	res, err := client.Get(target)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	var headers []byte

	if dumpHeaders {
		headers, _ = httputil.DumpResponse(res, false)
		fmt.Println(string(headers))
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(b))
}
