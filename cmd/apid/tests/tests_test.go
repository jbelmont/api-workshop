package tests

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/jbelmont/api-workshop/cmd/apid/config"
	"github.com/jbelmont/api-workshop/cmd/apid/handlers"
	"github.com/jbelmont/api-workshop/internal/platform/container"
	"github.com/jbelmont/api-workshop/internal/platform/database"
)

var app http.Handler

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	// Create test database.
	c, err := container.StartDB()
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err = container.StopDB(c); err != nil {
			log.Println(err)
		}
	}()

	// Initialize configuration.
	cfg := config.New()
	cfg.MongoHost = container.GetDBHost(c)

	// Initialize application.
	log.Println("main : Started : Initialize Mongo")
	masterDB, err := database.New(cfg.MongoHost, cfg.DBTimeout)
	if err != nil {
		log.Fatalf("startup : Register DB : %v", err)
	}
	defer masterDB.Close()

	apiDB, err := masterDB.Copy()
	if err != nil {
		log.Fatalf("copying: copying database: %v", err)
	}
	defer apiDB.Close()
	app = handlers.API(masterDB, cfg)

	return m.Run()
}
