package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// START OMIT
type Address struct {
	Street   string
	Postcode string
}

func changeAddress(rw http.ResponseWriter, req *http.Request) {
	limited := io.LimitReader(req.Body, 4*1024) // HL
	decoder := json.NewDecoder(limited)
	var addr Address
	err := decoder.Decode(&addr)
	if err != nil {
		panic(err)
	}
	log.Printf("Address set to: %#v\n", addr)
}

func main() {
	http.HandleFunc("/change-address", changeAddress)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// END OMIT
