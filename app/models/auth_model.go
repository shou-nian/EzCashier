package models

type LoginRequest struct {
	PhoneNum string `json:"phone_num" binding:"required,len=11"`
	Password string `json:"password" binding:"required,min=8,max=18"`
}
