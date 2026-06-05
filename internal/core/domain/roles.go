package domain

type Role string

var (
	RoleAdmin  Role = "admin"
	RoleSeller Role = "seller"
	RoleBuyer  Role = "buyer"
)
