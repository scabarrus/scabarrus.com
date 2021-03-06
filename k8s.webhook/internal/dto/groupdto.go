package dto

import "scabarrus.com/k8s.webhook/internal/domain"

// GroupDTO is a struct that provide business service implementation
type GroupDTO struct {
	GID int `json:"gid" mandatory:"true" example:"7001"`
	Group string `json:"group" mandatory:"true" example:"operator"`
	Description string `json:"description" example:"Operator group"`
	Users []UserDTO
	Roles []RoleDTO
}

func (g *GroupDTO)Convert(i interface{}){

	g.GID=i.(domain.Group).GID
	g.Group=i.(domain.Group).Group
	g.Description=i.(domain.Group).Description
	for _,role :=range i.(domain.Group).Roles{
		tmp:=RoleDTO{Role:role.Role,Namespace:role.Namespace,Verb:role.Verb,Resource:role.Resource}
		g.Roles=append(g.Roles, tmp)
	}
}