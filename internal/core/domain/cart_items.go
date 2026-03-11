package domain

type CartItems struct {
	Id          int `gorm:"primary_key"`
	CartId      int `gorm:"not null"`
	ProductId   int `gorm:"not null"`
	Quantity    int `gorm:"not null"`
	FkCartId    int `gorm:"not null"`
	FkProductId int `gorm:"not null"`
}
