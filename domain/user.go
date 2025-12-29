package domain

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    string
	Phone    string
	Password string
	Role     string
}
