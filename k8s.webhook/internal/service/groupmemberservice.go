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
type GroupMemberService struct {
		User string `json:"user" example:"user1" mandatory:"true"`
		Group string `json:"group" example:"operator"`
}


// Save godoc
// @Summary Create a member's group
// @Description create a member's group
// @Tags groups
// @Accept  json
// @Produce  json
// @Param group path string true "group name"
// @Param user body GroupMemberService false "dto"
// @Success 200 {object} dto.UserDTO true "dto"
// @Success 400 {object} format.Error
// @Router /groups/{group}/members [post]
// Save register a group member
// A group member is a user 
// Group is in URI and payload and member in payload only
// It return a json payload with group member details or an error message
func (m *GroupMemberService)Save(w http.ResponseWriter, r * http.Request){
	//Instanciate pg repository
	pg := repository.Postgres{}
	//Initialize the postgres repository
	pg.Initialization()
	fmt.Println("In Member service")
	vars := mux.Vars(r)
	group := vars["group"]
	_=json.NewDecoder(r.Body).Decode(&m)
	var e format.Error
	w.Header().Set("Content-Type", "application/json")
	//Control if group in URI match with payload
	if group != m.Group{
			e.FormatError("input Error - ","mismatch between group name in path ("+group+") and body ("+m.Group+")",r.RequestURI)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)
	}
	//Control mandatory field are not missing
	message,details,_ :=e.Unmarshal(&m)
	//Display an error message if mandatory attribute is missing
    if message != "" {	
		e.FormatError(message,details,r.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
	}else{ //Register member
		groupRepo :=domain.Group{Group:group}
		userRepo :=domain.User{User:m.User}
		ListGroupRepo :=[]domain.Group{}
		result:=groupRepo.FindByName(pg.Database)
		//if an error appears for DB select by name for a group
		if result.Error != nil{
			e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)
		}else{ //Retrieve user information
			result=userRepo.FindByName(pg.Database)
			//if an error appears for DB select by name for a user
			if result.Error != nil{
				e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(e)
			}else{ //Insert in association table the member 
				fmt.Println("Group : ", groupRepo.ID)
				fmt.Println("User: ", userRepo.ID)
				ListGroupRepo=append(ListGroupRepo,groupRepo)		
				joinRepo:= domain.User{User:m.User,Groups: ListGroupRepo}

				fmt.Println("Join : ",joinRepo)
				result=joinRepo.AssociationCreate(pg.Database)
				// if association failed
				if result.Error != nil{
					var e format.Error
					e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(e)
					
				}else{ //else return association details in payload
					userRepo.FindByName(pg.Database)
					userDTO := dto.UserDTO{}
					userDTO.Convert(userRepo)
					json.NewEncoder(w).Encode(userDTO)
				}
			}
		}
	}
}

// Delete godoc
// @Summary Delete a member's group
// @Description Delete a member's group
// @Tags groups
// @Accept  json
// @Produce  json
// @Param group path string true "group name"
// @Param member path string true "member name"
// @Success 200 {object} dto.UserDTO true "dto"
// @Success 400 {object} format.Error
// @Router /groups/{group}/members/{member} [delete]
// Delete remove a member of a group
// Group name is in URI but also in payload
// It return a json payload with group details or an error message
// Delete remove a member of a group
// Group name is in URI 
// It return a json payload with group member details or an error message
func (m *GroupMemberService)Delete(w http.ResponseWriter, r * http.Request){
	//Instanciate pg repository
	pg := repository.Postgres{}
	//Initialize the postgres repository
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
	}else{
		result=userRepo.FindByName(pg.Database)
		if result.Error != nil{
			var e format.Error
			e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)
		}else{
			fmt.Println("Group : ", groupRepo.ID)
			fmt.Println("User: ", userRepo.ID)
			ListGroupRepo=append(ListGroupRepo,groupRepo)		
			joinRepo:= domain.User{User:user,Groups: ListGroupRepo}

			fmt.Println("Join : ",joinRepo)
			err:=joinRepo.AssociationDelete(pg.Database,groupRepo)
			if err != nil{
				var e format.Error
				e.FormatError("SQL Error - ",err.Error(),r.RequestURI)
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

// FindByName godoc
// @Summary Find a member's group by it's name
// @Description find a member's group by it's name
// @Tags groups
// @Accept  json
// @Produce  json
// @Param group path string true "group name"
// @Param member path string true "member name"
// @Success 200 {object} dto.UserDTO true "dto"
// @Success 400 {object} format.Error
// @Router /groups/{group}/members/{member} [get]
func (m *GroupMemberService)FindByName(w http.ResponseWriter, r * http.Request){
	//Instanciate pg repository
	pg := repository.Postgres{}
	//Initialize the postgres repository
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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
	}else{
		result=userRepo.FindByName(pg.Database)
		if result.Error != nil{
			var e format.Error
			e.FormatError("SQL Error - ",result.Error.Error(),r.RequestURI)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(e)
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

// FindAll godoc
// @Summary Find all members 
// @Description find all roles 
// @Tags groups
// @Accept  json
// @Produce  json
// @Param group path string true "group name"
// @Success 200 {object} []dto.GroupDTO true "dto"
// @Success 400 {object} format.Error
// @Router /groups/{group}/members [get]
func (m *GroupMemberService)FindAll(w http.ResponseWriter, r * http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	fmt.Println("In Member service")
	vars := mux.Vars(r)
	group := vars["group"]

	groupRepo :=domain.Group{Group:group}
	userRepo :=[]domain.User{}
	w.Header().Set("Content-Type", "application/json")
	err:=groupRepo.FindAllMember(pg.Database,&userRepo)
	if err != nil{
		var e format.Error
		e.FormatError("SQL Error - ",err.Error(),r.RequestURI)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e)
	}else{
		var userDTO []dto.UserDTO
		for _,user := range userRepo{
			tmp:=dto.UserDTO{}
			tmp.Convert(user)
			userDTO=append(userDTO,tmp)
		}
		json.NewEncoder(w).Encode(userDTO)
	}
	
}
