package models

type Auth struct {
	ID     int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	APIKey string `gorm:"column:api-key;type:varchar(32);not null" json:"api_key"`
}

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"type:varchar(64);not null" json:"username"`
}

type UserProfile struct {
	UserID    int64  `gorm:"primaryKey" json:"user_id"`
	FirstName string `gorm:"type:varchar(32);not null" json:"first_name"`
	LastName  string `gorm:"type:varchar(64);not null" json:"last_name"`
	Phone     string `gorm:"type:varchar(64);not null" json:"phone"`
	Address   string `gorm:"type:varchar(64);not null" json:"address"`
	City      string `gorm:"type:varchar(64);not null" json:"city"`
}

type UserData struct {
	UserID int64  `gorm:"primaryKey" json:"user_id"`
	School string `gorm:"type:varchar(32);not null" json:"school"`
}

type GetUserParams struct {
	Username *string `json:"username,omitempty"`
}

type Profile struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	City      string `json:"city"`
	School    string `json:"school"`
}
