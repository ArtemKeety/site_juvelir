package model

import "time"

type Category struct {
	Id   uint   `gorm:"primary_key" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

type Product struct {
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Price     uint   `json:"price"`
	PhotoPath string `json:"photo_path"`
	Count     int    `json:"count"`
}

type TovarDetail struct {
	Id         uint    `gorm:"primaryKey" json:"id"`
	Name       string  `gorm:"column:name" json:"name"`
	Price      uint    `gorm:"column:price" json:"price"`
	Counts     uint8   `gorm:"column:counts" json:"counts"`
	IdCategory uint    `gorm:"column:idCategory" json:"id_category"`
	Photos     []Photo `gorm:"foreignKey:TovarId" json:"photos"`
}

func (TovarDetail) TableName() string {
	return "tovar"
}

type Photo struct {
	Id       uint   `gorm:"primaryKey" json:"id"`
	Filepath string `gorm:"column:filepath" json:"filepath"`
	TovarId  uint   `gorm:"column:idTovar" json:"tovar_id"`
}

func (Photo) TableName() string {
	return "photo"
}

type Basket struct {
	Id      uint    `gorm:"primary_key" json:"id"`
	UserId  uint    `gorm:"foreignKey:UserId" json:"user_id"`   // внешний ключ для user
	TovarId uint    `gorm:"foreignKey:TovarId" json:"tovar_id"` // внешний ключ для товара
	Counts  int     `gorm:"column:counts" json:"counts"`
	Tovar   Product `gorm:"foreignKey:TovarId" json:"tovar"`
}

type UserModel struct {
	Id        uint   `gorm:"primary_key" json:"id"`
	Login     string `gorm:"column:login" json:"login"`
	Password  string `gorm:"column:password" json:"password"`
	FirstName string `gorm:"column:first_name" json:"first_name"`
	NastName  string `gorm:"column:last_name" json:"last_name"`
	Email     string `gorm:"column:email" json:"email"`
}

type Session struct {
	Id       uint      `gorm:"primaryKey" json:"id"`
	UserId   uint      `gorm:"foreignKey" json:"user_id"`         // Внешний ключ на user
	Token    string    `gorm:"column:token" json:"token"`         // Токен с ограничением длины
	DateTime time.Time `gorm:"column:date_time" json:"date_time"` // Время создания сессии
}

type PageData struct {
	User   *UserModel // Информация о пользователе (nil, если не авторизован)
	Tovars []Product  // Список товаров
}

type PageDataForOneProduct struct {
	User  *UserModel  // Информация о пользователе (nil, если не авторизован)
	Tovar TovarDetail // Список товаров
}
