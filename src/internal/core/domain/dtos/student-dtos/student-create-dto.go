package dtos

type StudentCreateDTO struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Connection string `json:"connection"`
}
