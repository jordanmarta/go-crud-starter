package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
	Name     string `json:"name" binding:"required,min=4,max=100"`
	Age      int8   `json:"age" validate:"required,gte=1,lte=140"`
}

type UserUpdateRequest struct {
	Name string `json:"name" binding:"omitempty,min=4,max=100"`
	Age  int8   `json:"age" validate:"omitempty,gte=1,lte=140"`
}
