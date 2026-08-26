package models

type User struct {
	ID       int    `json:"id"`
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
