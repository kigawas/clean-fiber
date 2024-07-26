package schemas

import "github.com/kigawas/clean-fiber/models/domains"

type UserSchema struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type UserListSchema struct {
	Users []UserSchema `json:"users"`
}

func FromUser(user *domains.User) UserSchema {
	return UserSchema{
		ID:       user.ID,
		Username: user.Username,
	}
}

func FromUsers(users []domains.User) UserListSchema {
	userSchemas := []UserSchema{}
	for _, user := range users {
		userSchemas = append(userSchemas, FromUser(&user))
	}
	return UserListSchema{Users: userSchemas}
}
