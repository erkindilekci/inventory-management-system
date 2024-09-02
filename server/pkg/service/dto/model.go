package dto

type UserCreate struct {
	Username string
	Password string
	Role     string
}

type ProductCreate struct {
	Name     string
	Price    float32
	Quantity int64
	Category string
}
