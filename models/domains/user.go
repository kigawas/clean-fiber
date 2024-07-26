package domains

type User struct {
	ID       int    `gorm:"id"`
	Username string `gorm:"column:username;index:user_username_key,unique"`
	//BUG: https://github.com/go-gorm/gorm/issues/7100
}

func (User) TableName() string {
	return "user"
}
