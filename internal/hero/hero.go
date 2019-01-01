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
	"github.com/jbelmont/api-workshop/internal/platform/web"
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

// List returns a list of heroes from the database and applies query string parameters
func List(ctx context.Context, dbConn *database.DB, filters Filters, paging web.Paging) (*ListResults, error) {
	q, err := extractQueryFromFilters(filters)
	if err != nil {
		return nil, errors.Wrap(err, err.Error())
	}

	total := 0
	fn := func(coll *mgo.Collection) error {
		var err error
		total, err = coll.Find(q).Count()
		return err
	}

	if err := dbConn.Execute(ctx, heroCollection, fn); err != nil {
		return nil, errors.Wrap(err, "db.heroes.Count()")
	}

	heroes := []Hero{}
	if total > 0 {
		f := func(coll *mgo.Collection) error {
			q := coll.Find(q)

			// Apply sort
			if len(paging.Sort) > 0 {
				q = q.Sort(paging.Sort...)
			} else {
				// default sort by created
				q = q.Sort("created")
			}

			return q.Skip(paging.Size * paging.Index).Limit(paging.Size).All(&heroes)
		}
		if err := dbConn.Execute(ctx, heroCollection, f); err != nil {
			return nil, errors.Wrap(err, "db.heroes.Find()")
		}
	}

	listResults := ListResults{
		Heroes: heroes,
		Count:  len(heroes),
		Total:  total,
		Index:  paging.Index,
	}

	// Calculate the number of items to this page
	countToIndex := paging.Size * paging.Index
	// Calculate the number of items in next page
	nextPagesCount := total - (countToIndex + paging.Size)

	if nextPagesCount > 0 {
		listResults.NextIndex = new(int)
		*listResults.NextIndex = paging.Index + 1
	}
	if paging.Index > 1 {
		listResults.PreviousIndex = new(int)
		*listResults.PreviousIndex = paging.Index - 1
	}
	return &listResults, nil
}

func extractQueryFromFilters(filters Filters) (bson.M, error) {
	query := bson.M{}

	// Superpowers
	if len(filters.SuperPowers) > 0 {
		superpowers := make([]string, 0)
		for _, powers := range filters.SuperPowers {
			superpowers = append(superpowers, powers)
		}
		query["superpowers"] = bson.M{"$in": filters.SuperPowers}
	}

	return query, nil
}
