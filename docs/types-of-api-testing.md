API Workshop - Types of API Testing
## Sections:

* [Bread Crumb Navigation](#bread-crumb-navigation)

#### Types of API Testing

[Types of API Testing](https://en.wikipedia.org/wiki/API_testing#Types_of_API_testing)

API testing typically involves the following practices:

* Unit testing - Testing the functionality of individual operations.

Here is an example of a unit test in golang:

```go
package config_test

import (
	"os"
	"testing"

	"github.com/jbelmont/api-workshop/cmd/apid/config"
)

func TestNewConfig(t *testing.T) {
	cfg := config.New()

	if cfg.APIHost != ":8080" {
		t.Error("config should set default value of APIHost variable")
	}
	if cfg.MongoHost != "mongodb://mongo:27017" {
		t.Error("config should set default value for MongoHost variable")
	}
	if cfg.RedisHost != "redis:6379" {
		t.Error("config should set default value for RedisHost variable")
	}
}

func TestGetEnvironmentVariablesThatAreSet(t *testing.T) {
	apiHost := ":3000"
	os.Setenv("API_HOST", apiHost)
	defer os.Unsetenv("API_HOST")

	mongoHost := "mongodb://mongo:30015"
	os.Setenv("MONGO_HOST", mongoHost)
	defer os.Unsetenv("MONGO_HOST")

	redisHost := "redis:7005"
	os.Setenv("REDIS_HOST", redisHost)
	defer os.Unsetenv("REDIS_HOST")

	cfg := config.New()

	if cfg.APIHost != apiHost {
		t.Errorf("Expected %s but got %s", apiHost, cfg.APIHost)
	}
	if cfg.MongoHost != mongoHost {
		t.Errorf("Expected %s but got %s", mongoHost, cfg.MongoHost)
	}
	if cfg.RedisHost != redisHost {
		t.Errorf("Expected %s but got %s", redisHost, cfg.RedisHost)
	}
}
```

* Integration Testing - Testing (sometimes called integration and testing, abbreviated I&T) is the phase in software testing in which individual software modules are combined and tested as a group. It occurs after unit testing and before validation testing.

Here is an example of an integration test in golang:

```go
func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	ctx = context.Background()

	c, err := container.StartDB()
	if err != nil {
		log.Fatalln(err)
	}
	container.SetContainer(c)

	defer func() {
		if err = container.StopDB(c); err != nil {
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

	// Initialize context
	log.Println("main : Started : Initialize context")
	v := apiContext.Values{
		ID: bson.NewObjectId().Hex(),
	}
	ctx = context.WithValue(context.Background(), apiContext.KeyValues, &v)

	// Creating DB state.
	log.Println("main : Started : Set DB state for user")
	// Insert test user
	heroInfo := heroInfo()
	f := func(collection *mgo.Collection) error {
		return collection.Insert(heroInfo.hero)
	}
	if err := masterDB.Execute(ctx, heroCollection, f); err != nil {
		log.Fatalf("startup : Set DB state : %v", err)
	}
	return m.Run()
}
```

* Functional testing - Testing the functionality of broader scenarios, often using unit tests as building blocks for end-to-end tests. Includes test case definition, execution, validation, and regression testing.

* Load testing - Validating functionality and performance under load, often by reusing functional test cases.

* Runtime error detection - Monitoring an application the execution of automated or manual tests to expose problems such as race conditions, exceptions, and resource leaks.

* Security testing - Includes penetration testing and fuzz testing as well as validating authentication, encryption, and access control.

* Web UI testing - Performed as part of end-to-end integration tests that also cover APIs, enables teams to validate GUI items in the context of the larger transaction.

Here is an example of an end to end ui test with nightwatch.js

```js
'use strict';

module.exports = {
  'Code Craftsmanship Saturdays': browser => {
    const codeSaturdays = browser.page.codeCraftsmanshipSaturdaysPage();

    codeSaturdays
      .navigate()
      .waitForElementVisible('@codeCraftsmanshipContainerLabel', 1000)
      .assert.containsText('@codeCraftsmanshipContainerStrongLabel', 'Code Craftsmanship Saturdays')

    codeSaturdays
      .click('@trashBin')
      .assert.elementNotPresent('[data-email="tcox0@dion.ne.jp"]')

    codeSaturdays
      .click('@addSomeUserBtn')

    codeSaturdays.setInput('#emailInput', 'chuck@badass.net')
    codeSaturdays.setInput('#firstNameInput', 'Chuck')
    codeSaturdays.setInput('#lastNameInput', 'Norris')

    codeSaturdays
      .click('@addUserSubmit')
      .waitForElementVisible('[data-email="chuck@badass.net"]', 1000)

    codeSaturdays
      .click('[data-email="dpayne3@cdbaby.com"]')
      .assert.urlContains('user')
      .end();
  }
};
```

For more information about Software Testing in General checkout my repo [Software Testing at Github](https://github.com/jbelmont/software-testing)

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [API Testing](./api-testing.md) | [API Testing Tools](./api-testing-tools.md) →
