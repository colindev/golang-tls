package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	var addr, crtFile, keyFile string

	flag.StringVar(&addr, "addr", ":443", "https listen on")
	flag.StringVar(&crtFile, "crt", "", "Certficate")
	flag.StringVar(&keyFile, "key", "", "Private key")
	flag.Parse()

	log.Fatal(http.ListenAndServeTLS(addr, crtFile, keyFile, http.HandlerFunc(handler)))

}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
