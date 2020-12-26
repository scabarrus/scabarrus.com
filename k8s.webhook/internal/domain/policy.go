package domain

type Policy struct {
	ID            int `gorm:"primaryKey;unique"`
	Verbs         []string
	APIGroups     []string
	Resources     []string
	ResourceNames []string
}