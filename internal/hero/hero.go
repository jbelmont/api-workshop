package hero

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	apiContext "github.com/jbelmont/api-workshop/internal/platform/context"
	database "github.com/jbelmont/api-workshop/internal/platform/db"
)

const heroCollection = "heroes"

// Create creates new hero in DB
func Create(ctx context.Context, dbConn *database.DB, cH *CreateHero) (*Hero, error) {
	ctxValues := ctx.Value(apiContext.KeyValues).(*apiContext.Values)
	createdByID := ctxValues.ID
	h := Hero{
		ID:         bson.NewObjectId(),
		Name:       cH.Name,
		SuperPower: cH.SuperPower,
		Gender:     cH.Gender,
		Metadata: Metadata{
			CreatedAt:   time.Now(),
			CreatedByID: &createdByID,
			UpdatedAt:   time.Now(),
		},
	}

	f := func(collection *mgo.Collection) error {
		return collection.Insert(h)
	}

	if err := dbConn.Execute(ctx, heroCollection, f); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("db.%s.insert(%s)", heroCollection, database.Query(h)))
	}

	return &h, nil
}
