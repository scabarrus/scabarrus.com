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
		Users []UserService
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
// FindAll find all group
// It return a json payload with list of group details or an error message
func (u *GroupService)FindAll(w http.ResponseWriter, r *http.Request){
	//Instanciate pg repository
	pg := repository.Postgres{}
	//Initialize the postgres repository
	pg.Initialization()
	groupRepo :=domain.Group{}
	w.Header().Set("Content-Type", "application/json")
	result,groupList:=groupRepo.FindAll(pg.Database)
	// if select on all group failed
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
	}else{ // return list of group
		listDTO := []dto.GroupDTO{}
		// parse groups and append them in response
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
// FindByName find a group by it's name
// Group name is in URI 
// It return a json payload with group details or an error message
func (u *GroupService)FindByName(w http.ResponseWriter, r *http.Request){
	//Instanciate pg repository
	pg := repository.Postgres{}
	//Initialize the postgres repository
	pg.Initialization()
	vars := mux.Vars(r)
	group := vars["group"]
	groupRepo :=domain.Group{Group:group}
	w.Header().Set("Content-Type", "application/json")
	result:=groupRepo.FindByName(pg.Database)
	// if a select by name return an error
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
	}else{ // return group details in dto format
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
// Save create a group
// It return a json payload with group details or an error message
func (u *GroupService)Save(w http.ResponseWriter, r * http.Request){
	//Instanciate pg repository
	pg := repository.Postgres{}
	//Initialize the postgres repository
	pg.Initialization()
	groupDTO := dto.GroupDTO{}
	initDB:= domain.Group{}
	initDB.DBMigrate(pg.Database)
	
	_=json.NewDecoder(r.Body).Decode(&groupDTO)
	var e format.Error
	groupRepo := domain.Group{}.DTO(groupDTO.GID,groupDTO.Group,groupDTO.Description)
	w.Header().Set("Content-Type", "application/json")

	// control mandatory attributes are filled 
	message,details,_ :=e.Unmarshal(&groupDTO)

	// if a mandatory filed is not set
    if message != "" {
		e.FormatError(message,details,r.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
	}else{ // create in DB
		result:=groupRepo.Save(pg.Database)
		// if a create query failed
		if result.Error != nil{
			e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)
		}else{ // return group in dto format
			groupDTO.Convert(groupRepo)
			json.NewEncoder(w).Encode(groupDTO)
		}
	}
} 

// Modify godoc
// @Summary Modify a group
// @Description modify a group
// @Tags groups
// @Accept  json
// @Produce  json
// @Param groupdto body dto.GroupDTO true "dto"
// @Param group path string true "group name"
// @Success 200 {object} dto.GroupDTO true "dto"
// @Success 400 {object} format.Error
// @Router /groups/{group} [put]
// Modify update a group
// Group name is in URI 
// It return a json payload with group details or an error message
func (u *GroupService)Modify(w http.ResponseWriter, r *http.Request){
	//Instanciate pg repository
	pg := repository.Postgres{}
	//Initialize the postgres repository
	pg.Initialization()
	groupDTO := dto.GroupDTO{}
	_=json.NewDecoder(r.Body).Decode(&groupDTO)
	vars := mux.Vars(r)
	group := vars["group"]
	var e format.Error
	w.Header().Set("Content-Type", "application/json")
	// if group in uri not match with group in body
	if group != groupDTO.Group{
		e.FormatError("input Error - ","mismatch between group name in path ("+group+") and body ("+groupDTO.Group+")",r.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
	}

	// control all mandatories attributes are filled
	message,details,_ :=e.Unmarshal(&groupDTO)
	// send an error if a mandatory attribute is missing
    if message != "" {
		e.FormatError(message,details,r.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
	}else{ // Modify the group
		groupRepo := domain.Group{}.DTO(groupDTO.GID,groupDTO.Group,groupDTO.Description)
		result:=groupRepo.Modify(pg.Database)
		// if update in DB return an error 
		if result.Error != nil{
			var e format.Error
			e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)
		}else { // if update happens
			//if no row is affected, means a non editable attributes was modified
			if result.RowsAffected == 0{
				e.FormatError("Input Error - ","Non editable field are modified!",r.RequestURI)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(e)
			}else{ // return group details in dto format
				groupDTO.Convert(groupRepo)
				json.NewEncoder(w).Encode(groupDTO)
			}
		}
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
// Delete remove a group
// Group name is in URI 
// It return a json payload with group details or an error message
func (u *GroupService)Delete(w http.ResponseWriter, r *http.Request){
	//Instanciate pg repository
	pg := repository.Postgres{}
	//Initialize the postgres repository
	pg.Initialization()
	vars := mux.Vars(r)
	group := vars["group"]
	groupRepo :=domain.Group{Group:group}
	w.Header().Set("Content-Type", "application/json")
	result:=groupRepo.Delete(pg.Database)
	//if an error appears during DB delete query
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
	}else{// return deleted group details
		groupDTO := dto.GroupDTO{}
		groupDTO.Convert(groupRepo)
		json.NewEncoder(w).Encode(groupDTO)
	}

}
