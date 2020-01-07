package role

import (
	"ginserver/pkg/db"
	"time"
)

type Role struct {
	*db.Model `xorm:"-" json:"-"`
	Id        int    `xorm:"pk autoincr" json:"id"`
	Name      string `xorm:"unique"`
	State     bool
	Remark    string
	Created   time.Time `xorm:"created"`
	Updater   int
	Updated   time.Time `xorm:"updated"`
}

func (p *Role) TableName() string {
	return "s_role"
}

func (p *Role) SelectOneByUserId(userId int) (bool, error) {
	return db.GetDefaultDB().SQL(`
SELECT role.*
FROM s_user_role user_role
LEFT JOIN s_role role ON role.id = user_role.role_id
WHERE user_role.user_id = ?
`, userId).Get(p)
}
