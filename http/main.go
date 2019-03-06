package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// create client support https without cert
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", "https://localhost:8443/dashboard", nil)
	req.Header.Set("User-Agent", "my-header)")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	response, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", response)
}
