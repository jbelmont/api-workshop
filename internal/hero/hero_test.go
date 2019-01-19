package hero_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jbelmont/api-workshop/internal/hero"
	database "github.com/jbelmont/api-workshop/internal/platform/db"
	"github.com/jbelmont/api-workshop/internal/platform/docker"
	mgo "gopkg.in/mgo.v2"
)

var masterDB *database.DB

const (
	heroCollection = "heroes"
)

type HeroInfo struct {
	hero hero.CreateHero
}

func heroInfo() HeroInfo {
	return HeroInfo{
		hero: hero.CreateHero{
			Name:   "SuperMan",
			Gender: "Male",
			SuperPowers: []string{
				"Super Strength, Super Speed, Flying, and other cool abilities",
			},
		},
	}
}

// TestCreateInsertsProperly validates users can be created in the database.
func TestCreateInsertHero(t *testing.T) {
	// Insert User
	heroInfo := heroInfo()
	// Translate CreateProduct to Product in order to assert result.
	cuBytes, _ := json.Marshal(heroInfo.hero)
	var exH hero.Hero
	json.Unmarshal(cuBytes, &exH)
	h, err := hero.Create(masterDB, &heroInfo.hero)
	if err != nil {
		t.Fatalf("Should be able to create a hero : %s.", err)
	}
	if !h.ID.Valid() {
		t.Errorf("Should generate valid Hero ID, got: %s", h.ID)
	}
	assertEqualUsers(&exH, h, t)
}

func assertEqualUsers(actual *hero.Hero, expected *hero.Hero, t *testing.T) {
	if actual.Name != expected.Name {
		t.Errorf("Should create hero Name: %s, but got: %s", expected.Name, actual.Name)
	}
}

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	c, err := docker.StartDB()
	if err != nil {
		log.Fatalln(err)
	}
	docker.SetContainer(c)

	defer func() {
		if err = docker.StopDB(c); err != nil {
			log.Println(err)
		}
	}()

	dbTimeout := 25 * time.Second
	dbHost := os.Getenv("MONGO_HOST")

	log.Println("main : Started : Initialize Mongo")
	masterDB, err = database.New(dbHost, dbTimeout)
	if err != nil {
		log.Fatalf("startup : Register DB : %v", err)
	}
	defer masterDB.Close()

	// Creating DB state.
	log.Println("main : Started : Set DB state for user")
	// Insert test user
	heroInfo := heroInfo()
	f := func(collection *mgo.Collection) error {
		return collection.Insert(heroInfo.hero)
	}
	if err := masterDB.Execute(heroCollection, f); err != nil {
		log.Fatalf("startup : Set DB state : %v", err)
	}
	return m.Run()
}
