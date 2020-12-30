package service

import (
	"encoding/json"
	"net/http"

	format "scabarrus.com/k8s.webhook/internal/error"

	"github.com/gorilla/mux"
	"scabarrus.com/k8s.webhook/internal/domain"
	"scabarrus.com/k8s.webhook/internal/dto"
	"scabarrus.com/k8s.webhook/internal/repository"
)

// GroupService is a struct that provide business service implementation
type GroupService struct {
		GID int `json:"gid" example:"7001"`
		Group string `json:"group" example:"operator"`
		Description string `json:"description" example:"Operator group"`

}


// FindAll godoc
// @Summary Show all groups
// @Description get all groups
// @Tags groups
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.GroupDTO true "dto"
// @Success 400 {object} format.Error
// @Router /groups [get]
func (u *GroupService)FindAll(w http.ResponseWriter, r *http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	
	groupRepo :=domain.Group{}
	w.Header().Set("Content-Type", "application/json")
	result,groupList:=groupRepo.FindAll(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		listDTO := []dto.GroupDTO{}
		for _,group:=range groupList{
			groupDTO := dto.GroupDTO{}
			groupDTO.Convert(group)
			listDTO=append(listDTO,groupDTO)

		}
			json.NewEncoder(w).Encode(listDTO)
	}
}

// FindByName godoc
// @Summary Show a group details
// @Description get a group by it's name
// @Tags groups
// @Accept  json
// @Produce  json
// @Param group path string true "group name"
// @Success 200 {object} dto.GroupDTO true "dto"
// @Success 400 {object} format.Error
// @Router /groups/{group} [get]
func (u *GroupService)FindByName(w http.ResponseWriter, r *http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	vars := mux.Vars(r)
	group := vars["group"]
	groupRepo :=domain.Group{Group:group}
	w.Header().Set("Content-Type", "application/json")
	result:=groupRepo.FindByName(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		groupDTO := dto.GroupDTO{}
		groupDTO.Convert(groupRepo)
		json.NewEncoder(w).Encode(groupDTO)
	}
}

// Save godoc
// @Summary Create a group
// @Description create a group
// @Tags groups
// @Accept  json
// @Produce  json
// @Param group body dto.GroupDTO true "dto"
// @Success 200 {object} dto.GroupDTO true "dto"
// @Success 400 {object} format.Error
// @Router /groups [post]
func (u *GroupService)Save(w http.ResponseWriter, r * http.Request){
	pg := repository.Postgres{}
	groupDTO := dto.GroupDTO{}
	pg.Initialization()
	initDB:= domain.Group{}
	initDB.DBMigrate(pg.Database)
	
	_=json.NewDecoder(r.Body).Decode(&groupDTO)

	groupRepo := domain.Group{}.DTO(groupDTO.GID,groupDTO.Group,groupDTO.Description)
	w.Header().Set("Content-Type", "application/json")
	result:=groupRepo.Save(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		groupDTO.Convert(groupRepo)
		json.NewEncoder(w).Encode(groupDTO)
	}
} 

// Modify godoc
// @Summary Modify a group
// @Description modify a group
// @Tags groups
// @Accept  json
// @Produce  json
// @Param group path string true "group name"
// @Success 200 {object} dto.GroupDTO true "dto"
// @Success 400 {object} format.Error
// @Router /groups/{group} [put]
func (u *GroupService)Modify(w http.ResponseWriter, r *http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	groupDTO := dto.GroupDTO{}
	
	_=json.NewDecoder(r.Body).Decode(&groupDTO)


	groupRepo := domain.Group{}.DTO(groupDTO.GID,groupDTO.Group,groupDTO.Description)
	w.Header().Set("Content-Type", "application/json")
	result:=groupRepo.Modify(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else {
		groupDTO.Convert(groupRepo)
		json.NewEncoder(w).Encode(groupDTO)
	}

}

// Delete godoc
// @Summary Delete a group
// @Description delete a group
// @Tags groups
// @Accept  json
// @Produce  json
// @Param group path string true "group name"
// @Success 200 {object} dto.GroupDTO
// @Success 400 {object} format.Error
// @Router /groups/{group} [delete]
func (u *GroupService)Delete(w http.ResponseWriter, r *http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	vars := mux.Vars(r)
	group := vars["group"]
	groupRepo :=domain.Group{Group:group}
	w.Header().Set("Content-Type", "application/json")
	result:=groupRepo.Delete(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		groupDTO := dto.GroupDTO{}
		groupDTO.Convert(groupRepo)
		json.NewEncoder(w).Encode(groupDTO)
	}

}
