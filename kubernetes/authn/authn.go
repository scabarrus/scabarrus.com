package authn

import (
	"fmt"
	"net/http"
	"strings"
	"encoding/json"
	"strconv"
	"scabarrus.com/kubernetes/authn/repository"
)
type TokenReview struct {
	Version string `json:"apiVerson"`
	Kind string `json:"kind"`
	Status struct {
			Authentication bool `json:"authenticated"`
			UserInfo struct {
					Username string `json:"username"`
					Uid string `json:"uid"`
					Groups []string `json:"groups"`
			}`json:"user"`
	} `json:"status"`
}
type AuthRequest struct {
	Version string `json:"apiVerson"`
	Kind string `json:"kind"`
	Spec struct {
		Token string `json:"token"`
	} `json:"spec"`

}
type Authn struct {
	Token string
	TokenReview TokenReview

}
func (a Authn) Authentication (login string,password string) TokenReview{
	var db repository.User
	//var user internal.User=internal.User{}
	_,user:=db.FindByLogin(login)
	a.TokenReview.Version="authentication.k8s.io/v1beta1"
	a.TokenReview.Kind="TokenReview"
	if strings.TrimSpace(user.Password) == strings.TrimSpace(password) {
			a.Token=user.Password
			a.TokenReview.Status.Authentication=true
			a.TokenReview.Status.UserInfo.Username=user.User
			a.TokenReview.Status.UserInfo.Uid=strconv.Itoa(user.Uid)
		
			for i,_ := range user.FK_groups {
				
				a.TokenReview.Status.UserInfo.Groups=append(a.TokenReview.Status.UserInfo.Groups,user.FK_groups[i].Group)
			}
			return a.TokenReview
	}else{
			a.TokenReview.Status.Authentication=false
			return a.TokenReview
	}

}

func (a Authn)ReturnTokenReview (w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	request:=AuthRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
    	if err != nil {
        	http.Error(w, err.Error(), http.StatusBadRequest)
        	return
    	}
	fmt.Println("PAYLOAD request: ",request)
	fmt.Println("PAYLOAD token: ",request.Spec.Token)
	token:=request.Spec.Token
	fmt.Println("Token ",token)
	credential:=strings.Split(token,":")
	fmt.Println("Credential in header ",strings.TrimSpace(credential[0])," ", strings.TrimSpace(credential[1]))
	json.NewEncoder(w).Encode(a.Authentication(strings.TrimSpace(credential[0]),credential[1]))
}
