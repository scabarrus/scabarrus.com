package domain

import "gorm.io/gorm"
type Role struct {
	Role string `gorm:"primaryKey;unique"`
	Namespace string
	Verb string
	Group string
	Resource string
	Groups   []Group `gorm:"many2many:role_group;"`
}

//Factory method
func (r Role)DTO(role string,namespace string, verb string,group string, resource string)(Role){
	return Role{
		Role:      role,
		Namespace: namespace,
		Verb:      verb,
		Group:     group,
		Resource:  resource,
		Groups:    []Group{},
	}
}

//Save is method to save new User
func (r *Role) Save(db *gorm.DB)(* gorm.DB){
	return db.Debug().Create(&r)
		
}

func (r *Role) AssociationCreate(db *gorm.DB)(* gorm.DB){
	//return db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&u)
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Model(&r).Updates(&r)

}

func (r *Role) AssociationDelete(db *gorm.DB,g Group)(error){
	//return db.Session(&gorm.Session{FullSaveAssociations: true}).Model(&r).Delete(&r)
	return db.Model(&r).Association("Groups").Delete(g)
}

func (r *Role) AssociationFindByName(db *gorm.DB)(* gorm.DB){
	return db.Session(&gorm.Session{FullSaveAssociations: true}).Model(&r).Find(&r)
}
//FindAll is method to find all users
func (r *Role) FindAll(db *gorm.DB)(* gorm.DB,[]Role){
	roleList:=[]Role{}
	result:=db.Debug().Find(&roleList)
	return result,roleList
}

//FindByName is method to retrieve an User by his name
func (r *Role) FindByName(db *gorm.DB)(* gorm.DB){
	return db.Debug().Preload("Groups").Where("\"role\"=?",r.Role).First(&r)
}

//Modify is method to update an User
func (r *Role) Modify(db *gorm.DB)(* gorm.DB){
	return db.Debug().Where("\"role\"=?",r.Role).Updates(&r)
}

//Delete is method to remove an User
func (r *Role) Delete(db *gorm.DB)(* gorm.DB){
	return db.Debug().Where("\"role\"=?",r.Role).Delete(&r)
}

func (u *Role)DBMigrate(db *gorm.DB)(* gorm.DB){
	db.AutoMigrate(&Role{})
	return db
}