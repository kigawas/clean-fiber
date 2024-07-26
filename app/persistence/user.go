package persistence

import (
	"github.com/kigawas/clean-fiber/models/domains"
	"github.com/kigawas/clean-fiber/models/params"
	"github.com/kigawas/clean-fiber/models/queries"
	"github.com/kigawas/clean-fiber/models/schemas"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB, query *queries.UserQuery) (schemas.UserListSchema, error) {
	var users []domains.User
	result := db.Where("username LIKE ?", "%"+query.Username+"%").Find(&users)

	if result.Error != nil {
		return schemas.UserListSchema{}, result.Error
	}
	return schemas.FromUsers(users), nil
}

func GetUser(db *gorm.DB, id int) (schemas.UserSchema, error) {
	var user domains.User
	result := db.First(&user, id)

	if result.Error != nil {
		return schemas.UserSchema{}, result.Error
	}
	return schemas.FromUser(&user), nil
}

func CreateUser(db *gorm.DB, p *params.CreateUserParams) (schemas.UserSchema, error) {
	user := domains.User{
		Username: p.Username,
	}
	result := db.Create(&user)

	if result.Error != nil {
		return schemas.UserSchema{}, result.Error
	}
	return schemas.FromUser(&user), nil
}
