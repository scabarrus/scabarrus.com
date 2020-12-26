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

// UserService is a struct that provide business service implementation
type UserService struct {
		UID int `json:"uid" example:"5000"`
		User string `json:"user" example:"scabarrus"`
		Password string `json:"password" example:"B67zuopX#2"`

}


// FindAll godoc
// @Summary Retrieve all user 
// @Description Get details of a users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} User
// @Router /api/v1/users/{user} [get]
// FindAll is a method to retrieve all User 
func (u *UserService)FindAll(w http.ResponseWriter, r *http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	
	usersRepo :=domain.User{}
	w.Header().Set("Content-Type", "application/json")
	result,userList:=usersRepo.FindAll(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		listDTO := []dto.UserDTO{}
		for _,user := range userList{
			userDTO := dto.UserDTO{}
			userDTO.Convert(user)
			listDTO=append(listDTO,userDTO)
		}
		
		json.NewEncoder(w).Encode(listDTO)
	}
}
// FindByName godoc
// @Summary Retrieve a user by name
// @Description Get details of a user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} User
// @Router /api/v1/users/{user} [get]
// FindByName is a method to retrieve an User by his name
func (u *UserService)FindByName(w http.ResponseWriter, r *http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	vars := mux.Vars(r)
	user := vars["user"]
	userRepo :=domain.User{User:user}
	w.Header().Set("Content-Type", "application/json")
	result:=userRepo.FindByName(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		userDTO := dto.UserDTO{}
		userDTO.Convert(userRepo)
		json.NewEncoder(w).Encode(userDTO)
	}
}

// Save godoc
// @Summary Create a user 
// @Description Create a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body User true "Create user"
// @Success 200 {object} User
// @Router /api/v1/users/ [post]
// Save is a method to implement User creation logic
func (u *UserService)Save(w http.ResponseWriter, r * http.Request){
	pg := repository.Postgres{}
	userDTO := dto.UserDTO{}
	pg.Initialization()
		_=json.NewDecoder(r.Body).Decode(&userDTO)
	fmt.Println("userDTO : ",userDTO)
	userRepo := domain.User{}.DTO(userDTO.UID,userDTO.User,userDTO.Password)
	w.Header().Set("Content-Type", "application/json")
	result:=userRepo.Save(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
	}else{
		userDTO.Convert(userRepo)
		json.NewEncoder(w).Encode(userDTO)
	}
} 


// Modify godoc
// @Summary Modify a user 
// @Description Modify a user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} User
// @Router /api/v1/users/{user} [put]
// Modify is a method to modify an User by his name
func (u *UserService)Modify(w http.ResponseWriter, r *http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	userDTO := dto.UserDTO{}
	
	_=json.NewDecoder(r.Body).Decode(&userDTO)


	userRepo := domain.User{}.DTO(userDTO.UID,userDTO.User,userDTO.Password)
	w.Header().Set("Content-Type", "application/json")
	result:=userRepo.Modify(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{	
		userDTO.Convert(userRepo)
		json.NewEncoder(w).Encode(userDTO)
	}
}

// Delete godoc
// @Summary Delete a user 
// @Description Delete a user
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} User
// @Router /api/v1/users/{user} [delete]
// Delete is a method to delete an User by his name
func (u *UserService)Delete(w http.ResponseWriter, r *http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	vars := mux.Vars(r)
	user := vars["user"]
	userRepo :=domain.User{User:user}
	w.Header().Set("Content-Type", "application/json")
	result:=userRepo.Delete(pg.Database)
	if result.Error != nil{
		var e format.Error
		e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
		json.NewEncoder(w).Encode(e)
		w.WriteHeader(http.StatusBadRequest)
	}else{
		userDTO := dto.UserDTO{}
		userDTO.Convert(userRepo)
		json.NewEncoder(w).Encode(userDTO)
	}
}
