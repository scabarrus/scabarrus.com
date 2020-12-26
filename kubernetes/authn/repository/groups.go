package repository

import (
	"gorm.io/gorm"
)
type Groups struct {
	//gorm.Model
	Group string `gorm:"primaryKey;unique"`
	Description string
	Namespace string
	//FK_users []User `gorm:"many2many:groups_user;"`
}

func (g *Groups)Init(group Groups){
	g.Group=group.Group
	g.Description=group.Description
	g.Namespace=group.Namespace
}
//DBMigrate function creat in DB the User Entity table
//It return a gorm.DB struct
func (g *Groups)DBMigrate(db *gorm.DB) (*gorm.DB) {
	db.AutoMigrate(&Groups{})
	return db
}
