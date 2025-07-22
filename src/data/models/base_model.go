package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Id int `gorm:"primarykey"`

	CreatedAt  time.Time    `gorm:"type:TIMESTAMP with time zone;not null"`
	ModifiedAt sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`
	DeletedAt  sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`

	CreatedBy  int            `gorm:"not null"`
	ModifiedBy *sql.NullInt64 `gorm:"null"`
	DeletedBy  *sql.NullInt64 `gorm:"null"`
}

func (m *BaseModel) BefforCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("userId")
	userId := -1

	if value != nil {
		userId = int(value.(float64))
	}
	m.CreatedBy = userId
	m.CreatedAt = time.Now().UTC()

	return
}

func (m *BaseModel) BefforUpdate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("userId")
	userId := &sql.NullInt64{Valid: false}

	if value != nil {
		userId = &sql.NullInt64{Valid: true, Int64: int64(value.(float64))}
	}
	m.ModifiedBy = userId
	m.ModifiedAt = sql.NullTime{Valid: true, Time: time.Now().UTC()}

	return
}

func (m *BaseModel) BefforDelete(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("userId")
	userId := &sql.NullInt64{Valid: false}

	if value != nil {
		userId = &sql.NullInt64{Valid: true, Int64: int64(value.(float64))}
	}
	m.DeletedBy = userId
	m.DeletedAt = sql.NullTime{Valid: true, Time: time.Now().UTC()}

	return
}
