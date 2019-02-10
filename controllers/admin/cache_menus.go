package admin

import (
	"sort"
	"sync"

	"ginserver/models"
)

var cacheRoleMenu sync.Map

func GetCacheRoleMenu(roleId int) (menus []*models.SRoleMenus, err error) {
	if record, ok := cacheRoleMenu.Load(roleId); ok {
		if menus, ok = record.([]*models.SRoleMenus); ok {
			return
		}
	}
	if menus, err = genMenusByRoleId(roleId); err != nil {
		return
	}
	cacheRoleMenu.Store(roleId, menus)
	return
}

func SetCacheRoleMenu(roleId int) (err error) {
	var (
		menus interface{}
		ok    bool
	)
	if menus, ok = cacheRoleMenu.Load(roleId); !ok {
		return
	}
	if menus, err = genMenusByRoleId(roleId); err != nil {
		return
	}
	cacheRoleMenu.Store(roleId, menus)
	return
}

func genMenusByRoleId(roleId int) (menu []*models.SRoleMenus, err error) {
	var records []*models.SRoleMenus
	if records, err = new(models.SRoleMenus).SelectMainMenusByRoleId(roleId); err != nil {
		return
	}

	var (
		data = make(map[int]map[int]*models.SRoleMenus)
		ok   bool
	)
	for _, record := range records {
		if _, ok = data[record.ParentId]; !ok {
			data[record.ParentId] = make(map[int]*models.SRoleMenus)
		}
		data[record.ParentId][record.Id] = record
	}

	menu = buildMenuTree(0, data)
	return
}

func buildMenuTree(id int, data map[int]map[int]*models.SRoleMenus) []*models.SRoleMenus {
	list := make([]*models.SRoleMenus, 0)
	for index, record := range data[id] {
		if data[index] != nil {
			record.List = buildMenuTree(index, data)
		}
		list = append(list, record)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].OrderNo < list[j].OrderNo
	})
	return list
}
