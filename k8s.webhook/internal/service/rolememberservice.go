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

// RoleService is a struct that provide business service implementation
type RoleMemberService struct {
		Role string `json:"role" example:"role1"`
		Group string `json:"group" example:"dev"`
}

// Save godoc
// @Summary Create a member's group
// @Description create a member's group
// @Tags roles
// @Accept  json
// @Produce  json
// @Param role path string true "role name"
// @Param member body RoleMemberService true "payload"
// @Success 200 {object} dto.RoleDTO true "dto"
// @Success 400 {object} format.Error
// @Router /roles/{role}/members [post]
func (m *RoleMemberService)Save(w http.ResponseWriter, r * http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	fmt.Println("In Member service")
	vars := mux.Vars(r)
	role := vars["role"]
	_=json.NewDecoder(r.Body).Decode(&m)

	roleRepo :=domain.Role{Role:role}
	groupRepo :=domain.Group{Group:m.Group}
	ListGroupRepo :=[]domain.Group{}
	w.Header().Set("Content-Type", "application/json")
	result:=roleRepo.FindByName(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error 1 - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		result=groupRepo.FindByName(pg.Database)
		if result.Error != nil{
			var e format.Error
			e.FormatError("SQL Error 2 - ",result.Error.Error(),r.RequestURI)
			json.NewEncoder(w).Encode(e)
			w.WriteHeader(http.StatusBadRequest)
		}else{
			
			ListGroupRepo=append(ListGroupRepo,groupRepo)
			fmt.Println("role repo : ",roleRepo.Role)		
			joinRepo:= domain.Role{Role:roleRepo.Role,Groups: ListGroupRepo}

			fmt.Println("Join : ",joinRepo)
			result=joinRepo.AssociationCreate(pg.Database)
			if result.Error != nil{
				var e format.Error
				e.FormatError("SQL Error 3 - ",result.Error.Error(),r.RequestURI)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(e)
				
			}else{
				roleRepo.FindByName(pg.Database)
				roleDTO := dto.RoleDTO{}
				roleDTO.Convert(roleRepo)
				json.NewEncoder(w).Encode(roleDTO)
			}
		}
	}
}

// Save godoc
// @Summary Create a member's group
// @Description create a member's group
// @Tags roles
// @Accept  json
// @Produce  json
// @Param role path string true "role name"
// @Success 200 {object} dto.RoleDTO true "dto"
// @Success 400 {object} format.Error
// @Router /roles/{role}/members [delete]
func (m *RoleMemberService)Delete(w http.ResponseWriter, r * http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	fmt.Println("In Member service")
	vars := mux.Vars(r)
	role := vars["role"]
	group := vars["member"]

	roleRepo :=domain.Role{Role:role}
	groupRepo :=domain.Group{Group:group}
	listGroupRepo :=[]domain.Group{}
	w.Header().Set("Content-Type", "application/json")
	result:=roleRepo.FindByName(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		result=groupRepo.FindByName(pg.Database)
		if result.Error != nil{
			var e format.Error
			e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
			json.NewEncoder(w).Encode(e)
			w.WriteHeader(http.StatusBadRequest)
		}else{
			
			listGroupRepo=append(listGroupRepo,groupRepo)		
			joinRepo:= domain.Role{Role:role,Groups: listGroupRepo}

			fmt.Println("Join : ",joinRepo)
			err:=joinRepo.AssociationDelete(pg.Database,groupRepo)
			if err != nil{
				var e format.Error
				e.FormatError("SQL Error - ",err.Error(),r.RequestURI)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(e)
				
			}else{
				roleDTO := dto.RoleDTO{}
				roleDTO.Convert(joinRepo)
				json.NewEncoder(w).Encode(roleDTO)
			}
		}
	}
}

// FindByName godoc
// @Summary Find a member's role by it's name
// @Description find a member's role by it's name
// @Tags roles
// @Accept  json
// @Produce  json
// @Param role path string true "role name"
// @Param member path string true "member name"
// @Success 200 {object} dto.RoleDTO true "dto"
// @Success 400 {object} format.Error
// @Router /roles/{role}/members/{member} [get]
func (m *RoleMemberService)FindByName(w http.ResponseWriter, r * http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	fmt.Println("In Member service")
	vars := mux.Vars(r)
	role := vars["role"]
	group := vars["member"]

	roleRepo :=domain.Role{Role:role}
	groupRepo :=domain.Group{Group:group}
	ListGroupRepo :=[]domain.Group{}
	w.Header().Set("Content-Type", "application/json")
	result:=roleRepo.FindByName(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		result=groupRepo.FindByName(pg.Database)
		if result.Error != nil{
			var e format.Error
			e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
			json.NewEncoder(w).Encode(e)
			w.WriteHeader(http.StatusBadRequest)
		}else{
			
			ListGroupRepo=append(ListGroupRepo,groupRepo)		
			joinRepo:= domain.Role{Role:role,Groups: ListGroupRepo}

			fmt.Println("Join : ",joinRepo)
			result=joinRepo.AssociationFindByName(pg.Database)
			if result.Error != nil{
				var e format.Error
				e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(e)
				
			}else{
				roleDTO := dto.RoleDTO{}
				roleDTO.Convert(roleRepo)
				json.NewEncoder(w).Encode(roleDTO)
			}
		}
	}
}

// FindAll godoc
// @Summary Find all roles 
// @Description find all roles 
// @Tags roles
// @Accept  json
// @Produce  json
// @Param role path string true "role name"
// @Success 200 {object} []dto.RoleDTO true "dto"
// @Success 400 {object} format.Error
// @Router /roles/{role}/members [get]
func (m *RoleMemberService)FindAll(w http.ResponseWriter, r * http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	fmt.Println("In Member service")
	vars := mux.Vars(r)
	role := vars["role"]
	group := vars["member"]

	roleRepo :=domain.Role{Role:role}
	groupRepo :=domain.Group{Group:group}
	listGroupRepo :=[]domain.Group{}
	w.Header().Set("Content-Type", "application/json")
	result:=roleRepo.FindByName(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		result=roleRepo.FindByName(pg.Database)
		if result.Error != nil{
			var e format.Error
			e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
			json.NewEncoder(w).Encode(e)
			w.WriteHeader(http.StatusBadRequest)
		}else{
			
			listGroupRepo=append(listGroupRepo,groupRepo)		
			joinRepo:= domain.Role{Role:role,Groups: listGroupRepo}

			fmt.Println("Join : ",joinRepo)
			result=joinRepo.AssociationFindByName(pg.Database)
			if result.Error != nil{
				var e format.Error
				e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(e)
				
			}else{
				roleDTO := dto.RoleDTO{}
				roleDTO.Convert(roleRepo)
				json.NewEncoder(w).Encode(roleDTO)
			}
		}
	}
}
