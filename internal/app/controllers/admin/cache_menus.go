package admin

import (
	"sort"
	"sync"

	"ginserver/internal/app/models/s_role"
)

var cacheRoleMenu sync.Map

func GetCacheRoleMenu(roleId int) (menus []*s_role.RoleMenuDetail, err error) {
	if record, ok := cacheRoleMenu.Load(roleId); ok {
		if menus, ok = record.([]*s_role.RoleMenuDetail); ok {
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

func genMenusByRoleId(roleId int) ([]*s_role.RoleMenuDetail, error) {
	var (
		record  s_role.RoleMenuDetail
		records []*s_role.RoleMenuDetail
		err     error
	)
	if records, err = record.SelectMainMenuByRoleId(roleId); err != nil {
		return nil, err
	}

	var (
		data = make(map[int]map[int]*s_role.RoleMenuDetail)
		ok   bool
	)
	for _, record := range records {
		if _, ok = data[record.ParentId]; !ok {
			data[record.ParentId] = make(map[int]*s_role.RoleMenuDetail)
		}
		data[record.ParentId][record.Id] = record
	}

	return buildMenuTree(0, data), nil
}

// buildMenuTree to build main menu tree
func buildMenuTree(id int, data map[int]map[int]*s_role.RoleMenuDetail) []*s_role.RoleMenuDetail {
	list := make([]*s_role.RoleMenuDetail, 0, len(data))
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
