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

type RoleService struct {
	Role string `json:"role" example:"role1" mandatory:"true"`
	Namespace string `json:"namespace" example:"default" mandatory:"true"`
	Verb string `json:"verb" example:"list" mandatory:"true"`
	Group string `json:"group" example:"apps" mandatory:"true"`
	Resource string `json:"resource" example:"namespace" mandatory:"true"`
	Groups GroupService
}

 // FindAll godoc
// @Summary Show all roles
// @Description get all roles
// @Tags roles
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.RoleDTO true "dto"
// @Success 400 {object} format.Error
// @Router /roles [get]
// FindAll find all role 
// It return a json payload with list of role details or an error message
func (r *RoleService)FindAll(w http.ResponseWriter, req *http.Request){
	//Instanciate pg repository
	pg := repository.Postgres{}
	//Initialize the postgres repository
	pg.Initialization()
	rolesRepo :=domain.Role{}
	w.Header().Set("Content-Type", "application/json")
	result,roleList:=rolesRepo.FindAll(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),req.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		listDTO := []dto.RoleDTO{}
		for _,role := range roleList{
			roleDTO := dto.RoleDTO{}
			roleDTO.Convert(role)
			listDTO=append(listDTO,roleDTO)
		}
		
		json.NewEncoder(w).Encode(listDTO)
	}
}

// FindByName godoc
// @Summary Show a role details
// @Description get a role by it's name
// @Tags roles
// @Accept  json
// @Produce  json
// @Param role path string true "role name"
// @Success 200 {object} dto.RoleDTO true "dto"
// @Success 400 {object} format.Error
// @Router /roles/{role} [get]
// Find retrieve a role by name
// Role name is in URI 
// It return a json payload with role details or an error message
func (r *RoleService)FindByName(w http.ResponseWriter, req *http.Request){
	//Instanciate pg repository
	pg := repository.Postgres{}
	//Initialize the postgres repository
	pg.Initialization()
	vars := mux.Vars(req)
	role := vars["role"]
	roleRepo :=domain.Role{Role:role}
	w.Header().Set("Content-Type", "application/json")
	result:=roleRepo.FindByName(pg.Database)
	// if an error appears when trying to find a role by it's name
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),req.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
	}else{ //Display role information in dto format
		roleDTO := dto.RoleDTO{}
		fmt.Println("find by name response: ",roleRepo)
		roleDTO.Convert(roleRepo)
		json.NewEncoder(w).Encode(roleDTO)
	}
}


// Save godoc
// @Summary Create a role 
// @Description create a role 
// @Tags roles
// @Accept  json
// @Produce  json
// @Param role body dto.RoleDTO true "dto"
// @Success 200 {object} dto.RoleDTO true "dto"
// @Success 400 {object} format.Error
// @Router /roles [post]
// Save create a role
// Role details is in payload
// It return a json payload with role details or an error message
func (r *RoleService)Save(w http.ResponseWriter, req * http.Request){
	//Instanciate pg repository
	pg := repository.Postgres{}
	//Initialize the postgres repository
	pg.Initialization()
	roleDTO := dto.RoleDTO{}
	var e format.Error
	_=json.NewDecoder(req.Body).Decode(&roleDTO)
	roleRepo := domain.Role{}.DTO(roleDTO.Role,roleDTO.Namespace,roleDTO.Verb,roleDTO.Group,roleDTO.Resource)
	w.Header().Set("Content-Type", "application/json")
	// control no mandatory attribute is missing
	message,details,_ :=e.Unmarshal(&roleDTO)
	// if a mandatory attribut is missed
    if message != "" {
		e.FormatError(message,details,req.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
	}else{ // Save the new role
		result:=roleRepo.Save(pg.Database)
		//if an error appears when role is inserted
		if result.Error != nil{
			e.FormatError("SQL Error - ",result.Error.Error(),req.RequestURI)
			json.NewEncoder(w).Encode(e)
		}else{ // return the role payload in dto format
			roleDTO.Convert(roleRepo)
			json.NewEncoder(w).Encode(roleDTO)
		}
	}
} 


// Modify godoc
// @Summary Modify a role details
// @Description modify a role 
// @Tags roles
// @Accept  json
// @Produce  json
// @Param role body dto.RoleDTO true "dto"
// @Param role path string true "role name"
// @Success 200 {object} dto.RoleDTO true "dto"
// @Success 400 {object} format.Error
// @Router /roles/{role} [put]
// Modify update a role
// Role name is in URI 
// It return a json payload with role details or an error message
func (r *RoleService)Modify(w http.ResponseWriter, req *http.Request){
	//Instanciate pg repository
	pg := repository.Postgres{}
	//Initialize the postgres repository
	pg.Initialization()
	roleDTO := dto.RoleDTO{}
	
	_=json.NewDecoder(req.Body).Decode(&roleDTO)
	roleRepo := domain.Role{}.DTO(roleDTO.Role,roleDTO.Namespace,roleDTO.Verb,roleDTO.Group,roleDTO.Resource)
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	role := vars["role"]
	var e format.Error
	// if role in URI not match with URI in body
	if r.Role != roleDTO.Role{
			e.FormatError("input Error - ","mismatch between role name in path ("+role+") and body ("+roleDTO.Role+")",req.RequestURI)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)
	}
	// Controle no mandatory attributes is missing
	message,details,_ :=e.Unmarshal(&roleDTO)
    if message != "" {
		e.FormatError(message,details,req.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
	}else{ // Modify roles
		result:=roleRepo.Modify(pg.Database)
		// if an error appears during DB update
		if result.Error != nil{
			e.FormatError("SQL Error - ",result.Error.Error(),req.RequestURI)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)
		}else{	 // if a non editable attribute is changed
			if result.RowsAffected == 0{
				e.FormatError("Input Error - ","Non editable field are modified!",req.RequestURI)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(e)
			}else{ // return role in dto format
				roleDTO.Convert(roleRepo)
				json.NewEncoder(w).Encode(roleDTO)
			}
		}
	}
}

// Delete godoc
// @Summary Delete a role details
// @Description delete a role 
// @Tags roles
// @Accept  json
// @Produce  json
// @Param role path string true "role name"
// @Success 200 {object} dto.RoleDTO true "dto"
// @Success 400 {object} format.Error
// @Router /roles/{role} [delete]
// Delete remove a role
// Role name is in URI 
// It return a json payload with role details or an error message
func (r *RoleService)Delete(w http.ResponseWriter, req *http.Request){
	//Instanciate pg repository
	pg := repository.Postgres{}
	//Initialize the postgres repository
	pg.Initialization()
	vars := mux.Vars(req)
	role := vars["role"]
	roleRepo :=domain.Role{Role:role}
	w.Header().Set("Content-Type", "application/json")
	result:=roleRepo.Delete(pg.Database)
	// if an error appears during DB delete
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),req.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
	}else{ // Delete role in DB and return role in DTO
		roleDTO := dto.RoleDTO{}
		roleDTO.Convert(roleRepo)
		json.NewEncoder(w).Encode(roleDTO)
	}
}
