package domain

type Cart struct {
	Id     int64 `gorm:"primary_key" json:"id"`
	UserId int64 `json:"user_id"`
}
