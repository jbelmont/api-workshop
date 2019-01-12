package hero

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Hero holds value for /v1/heroes endpoint
type Hero struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	SuperPowers []string      `bson:"superpowers" json:"superpowers"`
	Gender      string        `bson:"gender" json:"gender"`
	Metadata    `bson:",inline"`
}

// CreateHero holds values for creating users.
type CreateHero struct {
	Name        string   `bson:"name" json:"name"`
	SuperPowers []string `bson:"superpowers" json:"superpowers"`
	Gender      string   `bson:"gender" json:"gender"`
	Metadata    `bson:",inline"`
}

// UpdateHero is what we accept from users to update a Hero record.
type UpdateHero CreateHero

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
	Created      time.Time `bson:"created" json:"created"`
	LastModified time.Time `bson:"lastModified" json:"lastModified"`
	IsRemoved    bool      `bson:"isRemoved" json:"-"`
}

// Filters represent restrictions that can be applied when listing heroes.
type Filters struct {
	SuperPowers []string `form:"superpowers" json:"superpowers" query:"superpowers"`
}
