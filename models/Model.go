package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// https://gorm.io/docs/models.html
type ModelDefault struct {
	ID          uint `json:"id" gorm:"primary_key, autoIncrement"` //Model id pk
	ModelAtTime      // Model from struct ModelAtTime
	ModelUUID        // Model from struct ModelUUID
}

type ModelAtTime struct {
	ModelDellog
	ModelUpdatelog
}

// json:"-"
type ModelDellog struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" sql:"default:null" swaggerignore:"true" example:"2021-02-02 11:11:11"` //Model Deleted At
}
type ModelUpdatelog struct {
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" example:"2021-02-02 11:11:11"` // Model Created At
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" example:"2021-02-02 11:11:11"` // Model UPdated At
}

type ModelUUIDPK struct {
	UID string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"` //Model UUID with primarykey
}

type ModelUUID struct {
	UID string `sql:"type:uuid;default:uuid_generate_v4()"` // Model UUID
}

// https://gorm.io/docs/hooks.html
func (u *ModelUUIDPK) BeforeCreate(tx *gorm.DB) (err error) {
	u.UID = uuid.NewV4().String()
	return
}

// https://gorm.io/docs/hooks.html
func (u *ModelUUID) BeforeCreate(tx *gorm.DB) (err error) {
	u.UID = uuid.NewV4().String()
	return
}
