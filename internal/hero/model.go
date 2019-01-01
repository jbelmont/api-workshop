package hero

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Hero holds value for /v1/heroes endpoint
type Hero struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"Name" json:"Name"`
	SuperPowers []string      `bson:"superpowers" json:"superpowers"`
	Gender      string        `bson:"gender" json:"gender"`
	Metadata    `bson:",inline"`
}

// CreateHero holds values for creating users.
type CreateHero struct {
	Name        string   `json:"name"`
	SuperPowers []string `json:"superpowers"`
	Gender      string   `json:"gender"`
	Metadata    `json:",inline"`
}

// ListResults holds a list of heroes and pagination information.
type ListResults struct {
	Heroes        []Hero `json:"items"`
	Count         int    `json:"count"`
	Total         int    `json:"total"`
	PreviousIndex *int   `json:"previousIndex,omitempty"`
	Index         int    `json:"index"`
	NextIndex     *int   `json:"nextIndex,omitempty"`
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

// Filters represent restrictions that can be applied when listing heroes.
type Filters struct {
	SuperPowers []string `form:"superpowers" json:"superpowers" query:"superpowers"`
}
