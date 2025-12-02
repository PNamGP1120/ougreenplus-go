package user

type CreateUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

type UpdateUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}
