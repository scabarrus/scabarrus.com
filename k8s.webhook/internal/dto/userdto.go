package dto

import (
	"scabarrus.com/k8s.webhook/internal/domain"
)

// UserDTO is a struct for DTO User
type UserDTO struct {
	UID int `json:"uid" mandatory:"true"`
	User string `json:"user" mandatory:"true"`
	Password string `json:"password" mandatory:"true"`
	Groups []GroupDTO

}

func (u *UserDTO)Convert(i interface{}){

	u.UID=i.(domain.User).UID
	u.User=i.(domain.User).User
	u.Password=i.(domain.User).Password
	for _,group :=range i.(domain.User).Groups{
		tmp:=GroupDTO{Group:group.Group,GID:group.GID,Description:group.Description}
		u.Groups=append(u.Groups, tmp)
	}

}