package models

type User struct {
	ID       int    `json:"id" binding:"required"`
	Name     string `json:"name" `
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}
type LoginReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Token struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
}

var UsersData = []User{
	{
		ID:       1,
		Name:     "kenzy ",
		Email:    "kenzy12@gmail.com",
		Password: "kekoo12",
		Role:     "admin",
	},
	{
		ID:       2,
		Name:     "ali ",
		Email:    "ali12@gmail.com",
		Password: "ali12",
		Role:     "user",
	},
	{
		ID:       3,
		Name:     "george",
		Email:    "george@gmail.com",
		Password: "george12",
		Role:     "admin",
	},
}
