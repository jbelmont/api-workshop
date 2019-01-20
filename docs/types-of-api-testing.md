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
