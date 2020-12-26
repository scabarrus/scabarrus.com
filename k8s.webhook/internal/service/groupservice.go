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


// FindByAll godoc
// @Summary Retrieve all groups
// @Description Get details of all groups
// @Tags groups
// @Accept  json
// @Produce  json
// @Success 200 {object} Groups
// @Router /api/v1/groups [get]
// FindAll is a method to retrieve all Group 
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
// @Summary Retrieve a group by name
// @Description Get details of a group
// @Tags groups
// @Accept  json
// @Produce  json
// @Success 200 {object} Groups
// @Router /api/v1/groups/{group} [get]
// FindByName is a method to retrieve an Group by his name
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
// @Description Create a group
// @Tags groups
// @Accept  json
// @Produce  json
// @Success 200 {object} Group
// @Router /api/v1/groups/{group} [post]
// Save is a method to implement Group creation logic
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
// @Description Modify a group
// @Tags Groups
// @Accept  json
// @Produce  json
// @Success 200 {object} Group
// @Router /api/v1/groups/{group} [put]
// Modify is a method to modify an Group by his name
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
// @Description Delete a group
// @Tags groups
// @Accept  json
// @Produce  json
// @Success 200 {object} Group
// @Router /api/v1/groups/{group} [delete]
// Delete is a method to delete an Group by his name
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
