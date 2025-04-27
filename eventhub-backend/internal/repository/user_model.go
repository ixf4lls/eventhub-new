package repository

type UserModel struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `gorm:"unique" json:"username"`
	PasswordHash string `json:"password_hash"`
	// TelegramUsername string `json:"telegram_username"`
	// PhoneNumber      string `json:"phone_number"`
}

func (UserModel) TableName() string {
	return "users"
}

type UserAsMember struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}
