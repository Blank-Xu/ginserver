package admin

import (
	"errors"
	"sync"

	"ginserver/models"
)

var cacheUsers sync.Map

func GetCacheUser(userId int) (user *models.SUser, err error) {
	if record, ok := cacheUsers.Load(userId); ok {
		if user, ok = record.(*models.SUser); ok {
			return
		}
	}
	return SetCacheUserById(userId)
}

func SetCacheUserById(userId int) (user *models.SUser, err error) {
	user = models.NewSUser(userId)
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

func SetCacheUser(user *models.SUser) {
	cacheUsers.Store(user.Id, user)
}
