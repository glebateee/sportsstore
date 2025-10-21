package auth

import (
	"platform/authorization/identity"
	"platform/services"
	"strings"
)

var users = map[int]identity.User{
	1: identity.NewBasicUser(1, "Alice", "Administrator"),
}

type userStore struct{}

// GetUserByID implements identity.UserStore.
func (u *userStore) GetUserByID(id int) (user identity.User, found bool) {
	user, found = users[id]
	return
}

// GetUserByName implements identity.UserStore.
func (u *userStore) GetUserByName(name string) (user identity.User, found bool) {
	for _, user = range users {
		if strings.EqualFold(user.GetDisplayName(), name) {
			return user, true
		}
	}
	return nil, false
}

func RegisterUserStoreService() {
	err := services.AddSingleton(func() identity.UserStore {
		return &userStore{}
	})
	if err != nil {
		panic(err)
	}
}
