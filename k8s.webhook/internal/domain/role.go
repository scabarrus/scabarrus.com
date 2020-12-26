package domain
type Role struct {
	ID int `gorm:"primaryKey;unique"`
	Name string `gorm:"unique"`
	Namespace string
	Policies []Policy
}