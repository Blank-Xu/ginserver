package admin

import (
	"sort"
	"sync"

	"ginserver/internal/app/models"
)

var cacheRoleMenu sync.Map

func GetCacheRoleMenu(roleId int) (menus []*models.SRoleMenuDetail, err error) {
	if record, ok := cacheRoleMenu.Load(roleId); ok {
		if menus, ok = record.([]*models.SRoleMenuDetail); ok {
			return
		}
	}
	if menus, err = genMenusByRoleId(roleId); err != nil {
		return
	}
	cacheRoleMenu.Store(roleId, menus)
	return
}

func SetCacheRoleMenu(roleId int) error {
	menus, err := genMenusByRoleId(roleId)
	if err != nil {
		return err
	}
	cacheRoleMenu.Store(roleId, menus)
	return nil
}

func genMenusByRoleId(roleId int) (menu []*models.SRoleMenuDetail, err error) {
	var records []*models.SRoleMenuDetail
	if records, err = new(models.SRoleMenuDetail).SelectMainMenuByRoleId(roleId); err != nil {
		return
	}

	var (
		data = make(map[int]map[int]*models.SRoleMenuDetail)
		ok   bool
	)
	for _, record := range records {
		if _, ok = data[record.ParentId]; !ok {
			data[record.ParentId] = make(map[int]*models.SRoleMenuDetail)
		}
		data[record.ParentId][record.Id] = record
	}

	menu = buildMenuTree(0, data)
	return
}

// buildMenuTree to build main menu tree
func buildMenuTree(id int, data map[int]map[int]*models.SRoleMenuDetail) []*models.SRoleMenuDetail {
	list := make([]*models.SRoleMenuDetail, 0)
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
