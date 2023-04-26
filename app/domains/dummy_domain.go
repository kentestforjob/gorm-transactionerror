package domains

import (
	"test/gormtransactionerr/utils"

	"gorm.io/gorm"
)

type Dummy struct {
	ID        uint32          `form:"id" json:"id" gorm:"primary_key" `
	UserId    *uint32         `form:"user_id" json:"user_id" `
	Email     string          `form:"email" json:"email" gorm:"index;uniqueIndex:idx_name;type:varchar(191);not null;" validate:"omitempty,email"`
	UpdatedAt utils.JSONTime  `form:"updated_at" json:"updated_at"`                  // last sync at
	CreatedAt utils.JSONTime  `form:"created_at" json:"created_at" gorm:"<-:create"` // first sync at
	DeletedAt *gorm.DeletedAt `form:"deleted_at" json:"deleted_at" gorm:"index;uniqueIndex:idx_name;"`
}

func (Dummy) TableName() string {
	return "dummy"
}
