package dto

import (
	"scabarrus.com/k8s.webhook/internal/domain"
)

type RoleDTO struct{
	Role string `json:"role" mandatory:"true"`
	Namespace string `json:"namespace"`
	Verb string `json:"verb"`
	Group string `json:"group"`
	Resource string `json:"resource"`
	Groups []GroupDTO
}

func (r *RoleDTO)Convert(i interface{}){

	r.Role=i.(domain.Role).Role
	r.Namespace=i.(domain.Role).Namespace
	r.Verb=i.(domain.Role).Verb
	r.Group=i.(domain.Role).Group
	r.Resource=i.(domain.Role).Resource
	for _,group :=range i.(domain.Role).Groups{
		tmp:=GroupDTO{Group:group.Group,GID:group.GID,Description:group.Description}
		r.Groups=append(r.Groups, tmp)
	}
	

}