package models

type Contact struct {
	Phone   uint   `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type User struct {
	Id       int     `json:"id" gorm:"primarykey"`
	FullName string  `json:"full_name"`
	Login    string  `json:"login" `
	Password string  `json:"password"`
	Contacts Contact `json:"contacts" gorm:"embedded;embeddedPrefix:contacts"`
	Role     string  `json:"role"`
	IsActive bool    `json:"is_active" gorm:"not null; default: true"`
}

type SignInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
