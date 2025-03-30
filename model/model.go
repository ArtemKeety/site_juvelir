package model

import "time"

type Category struct {
	Id   uint   `gorm:"primary_key"`
	Name string `gorm:"column:name"`
}

type Product struct {
	Id        uint
	Name      string
	Price     uint
	PhotoPath string
	Count     int
}

type TovarDetail struct {
	Id         uint    `gorm:"primaryKey"`
	Name       string  `gorm:"column:name"`
	Price      uint    `gorm:"column:price"`
	Counts     uint8   `gorm:"column:counts"`
	IdCategory uint    `gorm:"column:idCategory"`
	Photos     []Photo `gorm:"foreignKey:TovarId"` // Убедитесь, что внешний ключ для связи с Photos называется TovarId
}

func (TovarDetail) TableName() string {
	return "tovar"
}

type Photo struct {
	Id       uint   `gorm:"primaryKey"`
	Filepath string `gorm:"column:filepath"`
	TovarId  uint   `gorm:"column:idTovar"` // Убедитесь, что в БД внешний ключ называется idTovar (не TovarId)
}

func (Photo) TableName() string {
	return "photo"
}

type Basket struct {
	Id      uint    `gorm:"primary_key"`
	UserId  uint    `gorm:"foreignKey:UserId"`  // внешний ключ для user
	TovarId uint    `gorm:"foreignKey:TovarId"` // внешний ключ для
	Counts  int     `gorm:"column:counts"`
	Tovar   Product `gorm:"foreignKey:TovarId"` // связь с товаром
}

type UserModel struct {
	Id        uint   `gorm:"primary_key"`
	Login     string `gorm:"column:login"`
	Password  string `gorm:"column:password"`
	FirstName string `gorm:"column:first_name"`
	NastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
}

type Session struct {
	Id       uint      `gorm:"primaryKey"`
	UserId   uint      `gorm:"foreignKey"`       // Внешний ключ на user
	Token    string    `gorm:"column:token"`     // Токен с ограничением длины
	DateTime time.Time `gorm:"column:date_time"` // Время создания сессии
}

type PageData struct {
	User   *UserModel // Информация о пользователе (nil, если не авторизован)
	Tovars []Product  // Список товаров
}

type PageDataForOneProduct struct {
	User  *UserModel  // Информация о пользователе (nil, если не авторизован)
	Tovar TovarDetail // Список товаров
}
