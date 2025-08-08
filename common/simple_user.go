package common

type SimpleUser struct {
	SQLmodel  `json:",inline"`
	LastName  string `gorm:"column:last_name" json:"last_name"`
	FirstName string `gorm:"column:first_name" json:"first_name"`
	Role      string `gorm:"column:role" json:"role"`
	Avatar    *Image `gorm:"colume:avatar;type:json" json:"avatar,omitempty"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func (u *SimpleUser) Mask(isAdmin bool) {
	u.GenUID(DbTypeUser)
}
