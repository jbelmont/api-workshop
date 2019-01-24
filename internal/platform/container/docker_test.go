package container_test

import (
	"os"
	"testing"

	"github.com/jbelmont/api-workshop/internal/platform/container"
)

func TestSetContainer(t *testing.T) {
	ct := new(container.Container)
	ct.ID = "testID1234"
	ct.Port = "8708"

	err := container.SetContainer(ct)
	if err != nil {
		t.Error("Should not return error when setting environment variable in container.")
	}

	mgoHostEnv := "mongodb://localhost:8708/apitest"
	mongoHost := os.Getenv("MONGO_HOST")
	if mongoHost != mgoHostEnv {
		t.Errorf("Should set %s", mgoHostEnv)
	}
}

func TestStartDB(t *testing.T) {
	ct, err := container.StartDB()
	if err != nil {
		t.Error("Should not return error when starting db!")
	}

	if ct.ID == "" {
		t.Error("Should set ID field.")
	}
	if ct.Port == "" {
		t.Error("Should set Port field.")
	}
}

func TestStopDB(t *testing.T) {
	ct := new(container.Container)
	ct.ID = "stop1234"
	ct.Port = "8100"

	err := container.StopDB(ct)
	if err == nil {
		t.Errorf("Should not return error when stopping db: %s", err.Error())
	}
}
