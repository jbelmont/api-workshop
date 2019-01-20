package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jbelmont/api-workshop/internal/hero"
)

func TestCreateHero(t *testing.T) {
	heroPayload := `
  {
      "name": "Aquaman",
      "superpowers": [
        "Expert with magical Trident",
        "Enhanced vision",
        "Enhanced smell",
        "Enhanced stamina",
        "Expert combatant",
        "Expert tactician",
        "Super Strength",
        "Super Speed",
        "Marine Telepathy"
      ],
      "gender": "male"
  }
  `
	payload := []byte(heroPayload)
	r := httptest.NewRequest("POST", "/api/v1/heroes", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()
	app.ServeHTTP(w, r)

	if w.Code != http.StatusCreated {
		t.Fatalf("Should receive a status code of 201 for the response : %v", w.Code)
	}

	var cHero hero.CreateHero
	if err := json.NewDecoder(w.Body).Decode(&cHero); err != nil {
		t.Fatalf("could not decode json response: %v", err)
	}

	expected := hero.Hero{
		Name: "Aquaman",
		SuperPowers: []string{
			"Expert with magical Trident",
			"Enhanced vision",
			"Enhanced smell",
			"Enhanced stamina",
			"Expert combatant",
			"Expert tactician",
			"Super Strength",
			"Super Speed",
			"Marine Telepathy",
		},
		Gender: "male",
	}
	if cHero.Name != expected.Name {
		t.Errorf("Name should be set to %s", expected.Name)
	}
	for idx, power := range cHero.SuperPowers {
		if power != expected.SuperPowers[idx] {
			t.Errorf("Expected Power %s, Actual %s", expected.SuperPowers[idx], power)
		}
	}
	if cHero.Gender != expected.Gender {
		t.Errorf("Expected Gender %s, Actual %s", expected.Gender, cHero.Gender)
	}
}
