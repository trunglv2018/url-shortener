package db

import (
	"time"

	"github.com/trunglen/g/x/math"
)

const (
	defaultTableIDLength = 20
)

type IModel interface {
	BeforeCreate(prefix string)
	BeforeUpdate()
	BeforeDelete()
}
type BaseModel struct {
	Key       string `json:"_key"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
}

func (b *BaseModel) BeforeCreate(prefix string) {
	b.Key = math.RandString(prefix, defaultTableIDLength)
	b.CreatedAt = time.Now().Unix()
	b.UpdatedAt = time.Now().Unix()
}

func (b *BaseModel) BeforeUpdate() {
	b.UpdatedAt = time.Now().Unix()
}

func (b *BaseModel) BeforeDelete() {
	b.DeletedAt = time.Now().Unix()
}

func (m *BaseModel) SetID(id string) {
	m.Key = id
}
