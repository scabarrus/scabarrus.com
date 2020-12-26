package dto

import "scabarrus.com/k8s.webhook/internal/domain"

// GroupDTO is a struct that provide business service implementation
type GroupDTO struct {
	GID int `json:"gid" mandatory:"true"`
	Group string `json:"group" mandatory:"true"`
	Description string `json:"description"`

}

func (g *GroupDTO)Convert(i interface{}){

	g.GID=i.(domain.Group).GID
	g.Group=i.(domain.Group).Group
	g.Description=i.(domain.Group).Description

}