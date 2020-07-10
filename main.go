package main

import (
	"bytes"
	"github.com/google/brotli/go/cbrotli"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest(http.MethodGet, "http://localhost", nil)
	if err != nil {
		log.Fatal("error creating req", err)
	}
	req.Header.Set("Accept-Encoding", "br")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("error fetching url")
	}
	// illustrational purpose - not necessary
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal("error reading response")
	}
	log.Println("encoded resp \n", string(body))
	// put response as a reader back to body
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	defer resp.Body.Close()
	if resp.Header.Get("Content-Encoding") == "br" {
		// do br decoding
		reader := cbrotli.NewReader(resp.Body)
		defer reader.Close()
		respBody, err := ioutil.ReadAll(reader)
		if err != nil {
			log.Fatal("error decoding br response", err)
		}
		log.Println("decoded resp \n", string(respBody))
	} else {
		log.Fatal("response not br encoded")
	}
}
