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
func (r *RoleService)FindAll(w http.ResponseWriter, req *http.Request){
	pg := repository.Postgres{}
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
func (r *RoleService)FindByName(w http.ResponseWriter, req *http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	vars := mux.Vars(req)
	role := vars["role"]
	roleRepo :=domain.Role{Role:role}
	w.Header().Set("Content-Type", "application/json")
	result:=roleRepo.FindByName(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),req.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
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
func (r *RoleService)Save(w http.ResponseWriter, req * http.Request){
	pg := repository.Postgres{}
	roleDTO := dto.RoleDTO{}
	var e format.Error
	pg.Initialization()
		_=json.NewDecoder(req.Body).Decode(&roleDTO)
	fmt.Println("roleDTO : ",roleDTO)
	roleRepo := domain.Role{}.DTO(roleDTO.Role,roleDTO.Namespace,roleDTO.Verb,roleDTO.Group,roleDTO.Resource)
	w.Header().Set("Content-Type", "application/json")
	message,details,_ :=e.Unmarshal(&roleDTO)
    if message != "" {
		e.FormatError(message,details,req.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
	}else{
		result:=roleRepo.Save(pg.Database)
		if result.Error != nil{
			e.FormatError("SQL Error - ",result.Error.Error(),req.RequestURI)
			json.NewEncoder(w).Encode(e)
		}else{
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
func (r *RoleService)Modify(w http.ResponseWriter, req *http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	roleDTO := dto.RoleDTO{}
	
	_=json.NewDecoder(req.Body).Decode(&roleDTO)
	roleRepo := domain.Role{}.DTO(roleDTO.Role,roleDTO.Namespace,roleDTO.Verb,roleDTO.Group,roleDTO.Resource)
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	role := vars["role"]
	var e format.Error
	if r.Role != roleDTO.Role{
			e.FormatError("input Error - ","mismatch between role name in path ("+role+") and body ("+roleDTO.Role+")",req.RequestURI)
			json.NewEncoder(w).Encode(e)
			w.WriteHeader(http.StatusBadRequest)
	}
	message,details,_ :=e.Unmarshal(&roleDTO)
    if message != "" {
		e.FormatError(message,details,req.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
	}else{
		result:=roleRepo.Modify(pg.Database)
		if result.Error != nil{
			e.FormatError("SQL Error - ",result.Error.Error(),req.RequestURI)
			json.NewEncoder(w).Encode(e)
			w.WriteHeader(http.StatusBadRequest)
		}else{	
			if result.RowsAffected == 0{
				e.FormatError("Input Error - ","Non editable field are modified!",req.RequestURI)
				json.NewEncoder(w).Encode(e)
				w.WriteHeader(http.StatusBadRequest)
			}else{
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
func (r *RoleService)Delete(w http.ResponseWriter, req *http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	vars := mux.Vars(req)
	role := vars["role"]
	roleRepo :=domain.Role{Role:role}
	w.Header().Set("Content-Type", "application/json")
	result:=roleRepo.Delete(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),req.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		roleDTO := dto.RoleDTO{}
		roleDTO.Convert(roleRepo)
		json.NewEncoder(w).Encode(roleDTO)
	}
}
