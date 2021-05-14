package main

import (
	"encoding/json"
    "net/http"
)
type Data struct {
	Name string
	Location string
	Age int
	Infectedtype string
	State string
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var d Data
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {}else{}
	})
    http.ListenAndServe(":8000", nil)
	
}