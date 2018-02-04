package hero

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Hero holds value for /v1/heroes endpoint
type Hero struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `bson:"Name" json:"Name"`
	SuperPower string        `bson:"superPower" json:"superPower"`
	Gender     string        `bson:"gender" json:"gender"`
	Metadata   `bson:",inline"`
}

// CreateHero holds values for creating users.
type CreateHero struct {
	Name       string `json:"name"`
	SuperPower string `json:"superPower"`
	Gender     string `json:"gender"`
	Metadata   `json:",inline"`
}

// Metadata fields for a hero
type Metadata struct {
	CreatedAt   time.Time  `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	CreatedByID *string    `bson:"createdById,omitempty" json:"createdById,omitempty"`
	UpdatedAt   time.Time  `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
	UpdatedByID *string    `bson:"updatedById,omitempty" json:"updatedById,omitempty"`
	RemovedAt   *time.Time `bson:"removedAt,omitempty" json:"removedAt,omitempty"`
	RemovedByID *string    `bson:"removedById,omitempty" json:"removedById,omitempty"`
}
