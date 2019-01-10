API Workshop - API Blueprint

## Sections:

* [API Blueprint Introduction](#api-blueprint-introduction)
* [Metadata section](#metadata-section)
* [API name & overview section](#api-name-\&-overview-section)
* [Resource group section](#resource-group-section)
* [Resource section](#resource-section)
* [Resource Actions](#resource-actions)
* [Schema section](#schema-section)
* [Action section](#action-section)
* [Request section](#request-section)
* [Response section](#response-section)
* [URI parameters section](#uri-parameters-section)
* [Attributes section](#attributes-section)
* [Headers section](#headers-section)
* [Body section](#body-section)
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

#### Schema section

add content here

#### Action section

add content here

#### Request section

add content here

#### Response section

add content here

#### URI parameters section

add content here

#### Attributes section

add content here

#### Headers section

add content here

#### Body section

add content here

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [OpenAPI](./openapi.md) | [RAML](./raml.md) →
