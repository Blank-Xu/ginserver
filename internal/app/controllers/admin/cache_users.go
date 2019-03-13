package admin

import (
	"errors"
	"sync"

	"ginserver/internal/app/models/s_user"
)

var cacheUsers sync.Map

func GetCacheUser(userId int) (s_user.User, error) {
	var user s_user.User
	if record, ok := cacheUsers.Load(userId); ok {
		if user, ok = record.(s_user.User); ok {
			return user, nil
		}
	}
	return SetCacheUserById(userId)
}

func SetCacheUserById(userId int) (s_user.User, error) {
	var (
		user = s_user.User{Id: userId}
		has  bool
	)
	has, err := user.SelectOne(&user)
	if err != nil {
		return user, err
	}
	if !has {
		return user, errors.New("invalid params")
	}
	SetCacheUser(user)
	return user, nil
}

func SetCacheUser(user s_user.User) {
	cacheUsers.Store(user.Id, user)
}
