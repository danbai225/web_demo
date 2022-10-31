package base

import (
	"time"
)

type Base struct {
	CreatedAt *time.Time `json:"created_at;comment:创建时间"`
	UpdatedAt *time.Time `json:"updated_at;comment:更新时间"`
}
