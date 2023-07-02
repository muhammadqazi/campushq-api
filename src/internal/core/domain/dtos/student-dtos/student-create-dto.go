package dtos

type StudentCreateDTO struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Connection string `json:"connection"`
}
