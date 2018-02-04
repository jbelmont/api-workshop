package handlers

import (
	"context"
	"net/http"

	"github.com/pkg/errors"

	"github.com/jbelmont/api-workshop/internal/hero"
	database "github.com/jbelmont/api-workshop/internal/platform/db"
	"github.com/jbelmont/api-workshop/internal/platform/web"
)

// Hero holds database information for a given hero
type Hero struct {
	MasterDB *database.DB
}

// Create inserts a new User in the system.
func (h *Hero) Create(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	dbConn, err := h.MasterDB.Copy()
	if err != nil {
		return errors.Wrapf(web.ErrDBNotConfigured, "")
	}
	defer dbConn.Close()

	var cHero hero.CreateHero
	if err = web.Unmarshal(r.Body, &cHero); err != nil {
		return errors.Wrap(err, "unmarshalling new hero")
	}

	heroes, err := hero.Create(ctx, dbConn, &cHero)
	if err != nil {
		return errors.Wrapf(err, "User: %+v", &hero.Hero{})
	}

	web.Respond(ctx, w, heroes, http.StatusCreated)
	return nil
}
