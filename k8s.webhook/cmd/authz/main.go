package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"scabarrus.com/k8s.webhook/internal/service"
)

func main(){
	certDir:=os.Getenv("CRT_DIR")
	r := mux.NewRouter()
	authz:=service.Authz{}
	fmt.Println("authorization webhook")
	r.HandleFunc("/",authz.CheckAccess).Methods("POST")
	//http.ListenAndServe(":9444",r)
    log.Fatal(http.ListenAndServeTLS(":9444",certDir+"certificate.crt",certDir+"privateKey.key",r))
}