package samplemodel

import (
	"time"

	"github.com/google/uuid"
)

type SampleDataModel struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type SampleDataCreateModel struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}

type SampleDataUpdateModel struct {
	ID        uuid.UUID `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Sample struct {
	ID        uuid.UUID  `json:"id"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	CreatedAt time.Time  `json:"created_at"`
}

func (Sample) TableName() string {
	return "sample_datas"
}
