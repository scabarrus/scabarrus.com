package domain

import (
	"gorm.io/gorm"
)

//User struct is User entity
type User struct {
	gorm.Model
	UID      int    `gorm:"unique"` 
	User     string `gorm:"primaryKey;unique"`
	Password string
	Groups   []Group `gorm:"many2many:user_group;"`
	ID uint 
}

func (u User)DTO(uid int,user string, password string)(User){
	return User{UID:uid,User:user,Password:password}
}

//Save is method to save new User
func (u *User) Save(db *gorm.DB)(* gorm.DB){
	return db.Debug().Create(&u)
		
}

func (u *User) AssociationCreate(db *gorm.DB)(* gorm.DB){
	//return db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&u)
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Model(&u).Updates(&u)

}

func (u *User) AssociationDelete(db *gorm.DB)(* gorm.DB){
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Model(&u).Delete(&u)
}

func (u *User) AssociationFindByName(db *gorm.DB)(* gorm.DB){
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Model(&u).Find(&u)
}
//FindAll is method to find all users
func (u *User) FindAll(db *gorm.DB)(* gorm.DB,[]User){
	userList:=[]User{}
	result:=db.Debug().Find(&userList)
	return result,userList
}

//FindByName is method to retrieve an User by his name
func (u *User) FindByName(db *gorm.DB)(* gorm.DB){
	return db.Debug().Preload("Groups").Where("\"user\"=?",u.User).First(&u)
}

//Modify is method to update an User
func (u *User) Modify(db *gorm.DB)(* gorm.DB){
	return db.Debug().Where("\"uid\"=?",u.UID).Updates(&u)
}

//Delete is method to remove an User
func (u *User) Delete(db *gorm.DB)(* gorm.DB){
	return db.Debug().Where("\"user\"=?",u.User).Delete(&u)
}

func (u *User)DBMigrate(db *gorm.DB)(* gorm.DB){
	db.AutoMigrate(&User{})
	return db
}