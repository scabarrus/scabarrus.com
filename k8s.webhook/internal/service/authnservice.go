package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	auth "k8s.io/api/authentication/v1beta1"
	"scabarrus.com/k8s.webhook/internal/dto"
	lib "scabarrus.com/k8s.webhook/internal/utils"
)

type Authn struct{
	Token auth.TokenReview
}
//CheckAuthentication is a handler method that return a token review with authentication status
func (a Authn)CheckAuthentication(w http.ResponseWriter, r *http.Request){
	
	w.Header().Set("Content-Type", "application/json")
	a.Token.Status.Error=""
	_=json.NewDecoder(r.Body).Decode(&a.Token)
	fmt.Println("payload : ",a.Token)
	token:=strings.Split(a.Token.Spec.Token,":")
	if len(token[0]) <1 || len(token[1])<1{
		a.Token.Status.Authenticated=false
		a.Token.Status.Error="token is not correct!"
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(&a.Token)
	}else{
		fmt.Println("retrieve user |",strings.TrimSpace(token[0]),"|")
		url:=os.Getenv("USR_API")+strings.TrimSpace(token[0])
		method:="GET"
		headers:=make(map[string]string)
		payload := new(strings.Reader)
		client:=lib.CallAPI{url,method,headers,payload,false}
		response,err:=client.CallRest()
		
		userDTO:=dto.UserDTO{}
		err=json.NewDecoder(response.Body).Decode(&userDTO)
		fmt.Println("USER DTO ",userDTO)
		if err != nil{
			a.Token.Status.Authenticated=false
			a.Token.Status.Error="user-management response is not correct!"
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(&a)
			
		}else{
			if userDTO.Password == token[1]{
				fmt.Println("userDTO password and token is the same")
				a.Token.Status.Authenticated=true
				a.Token.Status.User.Username=userDTO.User
				a.Token.Status.User.UID=strconv.Itoa(userDTO.UID)
				a.Token.Status.User.Groups=nil
				for _,group :=range userDTO.Groups{
					a.Token.Status.User.Groups=append(a.Token.Status.User.Groups,group.Group)
				}
				json.NewEncoder(w).Encode(&a.Token)
			}else{
				fmt.Println("DTO password |",userDTO.Password,"| token password: |",token[1],"|")
				a.Token.Status.Authenticated=false
				a.Token.Status.Error="User or Password is not correct!"
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(&a.Token)
			}
		}
	}
}
