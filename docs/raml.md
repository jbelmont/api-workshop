API Workshop - RAML

## Sections:

* [What is RAML](#what-is-raml)
* [Root](#root)
* [Resources](#resources)
* [Methods](#methods)
* [URI Parameters](#uri-parameters)
* [Query Parameters](#query-parameters)
* [Responses](#responses)
* [Body Parameters](#body-parameters)
* [Extract Schemas](#extract-schemas)
* [Resource Types](#resource-types)
* [Parameters](#parameters)
* [Includes](#includes)
* [RAML Specification](#raml-specification)
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
    delete:
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

#### Body Parameters

You can define several ways for body parameters in *raml:*

```raml
body:
  application/json:
    type: |
      {
        "type": "object",
        "$schema": "http://json-schema.org/draft-04/schema",
        "id": "http://jsonschema.net",
        "required": true,
        "properties": {
          "name": {
            "type": "string",
            "required": true
          },
          "superpowers": {
            "type": "array",
            "required": true
          },
          "gender": {
            "type": "string",
            "required": true
          },
          "created": {
            "format": "date-time",
            "formatMinimum": "2013-11-17T00:00Z",
            "formatMaximum": "2015-11-17T00:00Z"
          },
          "lastModified": {
            "format": "date-time",
            "formatMinimum": "2013-11-17T00:00Z",
            "formatMaximum": "2015-11-17T00:00Z"
          }
        }
      }
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

The schema above defines what the json used needs to fulfill for each attribute.

The represented object has:

* "name" property of type "string", and is required
* "superpowers" property of type "array", and is required
* etc

#### Extract Schemas

> One interesting RAML feature is the ability to extract the schemas and reference them by name.


There are three major advantages of extracting schemas and referencing the schemas by name:

* Improve RAML readability
* Allow reusing the schemas in several sections
* content

```raml
types:
  hero: {
    "type": "object",
    "$schema": "http://json-schema.org/draft-04/schema",
    "id": "http://jsonschema.net",
    "required": true,
    "properties": {
      "name": {
        "type": "string",
        "required": true
      },
      "superpowers": {
        "type": "array",
        "required": true
      },
      "gender": {
        "type": "string",
        "required": true
      },
      "created": {
        "format": "date-time",
        "formatMinimum": "2013-11-17T00:00Z",
        "formatMaximum": "2015-11-17T00:00Z"
      },
      "lastModified": {
        "format": "date-time",
        "formatMinimum": "2013-11-17T00:00Z",
        "formatMaximum": "2015-11-17T00:00Z"
      }
    }
  }
```

```raml
body:
  application/json:
    type: hero
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

*Notice that above we simply used the *type* of hero defined in the types section.*

#### RESOURCE TYPES

RAML allows you to be able to declare different types of resources which are called:

**Resource Types in RAML**

```raml
resourceTypes:
  collection:
    get:
    post:
  collection-item:
    get:
    delete:
    update:
```

###### Defining and parameterizing resourceTypes

* Description
* GET method with:
  * description
  * response for HTTP status 200 (which body's content type is "application/json")
* POST method with:
  * description
  * "access_token" queryParameter
  * bodyParameter with "application/json" contentType and validated by a Schema
  * response with HTTP status 201 (which body's content type is "application/json")
* PUT method with:
  * description
  * "access_token" queryParameter
  * bodyParameter with "application/json" contentType and validated by a Schema
  * response with HTTP status 200 (which body's content type is "application/json")
* DELETE method with:
  * description
  * "access_token" queryParameter
  * response with HTTP status 204 (which body's content type is "application/json")

*Now with this stated above let us define some resource types.*

```raml
resourceTypes:
  collection:
    description: Collection of Heroes
    get:
      description: Get a list of heroes
      responses:
        200:
          body:
            application/json:
    post:
      description: |
        Add a new hero to collection
      queryParameters:
        access_token:
          description: "The access token provided by the authentication application"
          example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
          required: true
          type: string
      body:
        application/json:
          type: song
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
                  "created": "2019-01-18T22:46:54.5929963Z",
                  "lastModified": "2019-01-18T22:46:54.5930135Z"
                }
```

A Reserved Parameter is a parameter with a value automatically specified by its context.

For the resourceTypes case, there are two Reserved Parameters: resourcePath and resourcePathName.

For the /heroes example, the values will be "/heroes" and "hero" respectively.

```raml
resourceTypes:
  collection:
    description: Collection of available <<resourcePathName>> in Heroes
    get:
      description: Get a list of <<resourcePathName>> based on the hero.
      responses:
        200:
          body:
            application/json:
    post:
      description: |
        Add a new <<resourcePathName|!singularize>> to Heroes Collection.
      queryParameters:
        access_token:
          description: "The access token provided by the authentication application"
          example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
          required: true
          type: string
      body:
        application/json:
          type: <<resourcePathName|!singularize>>
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
                  "created": "2019-01-18T22:46:54.5929963Z",
                  "lastModified": "2019-01-18T22:46:54.5930135Z"
                }
```

```raml
/heroes:
  type: collection
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
    responses:
      200:
        body:
          application/json:
            example: |
              {
                "items": [
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
                    },
                    {
                        "id": "5c43432466508100c14a6538",
                        "name": "Immortal Hulk",
                        "superpowers": [
                            "Superhuman Speed",
                            "Superhuman Strength",
                            "Transformation",
                            "Self Sustenance",
                            "Superhuman Stamina",
                            "Superhuman Durability",
                            "Regenerative Healing Factor",
                            "Resistance to Psychic Control",
                            "Immunity to All Diseases and Viruses",
                            "Superhuman Leaping Ability",
                            "Astral Form Perception",
                            "Homing Ability",
                            "Gamma Radiation/Energy Manipulation and Emission",
                            "Adaptation to Hostile Environments"
                        ],
                        "gender": "male",
                        "created": "2019-01-19T15:32:52.824Z",
                        "lastModified": "2019-01-19T15:32:52.824Z"
                    }
                ],
                "count": 2,
                "total": 2,
                "index": 0
            }
  post:
    body:
      application/json:
        example: |
          {
              "id": "5c43432866508100c14a6539",
              "name": "Immortal Hulk",
              "superpowers": [
                  "Superhuman Speed",
                  "Superhuman Strength",
                  "Transformation",
                  "Self Sustenance",
                  "Superhuman Stamina",
                  "Superhuman Durability",
                  "Regenerative Healing Factor",
                  "Resistance to Psychic Control",
                  "Immunity to All Diseases and Viruses",
                  "Superhuman Leaping Ability",
                  "Astral Form Perception",
                  "Homing Ability",
                  "Gamma Radiation/Energy Manipulation and Emission",
                  "Adaptation to Hostile Environments"
              ],
              "gender": "male",
              "created": "2019-01-19T15:32:56.1248085Z",
              "lastModified": "2019-01-19T15:32:56.1248264Z"
          }
```

Some more resource usage and resourceType definitions in *raml*:

```raml
collection-item:
  description: Entity representing a <<resourcePathName|!singularize>>
  get:
    description: |
      Get the <<resourcePathName|!singularize>>
      with <<resourcePathName|!singularize>>Id =
      {<<resourcePathName|!singularize>>Id}
    responses:
      200:
        body:
          application/json:
      404:
        body:
          application/json:
            example: |
              {"message": "<<resourcePathName|!singularize>> not found" }
```

#### Parameters

> At the moment of defining the parameter in the resourceType (with the placeholder), there is no difference between a parameter and a reserved parameter. The actual difference only appears when passing the parameter at the resource level.

```raml
/{id}:
  type:
    collection-item:
      exampleItem: An example here
```

```raml
resourceTypes:
  collection:
    description: Collection of available <<resourcePathName>> in Heroes
    get:
      description: Get a list of <<resourcePathName>> based on the hero.
      responses:
        200:
          body:
            application/json:
    post:
      description: |
        Add a new <<resourcePathName|!singularize>> to heroes collection.
      queryParameters:
        access_token:
          description: "The access token provided by the authentication application"
          example: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
          required: true
          type: string
      body:
        application/json:
          type: <<resourcePathName|!singularize>>
      responses:
        200:
          body:
            application/json:
              example: |
                { "message": "The <<resourcePathName|!singularize>> has been properly entered" }
  collection-item:
    description: Entity representing a <<resourcePathName|!singularize>>
    get:
      description: |
        Get the <<resourcePathName|!singularize>>
        with <<resourcePathName|!singularize>>Id =
        {<<resourcePathName|!singularize>>Id}
      responses:
        200:
          body:
            application/json:
        404:
          body:
            application/json:
              example: |
                {"message": "<<resourcePathName|!singularize>> not found" }
```

#### Includes

The!includes directive in RAML allows you to build file-distributed API definitions, which is useful for both code reuse and readability.

Here is an example:

```raml
/heroes:
  type:
    collection:
      exampleCollection: !include heroes.sample
      exampleItem: !include hero-new.sample
  /{songId}:
    type:
      collection-item:
        exampleItem: !include hero-retrieve.sample
```

#### RAML Specification

Here is the official [RAML Specification](https://github.com/raml-org/raml-spec/blob/master/versions/raml-10/raml-10.md/)

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [API Blueprint](./api-blueprint.md) | [GraphQL](./graphql.md) →
