package user

import "api/app/models"

type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"` // 因为我们不希望将敏感信息输出给用户
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
