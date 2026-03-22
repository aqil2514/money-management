package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description" validate:"required"`
	Type        string         `json:"type" validate:"required,oneof=expense income transfer payable receivable"`
	ParentID    *uuid.UUID     `gorm:"type:uuid;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"parentId"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
}

func (t *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		t.ID = uuid.New()
	}

	return
}
