package database_test

import (
	"net/url"
	"testing"
	"time"

	"github.com/jbelmont/api-workshop/internal/platform/database"
	mgo "gopkg.in/mgo.v2"
)

func TestNewDB(t *testing.T) {
	url := url.URL{
		Host: "localhost:8080",
	}

	timeout, _ := time.ParseDuration("30m")

	_, err := database.New(url.String(), timeout)
	if err != nil {
		t.Error("Should return new db instance")
	}
}

func TestDBClose(t *testing.T) {
	url := url.URL{
		Host: "localhost:8080",
	}

	timeout, _ := time.ParseDuration("30m")

	db, err := database.New(url.String(), timeout)
	if err != nil {
		t.Error("Should return new db instance")
	}

	db.Close()
}

func TestDBCopy(t *testing.T) {
	url := url.URL{
		Host: "localhost:8080",
	}

	timeout, _ := time.ParseDuration("30m")

	db, err := database.New(url.String(), timeout)
	if err != nil {
		t.Error("Should return new db instance")
	}

	_, err = db.Copy()
	if err != nil {
		t.Error("Should return error when copying db!")
	}
}

func TestDBExecute(t *testing.T) {
	url := url.URL{
		Host: "localhost:8080",
	}

	timeout, _ := time.ParseDuration("45m")

	db, err := database.New(url.String(), timeout)
	if err != nil {
		t.Error("Should return new db instance")
	}

	f := func(collection *mgo.Collection) error {
		return nil
	}
	err = db.Execute("api", f)
	if err != nil {
		t.Error("Should be able to execute command.")
	}
}

func TestDBExecuteTimeout(t *testing.T) {
	url := url.URL{
		Host: "localhost:8080",
	}

	timeout, _ := time.ParseDuration("45m")

	db, err := database.New(url.String(), timeout)
	if err != nil {
		t.Error("Should return new db instance")
	}

	f := func(collection *mgo.Collection) error {
		return nil
	}
	err = db.ExecuteTimeout("api", f, timeout)
	if err != nil {
		t.Error("Should be able to execute command with timeout.")
	}
}

func TestQuery(t *testing.T) {
	q := struct {
		name string
	}{
		name: "john j rambo",
	}
	val := database.Query(q)
	if val == "" {
		t.Error("Should return representation of json value")
	}
}
