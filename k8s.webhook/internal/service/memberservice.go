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
type MemberService struct {
		User string `json:"user" example:"user1"`
		Group string `json:"group" example:"dev"`
}
/*
// FindByName godoc
// @Summary Retrieve a group by name
// @Description Get details of a group
// @Tags groups
// @Accept  json
// @Produce  json
// @Success 200 {object} Groups
// @Router /api/v1/groups/{group} [get]
// FindByName is a method to retrieve an Group by his name
func (m *MemberService)FindByName(w http.ResponseWriter, r *http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	vars := mux.Vars(r)
	group := vars["group"]
	user :=vars["user"]
	groupRepo :=domain.Group{Group:group}
	userRepo :=domain.User{User:user}
	w.Header().Set("Content-Type", "application/json")
	result:=groupRepo.FindByName(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		result=userRepo.FindByName(pg.Database)
		if result.Error != nil{
			var e format.Error
			e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
			json.NewEncoder(w).Encode(e)
			w.WriteHeader(http.StatusBadRequest)
		}else{

			json.NewEncoder(w).Encode(m)
	}
}
*/
// Save godoc
// @Summary Create a group member
// @Description Create a group member
// @Tags members
// @Accept  json
// @Produce  json
// @Success 200 {object} Members
// @Router /api/v1/groups/{group}/members/ [post]
// Save is a method to implement Group member creation logic
func (m *MemberService)Save(w http.ResponseWriter, r * http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	fmt.Println("In Member service")
	vars := mux.Vars(r)
	group := vars["group"]
	_=json.NewDecoder(r.Body).Decode(&m)

	groupRepo :=domain.Group{Group:group}
	userRepo :=domain.User{User:m.User}
	ListGroupRepo :=[]domain.Group{}
	w.Header().Set("Content-Type", "application/json")
	result:=groupRepo.FindByName(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		result=userRepo.FindByName(pg.Database)
		if result.Error != nil{
			var e format.Error
			e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
			json.NewEncoder(w).Encode(e)
			w.WriteHeader(http.StatusBadRequest)
		}else{
			fmt.Println("Group : ", groupRepo.ID)
			fmt.Println("User: ", userRepo.ID)
			ListGroupRepo=append(ListGroupRepo,groupRepo)		
			joinRepo:= domain.User{User:m.User,Groups: ListGroupRepo}

			fmt.Println("Join : ",joinRepo)
			result=joinRepo.AssociationCreate(pg.Database)
			if result.Error != nil{
				var e format.Error
				e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(e)
				
			}else{
				userRepo.FindByName(pg.Database)
				userDTO := dto.UserDTO{}
				userDTO.Convert(userRepo)
				json.NewEncoder(w).Encode(userDTO)
			}
		}
	}
}

// Save godoc
// @Summary Delete a group member
// @Description Delete a group member
// @Tags members
// @Accept  json
// @Produce  json
// @Success 200 {object} Members
// @Router /api/v1/groups/{group}/members/{member} [delete]
// Save is a method to implement Group member deletion logic
func (m *MemberService)Delete(w http.ResponseWriter, r * http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	fmt.Println("In Member service")
	vars := mux.Vars(r)
	group := vars["group"]
	user := vars["member"]

	groupRepo :=domain.Group{Group:group}
	userRepo :=domain.User{User:user}
	ListGroupRepo :=[]domain.Group{}
	w.Header().Set("Content-Type", "application/json")
	result:=groupRepo.FindByName(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		result=userRepo.FindByName(pg.Database)
		if result.Error != nil{
			var e format.Error
			e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
			json.NewEncoder(w).Encode(e)
			w.WriteHeader(http.StatusBadRequest)
		}else{
			fmt.Println("Group : ", groupRepo.ID)
			fmt.Println("User: ", userRepo.ID)
			ListGroupRepo=append(ListGroupRepo,groupRepo)		
			joinRepo:= domain.User{User:user,Groups: ListGroupRepo}

			fmt.Println("Join : ",joinRepo)
			result=joinRepo.AssociationDelete(pg.Database)
			if result.Error != nil{
				var e format.Error
				e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(e)
				
			}else{
				userDTO := dto.UserDTO{}
				userDTO.Convert(userRepo)
				json.NewEncoder(w).Encode(userDTO)
			}
		}
	}
}

// Save godoc
// @Summary Find a group member by name
// @Description find a group member by name
// @Tags members
// @Accept  json
// @Produce  json
// @Success 200 {object} Members
// @Router /api/v1/groups/{group}/members/{member} [get]
// Save is a method to implement Group member retrieving logic
func (m *MemberService)FindByName(w http.ResponseWriter, r * http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	fmt.Println("In Member service")
	vars := mux.Vars(r)
	group := vars["group"]
	user := vars["member"]

	groupRepo :=domain.Group{Group:group}
	userRepo :=domain.User{User:user}
	ListGroupRepo :=[]domain.Group{}
	w.Header().Set("Content-Type", "application/json")
	result:=groupRepo.FindByName(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		result=userRepo.FindByName(pg.Database)
		if result.Error != nil{
			var e format.Error
			e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
			json.NewEncoder(w).Encode(e)
			w.WriteHeader(http.StatusBadRequest)
		}else{
			fmt.Println("Group : ", groupRepo.ID)
			fmt.Println("User: ", userRepo.ID)
			ListGroupRepo=append(ListGroupRepo,groupRepo)		
			joinRepo:= domain.User{User:user,Groups: ListGroupRepo}

			fmt.Println("Join : ",joinRepo)
			result=joinRepo.AssociationFindByName(pg.Database)
			if result.Error != nil{
				var e format.Error
				e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(e)
				
			}else{
				userDTO := dto.UserDTO{}
				userDTO.Convert(userRepo)
				json.NewEncoder(w).Encode(userDTO)
			}
		}
	}
}

// Save godoc
// @Summary Find all group members
// @Description find all group members
// @Tags members
// @Accept  json
// @Produce  json
// @Success 200 {object} Members
// @Router /api/v1/groups/{group}/members [get]
// Save is a method to implement all Group member retrieving logic
func (m *MemberService)FindAll(w http.ResponseWriter, r * http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	fmt.Println("In Member service")
	vars := mux.Vars(r)
	group := vars["group"]
	user := vars["member"]

	groupRepo :=domain.Group{Group:group}
	userRepo :=domain.User{User:user}
	ListGroupRepo :=[]domain.Group{}
	w.Header().Set("Content-Type", "application/json")
	result:=groupRepo.FindByName(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		result=userRepo.FindByName(pg.Database)
		if result.Error != nil{
			var e format.Error
			e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
			json.NewEncoder(w).Encode(e)
			w.WriteHeader(http.StatusBadRequest)
		}else{
			fmt.Println("Group : ", groupRepo.ID)
			fmt.Println("User: ", userRepo.ID)
			ListGroupRepo=append(ListGroupRepo,groupRepo)		
			joinRepo:= domain.User{User:user,Groups: ListGroupRepo}

			fmt.Println("Join : ",joinRepo)
			result=joinRepo.AssociationFindByName(pg.Database)
			if result.Error != nil{
				var e format.Error
				e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(e)
				
			}else{
				userDTO := dto.UserDTO{}
				userDTO.Convert(userRepo)
				json.NewEncoder(w).Encode(userDTO)
			}
		}
	}
}
