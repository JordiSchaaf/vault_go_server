package models

type User struct {
	Id              int64  `gorm:"primary_key" json:"id"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	PermissionLevel int    `json:"permissionLevel"`
}

func (b *User) TableName() string {
	return "user"
}
