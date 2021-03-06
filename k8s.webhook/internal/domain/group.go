package domain

import (
	"gorm.io/gorm"
)
type Group struct {
	gorm.Model
	GID         int    `gorm:"unique"`
	Group       string `gorm:"primaryKey;unique"`
	Description string
	ID uint
	Users   []User `gorm:"many2many:user_group;"`
	Roles   []Role `gorm:"many2many:role_group;"`
}

func (g Group) DTO(gid int, group string, description string) Group{
	return Group{GID: gid, Group: group, Description: description}
}

func (g *Group) AssociationCreate(db *gorm.DB)(* gorm.DB){
	//return db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&u)
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Model(&g).Updates(&g)

}

func (g *Group)  AssociationDelete(db *gorm.DB,u User)(error){
	//return db.Session(&gorm.Session{FullSaveAssociations: true}).Delete(&g)
	return db.Model(&g).Association("Users").Delete(u)
}


func (g *Group) AssociationFindByName(db *gorm.DB,u *User)(error){
	return db.Model(&g).Debug().Where("group_group=?",g.Group).Association("Users").Find(&u)
}
func (g *Group) FindAllMember(db *gorm.DB,u *[]User)(error){
	//groupList := []Group{}
	return db.Model(&g).Debug().Where("group_group=?",g.Group).Association("Users").Find(&u)
}
//Save is method to save new Group
func (g *Group) Save(db *gorm.DB) *gorm.DB {
	return db.Create(&g)

}

func (g *Group) AssociationRoleGroupByName(db *gorm.DB)(*gorm.DB){
	return db.Debug().Preload("Roles").Where("\"group\"=?",g.Group).Find(&g)
}
func (g *Group) AssociationRoleFindByName(db *gorm.DB,u *User)(error){
	return db.Model(&g).Debug().Where("group_group=?",g.Group).Association("Users").Find(&u)
}
//FindAll is method to find all groups
func (g *Group) FindAll(db *gorm.DB) (*gorm.DB, []Group) {
	groupList := []Group{}
	result := db.Debug().Find(&groupList)
	return result, groupList
}

//FindByName is method to retrieve an Group by his name
func (g *Group) FindByName(db *gorm.DB) *gorm.DB {
	return db.Debug().Preload("Users").Where("\"group\"=?",g.Group).Find(&g)
}

//Modify is method to update an Group
func (g *Group) Modify(db *gorm.DB) *gorm.DB {
	return db.Debug().Where("\"g_id\"=?", g.GID).Updates(&g)
}

//Delete is method to remove an Group
func (g *Group) Delete(db *gorm.DB) *gorm.DB {
	return db.Debug().Where("\"group\"=?", g.Group).Delete(&g)
}

func (g *Group) DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Group{})
	return db
}