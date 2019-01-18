API Workshop - RAML

## Sections:

* [What is RAML](#what-is-raml)
* [Root](#root)
* [Resources](#resources)
* [Methods](#methods)
* [URI Parameters](#uri-parameters)
* [Query Parameters](#query-parameters)
* [Responses](#responses)
* [Base RAML file]()
* [Body Parameters]()
* [Extract Schemas]()
* [Resource Types]()
* [Parameters]()
* [Includes]()
* [Refactor]()
* [Traits]()
* [Final Tuning]()

* [Bread Crumb Navigation](#bread-crumb-navigation)

#### What is RAML

> RESTful API Modeling Language (RAML) makes it easy to manage the whole API lifecycle from design to sharing. It's concise - you only write what you need to define - and reusable. It is machine readable API design that is actually human friendly.

![Raml Homepage](../images/raml.png)

#### Root

* Everything you enter in at the root (or top) of the spec applies to the rest of your API.

* The baseURI you choose will be used with every Rest Call made

```raml
#%RAML 1.0
---
title: Code Craftsmanship Saturdays API
baseUri: http://localhost:8080
version: v1
```

#### Resources

In RAML you define resources with the following syntax:

```raml
/heroes:

/heroes:
  /{id}:
```

*These resources go under the Root*

#### Methods

You can use the following HTTP Verbs with RAML:

* [GET](https://tools.ietf.org/html/rfc2616#page-53)
* [PATCH](https://tools.ietf.org/html/rfc5789)
* [PUT](https://tools.ietf.org/html/rfc2616#page-55)
* [POST](https://tools.ietf.org/html/rfc2616#page-54)
* [DELETE](https://tools.ietf.org/html/rfc2616#page-56)
* [OPTIONS](https://tools.ietf.org/html/rfc2616#page-52)
* [HEAD](https://tools.ietf.org/html/rfc2616#page-54)

*Each HTTP method can only be used once per resource.*

Nest the methods to allow developers to perform these actions under your resources.

Note that you must use lower-case for methods in your RAML API definition:

```raml
/heroes:
  get:
  post:
  /{id}:
    put:
    delete:
    get:
```

#### URI Parameters

The resources that we defined are collections of smaller, relevant objects.

This is a URI parameter, denoted by surrounding curly brackets in RAML:

```raml
/heroes:
  /{id}:
```

Now in order to make a request to this resource the uri would look like this:

`http://localhost:8080/api/v1/heroes/5c3a241e092bd900c4444c88`

#### Query Parameters

API Consumers are also able to perform actions like filtering a collection and Query parameters allow you to accomplish filtering a collection.

```raml
/heroes:
  get:
    queryParameters:
      name:
      superpowers:
      gender:
  post:
```

> Query parameters may also be something that the server requires to process the API consumer's request, like an access token. Often, you need security authorization to alter a collection or record.

```raml
/heroes:
  get:
    queryParameters:
      name:
      superpowers:
      gender:
    put:
      queryParameters:
        access_token:
```

> An API's resources and methods often have a number of associated query parameters. Each query parameter may have any number of optional attributes to further define it.

We can specify attributes for each of the query string parameters like this:

```raml
/heroes:
  /{id}
    get:
      queryParameters:
        name:
          displayName: Name
          type: string
          description: A heroes name
          example: Superman Prime
          required: true
        superpowers:
          displayName: Super Powers
          type: array
          description: A list of super hero powers that each hero has
          example: ["super speed", "super strength", "heat vision", "flight", "invulnerable"]
          required: true
        gender:
          displayName: Gender
          type: string
          description: A heroes gender
          example: Male
          required: true
    put:
      queryParameters:
        access_token:
          displayName: Access Token
          type: string
          description: Token giving you permission to make the rest call
          required: true
```

#### Responses

* Responses MUST be a map of one or more HTTP status codes, and each response may include descriptions, examples, or schemas.

```raml
/heroes:
  /{id}:
    get:
      description: Retrieve a specific hero
      responses:
        200:
          body:
            application/json:
              example: |
                {
                  "id": "5c42575e77802e00bf9ed411",
                  "name": "Aquaman",
                  "superpowers": [
                      "telepathic abilities",
                      "super strength",
                      "Intense heat resistance",
                      "Superhuman Hearing and Sonar"
                  ],
                  "gender": "male",
                  "created": "2019-01-18T22:46:54.592Z",
                  "lastModified": "2019-01-18T22:46:54.593Z"
                }
```

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [API Blueprint](./api-blueprint.md) | [GraphQL](./graphql.md) →
