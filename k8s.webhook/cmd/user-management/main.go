package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"scabarrus.com/k8s.webhook/internal/service"
)

func helloWorld(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,r.URL.Path)
	fmt.Fprintf(w,r.Method)
}
// @title User-Managmeent API
// @version 1.0
// @description This is a sample serice for managing user and role for kubernetes cluster
// @termsOfService http://swagger.io/terms/
// @contact.name scabarrus
// @contact.email scabarrus@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9443
// @BasePath /api/v1/
func main (){
	os.Setenv("POSTGRES_CON", "user=admin password=password dbname=admin host=192.168.169.129 port=31018 sslmode=disable")
	r := mux.NewRouter()
	
	dbm:=service.DBMigrate{}
	r.HandleFunc("/api/admin/db-migrate",dbm.MigrateService).Methods("GET")
	// Manage Users
	userService:=service.UserService{}
	r.HandleFunc("/api/v1/users", userService.FindAll).Methods("GET")
	r.HandleFunc("/api/v1/users", userService.Save).Methods("POST")
	r.HandleFunc("/api/v1/users/{user}", userService.FindByName).Methods("GET")
	r.HandleFunc("/api/v1/users/{user}", userService.Modify).Methods("PUT")
	r.HandleFunc("/api/v1/users/{user}", userService.Delete).Methods("DELETE")
	
	// Manage Groups
	groupService:=service.GroupService{}
	r.HandleFunc("/api/v1/groups", groupService.FindAll).Methods("GET")
	r.HandleFunc("/api/v1/groups", groupService.Save).Methods("POST")
	r.HandleFunc("/api/v1/groups/{group}", groupService.FindByName).Methods("GET")
	r.HandleFunc("/api/v1/groups/{group}", groupService.Modify).Methods("PUT")
	r.HandleFunc("/api/v1/groups/{group}", groupService.Delete).Methods("DELETE")

	//Manage user's group
	memberService:=service.MemberService{}
	/*r.HandleFunc("/api/v1/groups/{group}/users", helloWorld).Methods("GET")*/
	r.HandleFunc("/api/v1/groups/{group}/members", memberService.Save).Methods("POST")
	r.HandleFunc("/api/v1/groups/{group}/members/{member}", memberService.FindByName).Methods("GET")
	r.HandleFunc("/api/v1/groups/{group}/members/{member}", memberService.Delete).Methods("DELETE")

	/*
	// Manage Roles
	r.HandleFunc("/api/v1/roles", helloWorld).Methods("GET")
	r.HandleFunc("/api/v1/roles", helloWorld).Methods("POST")
	r.HandleFunc("/api/v1/roles/{role}", helloWorld).Methods("GET")
	r.HandleFunc("/api/v1/roles/{role}", helloWorld).Methods("PUT")
	r.HandleFunc("/api/v1/roles/{role}", helloWorld).Methods("DELETE")

	// Manage policies
	r.HandleFunc("/api/v1/roles/{role}/policies", helloWorld).Methods("GET")
	r.HandleFunc("/api/v1/roles/{role}/policies/{policy}", helloWorld).Methods("POST")
	r.HandleFunc("/api/v1/roles/{role}/policies/{policy}", helloWorld).Methods("GET")
	r.HandleFunc("/api/v1/roles/{role}/policies/{policy}", helloWorld).Methods("PUT")
	r.HandleFunc("/api/v1/roles/{role}/policies/{policy}", helloWorld).Methods("DELETE")

	// Manage Namespace
	r.HandleFunc("/api/v1/namespaces", helloWorld).Methods("GET")
	r.HandleFunc("/api/v1/namespaces", helloWorld).Methods("POST")
	r.HandleFunc("/api/v1/namespaces/{namespace}", helloWorld).Methods("GET")
	r.HandleFunc("/api/v1/namespaces/{namespace}", helloWorld).Methods("PUT")
	r.HandleFunc("/api/v1/namespaces/{namespace}", helloWorld).Methods("DELETE")
	*/
    http.ListenAndServe(":9443", r)
}
/*
func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})
	*/