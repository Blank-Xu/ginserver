package admin

import (
	"errors"
	"sync"

	"ginserver/models/system/user"
)

var cacheUsers sync.Map

func GetCacheUser(userId int) (user.User, error) {
	var userInfo user.User
	if record, ok := cacheUsers.Load(userId); ok {
		if userInfo, ok = record.(user.User); ok {
			return userInfo, nil
		}
	}
	return SetCacheUserById(userId)
}

func SetCacheUserById(userId int) (user.User, error) {
	userInfo := user.User{Id: userId}

	has, err := userInfo.SelectOne(&userInfo)
	if err != nil {
		return userInfo, err
	}
	if !has {
		return userInfo, errors.New("invalid params")
	}
	SetCacheUser(userInfo)
	return userInfo, nil
}

func SetCacheUser(user user.User) {
	cacheUsers.Store(user.Id, user)
}
