package user

import "api/app/models"

type User struct {
	models.BaseModel

	Name     string `gorm:"type:varchar(100);not null;comment:用户名;" json:"name,omitempty"`
	Email    string `gorm:"type:varchar(100);comment:邮箱;" json:"-"` // 因为我们不希望将敏感信息输出给用户
	Phone    string `gorm:"type:varchar(100);comment:手机号码;" json:"-"`
	Password string `gorm:"type:varchar(256);comment:密码;" json:"-"`

	models.CommonTimestampsField
}
