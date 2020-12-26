package service

import (
	"net/http"

	"gorm.io/gorm"
	"scabarrus.com/k8s.webhook/internal/domain"
	"scabarrus.com/k8s.webhook/internal/repository"
)


type DBMigrate struct {
	name string
}

func (dbm *DBMigrate) MigrateRepo(db *gorm.DB) {
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Group{})
}


func (dbm *DBMigrate) MigrateService(w http.ResponseWriter, r *http.Request){
	pg := repository.Postgres{}
	pg.Initialization()
	dbm.MigrateRepo(pg.Database)
}