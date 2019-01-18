API Workshop - API Blueprint

## Sections:

* [API Blueprint Introduction](#api-blueprint-introduction)
* [Metadata section](#metadata-section)
* [API name & overview section](#api-name-\&-overview-section)
* [Resource group section](#resource-group-section)
* [Resource section](#resource-section)
* [Resource Actions](#resource-actions)
* [URI Template](#uri-template)
* [URI Parameters](#uri-parameters)
* [Bread Crumb Navigation](#bread-crumb-navigation)

#### API Blueprint Introduction

The first step in creating an api blueprint document is to specify the API Name and metadata.

*hello-world.md*

```apib
FORMAT: 1A

# API Workshop

A simple description right here.
```

#### Metadata section

An api blueprint document starts with a metadata section.

In the example above we specified a key of *FORMAT* and a value of *1A*.

The format keyword denotes the version of the API Blueprint which in this case is *1A*

#### API name & overview section

In your api blueprint document the first heading denoted by `#` should be the name of the api, which in this case is api-workshop.

Headings correspond to markdown headers which go up to 6 levels:

```apib
# foo         <h1>foo</h1>
## foo        <h2>foo</h2>
### foo       <h3>foo</h3>
#### foo      <h4>foo</h4>
##### foo     <h5>foo</h5>
###### foo    <h6>foo</h6>
```

Notice here that the headings correspond to h1 to h6 html headings.

The API Name here uses one hash to distinguish it as the first level.

The number of # you use will determine the level of the heading.

#### Resource group section

We use the *Group* keyword at the start of a heading for creating a group of related resources.

```apib
# Group Heroes

Resources related to heroes in the API.
```

#### Resource section

Now within the heroes resource groupe we can have a heroes collection.

The heroes resource shows a group of heroes.

*The heading specifies the URI used to access the resource inside of square brackets at the end of the heading.*

```apib
## Heroes Collection [/heroes]
```

#### Resource Actions

API Blueprint allows you to specify each action you may make on a resource.

An action is specified with a sub-heading inside of a resource with the name of the action followed by the HTTP method.

```apib
### List All Heroes [GET]
```

An action should include at least one response from the api and should include a status code and optionally a response body.

A response is defined as a list item within an action.

Lists are created by preceding list items with either a `+`, `*` or `-`.

This action returns a *200* status code along with a JSON body for the heroes endpoint:

```apib
+ Response 200 (application/json)
{
    "items": [
        {
            "id": "5c2b9750da0b9600a8210156",
            "Name": "Superman",
            "superPower": "",
            "gender": "male",
            "createdAt": "2019-01-01T16:37:36.905Z",
            "createdById": "",
            "updatedAt": "2019-01-01T16:37:36.905Z"
        },
        {
            "id": "5c30eed2cbf27800b60a2a78",
            "Name": "Superman",
            "superPower": "",
            "gender": "male",
            "createdAt": "2019-01-05T17:52:18.06Z",
            "createdById": "",
            "updatedAt": "2019-01-05T17:52:18.06Z"
        }
    ],
    "count": 2,
    "total": 2,
    "index": 0
}
```

When you specify a media type after the response status code blueprint generates a `Content-Type` HTTP header.

*The Content-Type header does not need to be explicitly set.*

The heroes resource has a second action which allows you to create a new hero.

This specific action includes a description that shows you the request payload needed to make this rest call.

```apib
### Create a New Hero [POST]

You may create your own hero using this action.

It takes a JSON object containing a hero and a heroes abilities and name.

+ name (string) - The hero's name
+ superPowers (array[string]) - A list of superpowers that a hero has.
+ gender (string) - The hero's gender
```

This action takes a JSON payload as part of the request as follows:

```apib
+ Request (application/json)

{
	"name": "Aquaman",
	"superPowers": [
		"super speed", "super strength", "telepathic abilities with animals"
	],
	"gender": "male"
}
```

This resource returns a 201 status code, along with HTTP headers and a response body:

```apib
Response 201 (application/json)

  + Headers

          Content-Type: application/json
          X-Workshop-Trace-Id: e697253f-b5af-43ed-8592-84dc4be8f072
          Date: Thu, 10 Jan 2019 23:00:44 GMT
          Content-Length: 214

    + Body

          {
            "id": "5c37ce9cac2eb400b9b783f9",
            "Name": "Aquaman",
            "superPower": "",
            "gender": "male",
            "createdAt": "2019-01-10T23:00:44.6509367Z",
            "createdById": "",
            "updatedAt": "2019-01-10T23:00:44.6509553Z"
          }
```

The next resource is a specific hero id, which represents a single hero.

```apib
## A specific hero [/api/v1/heroes/{hero_id}]
```

This resource returns a 200 status code, along with HTTP headers and a response body:

```apib
Response 200 (application/json)

  + Headers

          Content-Type: application/json
          X-Workshop-Trace-Id: ce21f483-bbdc-474f-a960-f1ac10e9b792
          Date: Sat, 12 Jan 2019 14:36:32 GMT
          Content-Length: 271

    + Body

          {
            "id": "5c39fb67ee458c00c072eade",
            "Name": "Aquaman",
            "superpowers": [
                "super speed",
                "super strength",
                "telepathic abilities with animals"
            ],
            "gender": "male",
            "created": "2019-01-12T14:36:23.606Z",
            "lastModified": "2019-01-12T14:36:23.606Z"
          }
```

#### URI Template

The URI for the Heroes resource uses a variable component, expressed by URI Template.

In this case, there is an ID variable called hero_id, represented in the URI template as {hero_id}.

#### URI Parameters

URI parameters should describe the URI using a list of Parameters.

For "Heroes" it would be as follows:

```apib
+ Parameters
    + hero_id (number) - ID of the Hero in the form of an MongoDB ObjectID
```

The hero_id variable of the URI template is a parameter for every action on this resource.

The hero_id variable is defined here using an arbitrary type number, followed by a description for the parameter.

###### Response with a Body

Here is rest call to the *Retrieve a Hero with an ID` endpoint that has json payload:

```apib
Response 200 (application/json)

  + Headers

          Content-Type: application/json
          X-Workshop-Trace-Id: ce21f483-bbdc-474f-a960-f1ac10e9b792
          Date: Sat, 12 Jan 2019 14:36:32 GMT
          Content-Length: 271

    + Body

          {
            "id": "5c39fb67ee458c00c072eade",
            "Name": "Aquaman",
            "superpowers": [
                "super speed",
                "super strength",
                "telepathic abilities with animals"
            ],
            "gender": "male",
            "created": "2019-01-12T14:36:23.606Z",
            "lastModified": "2019-01-12T14:36:23.606Z"
          }
```

###### Response without a body

This resource has a delete action and will return a 204 response code without a json response body:

```apib
### Delete [DELETE]

+ Response 204
```

###### Complete CRUD API Example

Here is the simple CRUD API example:

```apib
FORMAT: 1A
HOST: http://api-workshop.apiblueprint.org/

# marcelbelmont

Heroes is a simple CRUD API allowing consumers to create, list, retrieve, update and delete superheroes.

## Heroes Collection [/heroes]

### List All Heroes [GET]

+ Response 200 (application/json)

        {
            "items": [
                {
                    "id": "5c3a241e092bd900c4444c88",
                    "name": "Superman Prime",
                    "superpowers": [
                        "super speed",
                        "super strength",
                        "heat vision",
                        "Strongest Version of Superman"
                    ],
                    "gender": "male",
                    "created": "0001-01-01T00:00:00Z",
                    "lastModified": "2019-01-12T17:30:41.818Z"
                },
                {
                    "id": "5c3bd46c092bd9004c27cdee",
                    "name": "Aquaman",
                    "superpowers": [
                        "telepathic abilities",
                        "super strength",
                        "Intense heat resistance",
                        "Superhuman Hearing and Sonar"
                    ],
                    "gender": "male",
                    "created": "2019-01-14T00:14:36.158Z",
                    "lastModified": "2019-01-14T00:14:36.158Z"
                }
            ],
            "count": 2,
            "total": 2,
            "index": 0
        }

### Create a New Hero [POST]

You may create your own hero using this action. It takes a JSON
object containing a hero's name, superpowers, and gender.

+ Request (application/json)

        {
            "name": "Aquaman",
            "superpowers": [
                "telepathic abilities",
                "super strength",
                "Intense heat resistance",
                "Superhuman Hearing and Sonar"
            ],
            "gender": "male"
        }

+ Response 201 (application/json)

    + Headers

            X-Workshop-Trace-Id: 6831eb5c-de6c-4286-8af7-cc9f08100386
            Date: Mon, 14 Jan 2019 00:14:36 GMT
            Content-Length: 313

    + Body

            {
                "id": "5c3bd46c092bd9004c27cdee",
                "name": "Aquaman",
                "superpowers": [
                    "telepathic abilities",
                    "super strength",
                    "Intense heat resistance",
                    "Superhuman Hearing and Sonar"
                ],
                "gender": "male",
                "created": "2019-01-14T00:14:36.158902Z",
                "lastModified": "2019-01-14T00:14:36.1589207Z"
            }

## Heroes By ID [/heroes/{id}]

### Retrieve a hero by id [GET]

This endpoint retrieves a hero by ID and uses a path parameter of {id}

+ Response 200 (application/json)

        {
            "id": "5c3a241e092bd900c4444c88",
            "name": "Superman Prime",
            "superpowers": [
                "super speed",
                "super strength",
                "heat vision",
                "Strongest Version of Superman"
            ],
            "gender": "male",
            "created": "0001-01-01T00:00:00Z",
            "lastModified": "2019-01-12T17:30:41.818Z"
        }

### Update a hero by id [PUT]

This endpoint updates a hero's name, powers and gender.

+ Request (application/json)

        {
            "name": "Superman Prime",
            "superpowers": [
                "super speed", "super strength", "heat vision", "Strongest Version of Superman"
            ],
            "gender": "male"
        }

+ Response 200

    + Headers

            X-Workshop-Trace-Id: 8b75a7c0-96d1-4c62-bb96-35705cda80a2
            Date: Sat, 12 Jan 2019 17:30:41 GMT
            Content-Length: 250

    + Body

            {
                "id": "5c3a241e092bd900c4444c88",
                "name": "Superman",
                "superpowers": [
                    "super speed",
                    "super strength",
                    "heat vision"
                ],
                "gender": "male",
                "created": "2019-01-12T17:30:06.042Z",
                "lastModified": "2019-01-12T17:30:06.042Z"
            }


### Delete a hero by id [DELETE]

+ Response 204

    + Headers

            X-Workshop-Trace-Id: c4711ca8-5d82-4d3b-9857-6aa152a1b021
            Date: Mon, 14 Jan 2019 00:34:15 GMT
```

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [OpenAPI](./openapi.md) | [RAML](./raml.md) →
