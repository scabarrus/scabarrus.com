package repository

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
type Postgres struct {
	Database *gorm.DB
}
func (pg *Postgres)Initialization()(*gorm.DB,error){
	dsn := os.Getenv("POSTGRES_CON")
	var err error
	pg.Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//pg.Database.Migrate(&u)
	
	return pg.Database,err 
}

func (pg *Postgres)InitDB(i interface{}){
	pg.Database.Create(i)
}



