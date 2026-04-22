package dto

type Registraton_request struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=4"`
	Role_id  int    `json:"role_id" binding:"required"`
}

type Login_request struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=4"`
}
