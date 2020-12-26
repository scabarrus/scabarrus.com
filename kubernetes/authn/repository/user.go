package repository

import (
	"gorm.io/gorm"
)
type User struct{
	//gorm.Model
	User string `gorm:"primaryKey;unique"`
	Password string
	Uid int
	FK_groups []Groups `gorm:"many2many:user_groups;"`
	//FK_groups []Groups `gorm:"association_foreignkey:Refer"`
	PostgresRepository	Postgres `gorm:"-"`
}


//TableName function define the name of the table
//It return a string
func (u *User)TableName() string {
	return "public.user"
}


func (u *User)FindByLogin(login string)(*gorm.DB,*User){
	//groups:=Groups{}
	u.PostgresRepository.Initialization()	
	result:=u.PostgresRepository.Database.Preload("FK_groups").Where("\"public\".user.user=?",login).Find(&u) 
	return result,u
}


func (u *User)Init(user User){
	u.User=user.User
	u.Password=user.Password
	u.FK_groups=user.FK_groups

}
//DBMigrate function creat in DB the User Entity table
//It return a gorm.DB struct
func (u *User)DBMigrate(db *gorm.DB) (*gorm.DB) {
	db.AutoMigrate(&User{})
	return db
}

