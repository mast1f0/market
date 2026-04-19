package domain

type CartItems struct {
	Id        int64 `gorm:"primaryKey" json:"id,omitempty"`
	CartId    int64 `gorm:"not null;index" json:"cart_id"`
	ProductId int64 `gorm:"not null" json:"product_id"`
	Quantity  int   `gorm:"not null" json:"quantity"`
}
