package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	authz "k8s.io/kubernetes/pkg/apis/authorization"
	"scabarrus.com/k8s.webhook/internal/dto"
	lib "scabarrus.com/k8s.webhook/internal/utils"
)
type Authz struct{
	Subject authz.SubjectAccessReview
}

//GetUserGroup look for user's group
func (a Authz) GetUserGroup(user string, api string)(dto.UserDTO){
	url:=api+"api/v1/users/"+user
	method:="GET"
	headers:=make(map[string]string)
	payload := new(strings.Reader)
	client:=lib.CallAPI{url,method,headers,payload,false}
	response,_:=client.CallRest()
	
	userDTO:=dto.UserDTO{}
	_=json.NewDecoder(response.Body).Decode(&userDTO)
	return userDTO
}

//CheckRole check if one role provide right to perform the action
func (a Authz) CheckRole(api string)(bool){
	
	for _,member := range a.Subject.Spec.Groups{
		url:=api+"api/v1/groups/"+member+"/roles"
		method:="GET"
		headers:=make(map[string]string)
		payload := new(strings.Reader)
		client:=lib.CallAPI{url,method,headers,payload,false}
		response,_:=client.CallRest()
		groupDTO:=dto.GroupDTO{}
		_=json.NewDecoder(response.Body).Decode(&groupDTO)
		//if group are linked to a role
		if groupDTO.Roles != nil {
			//for each role
			for _,role:= range groupDTO.Roles{
				fmt.Println("role: |",role.Namespace,"| request |",a.Subject.Spec.ResourceAttributes.Namespace,"|")
				//check if namespace match with request
				//&& role.Group==a.Subject.Spec.ResourceAttributes.Group && role.Resource==a.Subject.Spec.ResourceAttributes.Resource
				if strings.Contains(role.Namespace,a.Subject.Spec.ResourceAttributes.Namespace)  {
						fmt.Println("find a matching namespace in role")
						if strings.Contains(role.Verb,a.Subject.Spec.ResourceAttributes.Verb) {
							fmt.Println("find matching verb in role")
							if strings.Contains(role.Resource,a.Subject.Spec.ResourceAttributes.Resource) {
								fmt.Println("find matching resource in role")
								//if  strings.Contains(role.Group,a.Subject.Spec.ResourceAttributes.Group) {
								//	fmt.Println("find matching group in role")
									return true
								//}
						}
					}
				}
			}
		}
	
	}
	return false
}
func (a Authz)CheckAccess(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	_=json.NewDecoder(r.Body).Decode(&a.Subject)
	fmt.Println("Subject ",a.Subject)
	api:=os.Getenv("USR_API")
	_=a.GetUserGroup(a.Subject.Spec.User,api)
	fmt.Println("groups : ",a.Subject.Spec.Groups)
	access:=a.CheckRole(api)

	
	if access == true{
		a.Subject.Status.Allowed=true
	}else{
		a.Subject.Status.Allowed=false
	}
	json.NewEncoder(w).Encode(a.Subject)
	
}
