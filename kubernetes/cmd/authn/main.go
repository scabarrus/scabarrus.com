package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"scabarrus.com/kubernetes/authn"
	"scabarrus.com/kubernetes/authn/repository"
)
	
const(
    //certDir= "/go/src/scabarrus.com/kubernetes/configs/"
    certDir= "scabarrus.com/kubernetes/configs/"
)
func main() {
    fmt.Println("Authorization webhook")
    var a authn.Authn
    a=authn.Authn{}
    os.Setenv("POSTGRES_CON", "user=admin password=password dbname=admin host=192.168.169.129 port=31018 sslmode=disable")
    
    pg := repository.Postgres{}
    pg.Initialization()
    
    var groupsModel repository.Groups
    groupsModel.DBMigrate(pg.Database)
    
    var userModel repository.User
    userModel.DBMigrate(pg.Database)

    var groups []repository.Groups
    groups=append(groups,repository.Groups{Group:"admin",Description: "description",Namespace: "namespace-adm"})
    groups=append(groups,repository.Groups{Group:"dev",Description: "description",Namespace: "namespace-dev"})
    pg.InitDB(groups)
    
    pg.InitDB(repository.User{User:"marvel",Password:"password",Uid: 6002,FK_groups:groups})
    pg.InitDB(repository.User{User:"spiderman",Password:"password",Uid: 6003,FK_groups:groups})
    http.HandleFunc("/",a.ReturnTokenReview)
    //http.ListenAndServe(":9443",nil)
    log.Fatal(http.ListenAndServeTLS(":9443",certDir+"certificate.crt",certDir+"privateKey.key",nil))
   

}
