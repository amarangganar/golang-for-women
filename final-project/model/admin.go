package model

type UserSignIn struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required,alphanum,min=8"`
}

type UserRegister struct {
	Name string `form:"name" json:"name" binding:"required"`
	UserSignIn
}
