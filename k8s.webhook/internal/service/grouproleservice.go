package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"scabarrus.com/k8s.webhook/internal/domain"
	"scabarrus.com/k8s.webhook/internal/dto"
	format "scabarrus.com/k8s.webhook/internal/error"
	"scabarrus.com/k8s.webhook/internal/repository"
)

// GroupService is a struct that provide business service implementation
type GroupRoleService struct {
	Role string `json:"user" example:"user1" `
	Group string `json:"group" example:"operator" mandatory:"true"`
}


// FindByName godoc
// @Summary Find a role's group by it's name
// @Description find a role's group by it's name
// @Tags groups
// @Accept  json
// @Produce  json
// @Param group path string true "group name"
// @Success 200 {object} dto.GroupDTO true "dto"
// @Success 400 {object} format.Error
// @Router /groups/{group}/roles [get]
func (m *GroupRoleService)FindByName(w http.ResponseWriter, r * http.Request){
	//Instanciate pg repository
	pg := repository.Postgres{}
	//Initialize the postgres repository
	pg.Initialization()
	vars := mux.Vars(r)
	group := vars["group"]

	groupRepo :=domain.Group{Group:group}
	//ListRoleRepo :=[]domain.Role{}
	w.Header().Set("Content-Type", "application/json")
	result:=groupRepo.FindByName(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
	}else{
			groupRepo.AssociationRoleGroupByName(pg.Database)
			fmt.Println("group Repo ",groupRepo)
			groupDTO:=dto.GroupDTO{}
			groupDTO.Convert(groupRepo)
			json.NewEncoder(w).Encode(groupDTO)
		}
}