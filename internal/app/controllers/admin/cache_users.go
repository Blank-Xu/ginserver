package admin

import (
	"errors"
	"sync"

	"ginserver/internal/app/models/s_user"
)

var cacheUsers sync.Map

func GetCacheUser(userId int) (user *s_user.User, err error) {
	if record, ok := cacheUsers.Load(userId); ok {
		if user, ok = record.(*s_user.User); ok {
			return
		}
	}
	return SetCacheUserById(userId)
}

func SetCacheUserById(userId int) (user *s_user.User, err error) {
	user = &s_user.User{Id: userId}
	var has bool
	if has, err = user.SelectOne(user); err != nil {
		return
	}
	if !has {
		return nil, errors.New("invalid params")
	}
	SetCacheUser(user)
	return
}

func SetCacheUser(user *s_user.User) {
	cacheUsers.Store(user.Id, user)
}
