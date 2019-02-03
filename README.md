[![Build Status](https://travis-ci.org/jbelmont/api-workshop.svg?branch=master)](https://travis-ci.org/jbelmont/api-workshop)
[![Coverage Status](https://coveralls.io/repos/github/jbelmont/api-workshop/badge.svg?branch=master)](https://coveralls.io/github/jbelmont/api-workshop?branch=master)

# API Workshop

A workshop  on API Development, API Security, API Testing and API Description Languages

* [API Workshop Gitbook](#api-workshop-gitbook)
* [Workshop Details](#workshop-details)
* [API Summary and Description](docs/api-summary.md)
* [Restful API Introduction](docs/restful-intro.md)
  * [Roy Fieldings Dissertation on REST](docs/rest-dissertation.md)
  * [Rest Architecture Constraints](docs/rest-constraints.md)
  * [Richardson Maturity Model](docs/maturity-model.md)
* [HTTP Concepts](docs/http-concepts.md)
  * [HTTP Verbs](docs/http-verbs.md)
  * [HTTP Status Codes](docs/status-codes.md)
  * [HTTP Headers](docs/http-headers.md)
  * [HTTP Request and Response](docs/request-response.md)
  * [URI Design](docs/uri-design.md)
  * [Query String](docs/query-string.md)
* [API Design](docs/api-design.md)
* [API Testing](docs/api-testing.md)
  * [Types of API Testing](docs/types-of-api-testing.md)
  * [API Testing Tools](docs/api-testing-tools.md)
* [API Security](docs/api-security.md)
  * [Authentication](docs/authentication.md)
    * [OpenID Connect](docs/openid-connect.md)
  * [Authorization](docs/authorization.md)
    * [OAuth](docs/oauth.md)
  * [Federation](docs/federation.md)
  * [Delegation](docs/delegation.md)
* [API Description Languages](docs/api-description-languages.md)
  * [OpenAPI](docs/openapi.md)
  * [API Blueprint](docs/api-blueprint.md)
  * [RAML](docs/raml.md)
* [GraphQL](docs/graphql.md)


# API Workshop Gitbook

Here is the gitbook for the [API Workshop](https://www.marcelbelmont.com/api-workshop/)

## Workshop Details

Install [Golang](https://golang.org/dl/) and make sure to set $GOPATH to `$HOME/go`

Try to put your GO Files in the following location:

`$HOME/go/src/github.com/{github_username` 

Go is very opinionated about this and linters and tools are looking for a specific path.

Install Docker and Docker-compose. 

Here is my version information:

*Docker Version:*

```bash
docker version
Client: Docker Engine - Community
 Version:           18.09.1
 API version:       1.39
 Go version:        go1.10.6
 Git commit:        4c52b90
 Built:             Wed Jan  9 19:33:12 2019
 OS/Arch:           darwin/amd64
 Experimental:      false

Server: Docker Engine - Community
 Engine:
  Version:          18.09.1
  API version:      1.39 (minimum version 1.12)
  Go version:       go1.10.6
  Git commit:       4c52b90
  Built:            Wed Jan  9 19:41:49 2019
  OS/Arch:          linux/amd64
  Experimental:     true
```

*Docker Compose Version:*

```bash
> docker-compose version
docker-compose version 1.23.2, build 1110ad01
docker-py version: 3.6.0
CPython version: 3.6.6
OpenSSL version: OpenSSL 1.1.0h  27 Mar 2018
```

#### Install Dep binary

Please run the following commands to install Dep which is a dependency manager in Golang:

###### Dep Install Script with official shell script

- `curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh`

###### Homebrew Installation for Mac OS x

Install or upgrade to the latest released version with Homebrew:

1. brew install dep

2. brew upgrade dep


###### Install GoMetalinter with Curl

Binary Releases
To install the latest stable release:

curl -L https://git.io/vp6lP | sh

Alternatively you can install a specific version from the releases list.

###### Install GoMetalinter with Homebrew 

1. brew tap alecthomas/homebrew-tap

2. brew install gometalinter

##### Add common.env file to the repository

Please make sure to rename the *common.env.sample* to *common.env* because the API needs it to run.

##### Cloning instructions to proper GOPATH locations

I would highly suggest that you place the cloned repository into the following directory path:

`$HOME/go/src/github.com/{github_username}`

*Golang is very opionated about the package location and structure.*

#### Running API

Run `make build` to build the api and in particular the docker containers.

Run `make dev` in order to get API running locally

This gets both [Redis](https://redis.io/) and [MongoDB](https://www.mongodb.com/) running in docker containers as well as a Golang Container running.

###### Using Rest Client to work with API

I would suggest to use a Rest Client such as Postman to work with the running api:

*Example making GET Request to the Heroes LIST Endpoint:*

```bash
curl -X GET \
  http://localhost:8080/api/v1/heroes \
  -H 'Accept: application/json'
```

*Example making a POST Request to the Heroes Create Endpoint:*

```bash

curl -X POST \
  http://localhost:8080/api/v1/heroes \
  -H 'Content-Type: application/json' \
  -d '{
    "name": "Batman",
    "superpowers": [
        "Super Rich",
        "Tech Gadgets",
        "Batmobile",
        "Martial Artists",
        "Super Style",
        "Alfred"
    ],
    "gender": "male"
  }'
```

*Example making a GET Request retrieving a specific Hero:*

```bash
curl -X GET \
  http://localhost:8080/api/v1/heroes/5c4ca572912242004355b5d4 \
  -H 'Accept: application/json'
```

*Example making a DELETE Request to remove a specific hero:*

```bash
curl -X DELETE \
  http://localhost:8080/api/v1/heroes/5c4ca572912242004355b5d4 \
  -H 'Accept: application/json'
```

*Example making an UPDATE Request to update a specific hero:*

```bash
curl -X PUT \
  http://localhost:8080/api/v1/heroes/5c44982798204b00946bb860 \
  -H 'Accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
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
        "Marine Telepathy",
        "Super Reflexes"
    ],
    "gender": "male"
}'
```

#### Running API Tests

Run `make test` in order to run the tests in a Go container

##### TODO Section for Workshop

- [] Wrap up added sections in API Design Section
- [] Wrap up GraphQL Section
- [] Add section for API Introspection
- [] Add section for tool like API Dredd to test API Documentation
- [] Wrap new section to disambiguate differences between OAuth and OAuth2
