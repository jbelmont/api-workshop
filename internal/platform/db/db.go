package mongo

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	mgo "gopkg.in/mgo.v2"
)

const apiDatabase = "api"

// DB holds values for DB drivers.
type DB struct {
	// MongoDB Support.
	database *mgo.Database
	session  *mgo.Session
}

// New returns a new DB value for use with MongoDB.
func New(url string, timeout time.Duration) (*DB, error) {

	if timeout == 0 {
		timeout = 60 * time.Second
	}

	// Create a session which maintains a pool of socket connections.
	session, err := mgo.DialWithTimeout(url, timeout)
	if err != nil {
		return nil, errors.Wrapf(err, "mgo.DialWithTimeout: %s,%v", url, timeout)
	}

	session.SetMode(mgo.Monotonic, true)

	db := DB{
		database: session.DB(apiDatabase),
		session:  session,
	}

	return &db, nil
}

// Close closes a DB value being used with MongoDB.
func (db *DB) Close() {
	db.session.Close()
}

// Copy returns a new DB value for use with MongoDB based on the master session.
func (db *DB) Copy() (*DB, error) {
	session := db.session.Copy()

	newDB := DB{
		database: session.DB(apiDatabase),
		session:  session,
	}

	return &newDB, nil
}

// Execute is used to execute MongoDB commands.
func (db *DB) Execute(collection string, fn func(*mgo.Collection) error) error {
	if db == nil || db.session == nil {
		return errors.New("DB error: db == nil || db.session == nil")
	}

	return fn(db.database.C(collection))
}

// ExecuteTimeout is used to execute MongoDB commands with a timeout.
func (db *DB) ExecuteTimeout(collection string, fn func(*mgo.Collection) error, timeout time.Duration) error {
	if db == nil || db.session == nil {
		return errors.New("DB error: db == nil || db.session == nil")
	}

	db.session.SetSocketTimeout(timeout)

	return fn(db.database.C(collection))
}

// Query provides a string version of the value
func Query(value interface{}) string {
	json, err := json.Marshal(value)
	if err != nil {
		return ""
	}

	return string(json)
}
