package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"scabarrus.com/k8s.webhook/internal/service"
)

func helloWorld(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,r.URL.Path)
	fmt.Fprintf(w,r.Method)
}
func main(){
	certDir:=os.Getenv("CRT_DIR")
	r := mux.NewRouter()
	authn:=service.Authn{}
	fmt.Println("authentication webhook")
	r.HandleFunc("/",authn.CheckAuthentication).Methods("POST")
	//http.ListenAndServe(":9444",r)
    log.Fatal(http.ListenAndServeTLS(":9444",certDir+"certificate.crt",certDir+"privateKey.key",r))
	

}
