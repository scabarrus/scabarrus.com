package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	certDir="./"
)
func validate(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"validate")
	var request map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		fmt.Println("error : ",err)
	}
	fmt.Println("Payload : ",request)
}

func mutating(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"validate")
	var request map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil{
		fmt.Println("error : ",err)
	}
	fmt.Println("Payload : ",request)

}

func main(){
	r := mux.NewRouter()
	// Replace http.HandleFunc by router.HandleFunc.
	r.HandleFunc("/validate", validate)
	r.HandleFunc("/mutating", mutating)
	// Replace 2nd parameter by the configured mux router.
    log.Fatal(http.ListenAndServeTLS(":9443",certDir+"certificate.crt",certDir+"privateKey.key",nil))
}