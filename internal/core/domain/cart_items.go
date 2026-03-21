package domain

type CartItems struct {
	Id        int64 `gorm:"primary_key, omitempty" json:"id"`
	CartId    int64 `gorm:"not null" json:"cart_id"`
	ProductId int64 `gorm:"not null" json:"product_id"`
	Quantity  int   `gorm:"not null" json:"quantity"`
}
