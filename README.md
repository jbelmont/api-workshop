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

Make sure that you have [Golang](https://golang.org/dl/) Installed

#### Install Dep binary

Please run the following commands to install Dep which is a dependency manager in Golang:

###### Dep Install Script with official shell script

- `curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh`

###### Homebrew Installation for Mac OS x

Install or upgrade to the latest released version with Homebrew:

1. brew install dep
2. brew upgrade dep


###### Install GoMetalinter

Binary Releases
To install the latest stable release:

curl -L https://git.io/vp6lP | sh
Alternatively you can install a specific version from the releases list.

###### Install GoMetalinter with Homebrew 

brew tap alecthomas/homebrew-tap
brew install gometalinter

Once you have dep installed run the *make* command in your terminal

Run `make` in the root of the repository to install go libraries needed for the workshop.

##### Add common.env file to the repository

Please make sure to rename the *common.env.sample* to *common.env* because the API needs it to run.

##### Cloning instructions to proper GOPATH locations

I would highly suggest that you place the cloned repository into the following directory path:

`$HOME/go/src/github.com/{github_username}`

*Golang is very opionated about the package location and structure.*

#### Running API

Run `make dev` in order to get API running locally

This gets both [Redis](https://redis.io/) and [MongoDB](https://www.mongodb.com/) running in docker containers as well as a Golang Container running.

#### Running API Tests

Run `make test` in order to run the tests in a Go container
