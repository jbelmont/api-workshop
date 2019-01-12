package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	jsonpatch "github.com/evanphx/json-patch"

	"github.com/jbelmont/api-workshop/internal/hero"
	database "github.com/jbelmont/api-workshop/internal/platform/db"
	"github.com/jbelmont/api-workshop/internal/platform/web"
)

// Hero holds database information for a given hero
type Hero struct {
	MasterDB *database.DB
}

const (
	defaultPageSize = 50
	maxPageSize     = 100
)

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
		return errors.Wrapf(err, "Hero: %+v", &hero.Hero{})
	}

	web.Respond(ctx, w, heroes, http.StatusCreated)
	return nil
}

// List returns a paginated list of heroes
func (h *Hero) List(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	dbConn, err := h.MasterDB.Copy()
	if err != nil {
		return errors.Wrapf(web.ErrDBNotConfigured, "")
	}
	defer dbConn.Close()

	// Default page size can be pulled from configuration.
	paging := web.PopulatePaging(*r.URL, defaultPageSize, maxPageSize)

	filters, err := h.populateFilters(r)
	if err != nil {
		return errors.Wrap(err, "something happened while populating filters")
	}

	heroes, err := hero.List(ctx, dbConn, *filters, paging)
	if err != nil {
		return errors.Wrap(err, "something happened while listing heroes")
	}

	web.Respond(ctx, w, heroes, http.StatusOK)
	return nil
}

func (h *Hero) populateFilters(r *http.Request) (*hero.Filters, error) {
	qv := r.URL.Query()

	filter := hero.Filters{
		SuperPowers: qv["superpowers"],
	}
	return &filter, nil
}

// Retrieve returns a hero from the heroes collection.
func (h *Hero) Retrieve(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	dbConn, err := h.MasterDB.Copy()
	if err != nil {
		return errors.Wrapf(web.ErrDBNotConfigured, "")
	}
	defer dbConn.Close()

	hero, err := hero.Retrieve(ctx, dbConn, params["id"])
	if err != nil {
		return errors.Wrap(err, "calling hero retrieve service")
	}

	web.Respond(ctx, w, hero, http.StatusOK)
	return nil
}

// Delete deletes a hero from the heroes collection.
func (h *Hero) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	dbConn, err := h.MasterDB.Copy()
	if err != nil {
		return errors.Wrapf(web.ErrDBNotConfigured, "")
	}
	defer dbConn.Close()

	if err := hero.Delete(ctx, dbConn, params["id"]); err != nil {
		return errors.Wrap(err, "calling hero delete service")
	}

	web.Respond(ctx, w, nil, http.StatusNoContent)
	return nil
}

// Update accepts PATCH requests in application/merge-patch+json form and modifies a hero collection information
func (h *Hero) Update(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	// Force the endpoint to provide application/merge-patch+json Content-Type
	if typ, want := r.Header.Get("Content-Type"), "application/merge-patch+json"; typ != want {
		return web.AppError{
			Message: fmt.Sprintf("Unsupported content type: %q. Must be %q.", typ, want),
			Status:  http.StatusUnsupportedMediaType,
			Code:    "CONTENT_TYPE",
		}
	}
	dbConn, err := h.MasterDB.Copy()
	if err != nil {
		return errors.Wrapf(web.ErrDBNotConfigured, "")
	}
	defer dbConn.Close()

	id := params["id"]

	retrievedHero, err := hero.Retrieve(ctx, dbConn, id)
	if err != nil {
		return errors.Wrap(err, "calling hero retrieve service")
	}

	// read request body to get updated object's JSON
	uHero, err := ioutil.ReadAll(r.Body)

	var upd hero.UpdateHero

	// now merge the mergePatch into the originalJSON
	originalHero, err := json.Marshal(retrievedHero)
	if err != nil {
		return errors.New("Error trying to marshal a retrieved hero")
	}

	resultJSON, err := jsonpatch.MergePatch(originalHero, uHero)
	if err != nil {
		return errors.Wrap(err, "something happened while merging structs together")
	}

	// Unmarshal the result JSON into the upd struct passed in
	err = json.Unmarshal(resultJSON, upd)

	if err = hero.Update(ctx, dbConn, upd, id); err != nil {
		return errors.Wrap(err, "calling user update service")
	}

	web.Respond(ctx, w, retrievedHero, http.StatusOK)

	return nil
}
