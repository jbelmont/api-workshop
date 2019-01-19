package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jbelmont/api-workshop/cmd/apid/config"
	"github.com/jbelmont/api-workshop/internal/hero"
	database "github.com/jbelmont/api-workshop/internal/platform/db"
	"github.com/jbelmont/api-workshop/internal/platform/web"
)

// Hero holds database information for a given hero
type Hero struct {
	config   config.Config
	MasterDB *database.DB
}

// [go] cannot use h.List (type func(http.ResponseWriter, *http.Request, map[string]string) error) as type httptreemux.HandlerFunc in argument to api.GET

const (
	defaultPageSize = 50
	maxPageSize     = 100
)

// Create inserts a new User in the system.
func (h *Hero) Create(w http.ResponseWriter, r *http.Request, params map[string]string) {
	dbConn, err := h.MasterDB.Copy()
	if err != nil {
		log.Print("DB not initialized")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	var cHero hero.CreateHero

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Could not read body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(body, &cHero); err != nil {
		log.Printf(err.Error())
		return
	}

	heroes, err := hero.Create(dbConn, &cHero)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	// Write the response.
	response, err := json.Marshal(&heroes)
	if err != nil {
		log.Printf("ERROR: Failed while marshalling response: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)

	return
}

// List returns a paginated list of heroes
func (h *Hero) List(w http.ResponseWriter, r *http.Request, params map[string]string) {
	dbConn, err := h.MasterDB.Copy()
	if err != nil {
		log.Print("ERROR: Trying to initialize database", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	// Default page size can be pulled from configuration.
	paging := web.PopulatePaging(*r.URL, defaultPageSize, maxPageSize)

	filters, err := h.populateFilters(r)
	if err != nil {
		log.Printf("ERROR: Failed while parsing filters: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	heroes, err := hero.List(dbConn, *filters, paging)
	if err != nil {
		log.Printf("ERROR: Trying to list heroes: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write the response.
	response, err := json.Marshal(&heroes)
	if err != nil {
		log.Printf("ERROR: Failed while marshalling response: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}

func (h *Hero) populateFilters(r *http.Request) (*hero.Filters, error) {
	qv := r.URL.Query()

	filter := hero.Filters{
		SuperPowers: qv["superpowers"],
	}
	return &filter, nil
}

// Retrieve returns a hero from the heroes collection.
func (h *Hero) Retrieve(w http.ResponseWriter, r *http.Request, params map[string]string) {
	dbConn, err := h.MasterDB.Copy()
	if err != nil {
		log.Print("ERROR: Trying to initialize database", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	hero, err := hero.Retrieve(dbConn, params["id"])
	if err != nil {
		log.Printf("ERROR: Calling hero Retrieve Service function: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write the response.
	response, err := json.Marshal(&hero)
	if err != nil {
		log.Printf("ERROR: Failed while marshalling response: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}

// Delete deletes a hero from the heroes collection.
func (h *Hero) Delete(w http.ResponseWriter, r *http.Request, params map[string]string) {
	dbConn, err := h.MasterDB.Copy()
	if err != nil {
		log.Print("ERROR: Trying to initialize database", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	if err := hero.Delete(dbConn, params["id"]); err != nil {
		log.Printf("ERROR: Calling hero Delete Service function: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write the response.
	w.WriteHeader(http.StatusNoContent)
	return
}

// Update accepts updates a hero in the heroes collection.
func (h *Hero) Update(w http.ResponseWriter, r *http.Request, params map[string]string) {
	dbConn, err := h.MasterDB.Copy()
	if err != nil {
		log.Print("ERROR: Trying to initialize database", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	id := params["id"]

	retrievedHero, err := hero.Retrieve(dbConn, id)
	if err != nil {
		log.Printf("ERROR: Calling hero Retrieve Service function: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var uHero hero.UpdateHero
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("ERROR: could not read body: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &uHero)
	if err != nil {
		log.Printf("ERROR: Failed while unmarshalling: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = hero.Update(dbConn, uHero, id); err != nil {
		log.Print("ERROR: Trying to Call Update Service Function", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write the response.
	response, err := json.Marshal(&retrievedHero)
	if err != nil {
		log.Printf("ERROR: Failed while marshalling response: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	return
}
