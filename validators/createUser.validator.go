package validators

type CreateUser struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
