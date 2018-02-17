API Workshop - Richardson Maturity Model

## Sections:

* [Bread Crumb Navigation](#bread-crumb-navigation)

#### What is the Richardson Maturity Model (RMM)

The Richardson Maturity Model (RMM) was coined by [Leonard Richardson](https://www.crummy.com/)

Leonard Richardson gave a talk on [QCon](https://qconferences.com/) presenting this model

RMM breaks down the maturity of Restful APIs into 4 levels:

* Level 0
* Level 1
* Level 2
* Level 3

#### Level 0

Level 0 lies on the realm of early web services.

These web services have the following:

* One [URI](https://en.wikipedia.org/wiki/Uniform_Resource_Identifier) and one [HTTP Method](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)

* [XML-RPC](https://en.wikipedia.org/wiki/XML-RPC) and [SOAP Services](https://en.wikipedia.org/wiki/SOAP)
  
#### Level 1

Level 1 is many [URIs](https://en.wikipedia.org/wiki/Uniform_Resource_Identifier) and one HTTP method

#### Level 2

Level 2 consists of:

* Many [URIs](https://en.wikipedia.org/wiki/Uniform_Resource_Identifier)
* Each URI consisting of multiple [HTTP Methods](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
* Proper [HTTP Status Codes](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Status_codes)

In this level you split out read operations, operations that want to fetch data, and treats them specially.

```http
GET /search?q=dogs
```

Here we do a `GET` request where we want to search for dogs

Most Restful APIs today are at level 2. 

They have many URIs and they utilize a lot of the HTTP verbs:

* GET
* POST
* PUT
* DELETE
* PATCH

A lot of Earlier APIs mainly used GET and POST instead of breaking down operations into further semantic meaning like a PATCH or a DELETE verb would signify

Richardson makes the point that [XML-RPC](https://en.wikipedia.org/wiki/XML-RPC) only uses POST operations so you can't really utilize safe and idempotent operations like a GET request.

Everything in XML-RPC is a POST request and you just wrap a response with the metainformation you need.

#### Level 3

Level 3 maturity is what [Roy Fielding](https://en.wikipedia.org/wiki/Roy_Fielding) defines as a Restful API.

In this level you implement [Hypermedia As The Engine Of Application State (HATEOAS)](https://en.wikipedia.org/wiki/HATEOAS)

It is easier for your clients to move from one resource to another if you embed the actual URIs in the HTTP response.

Richardson uses Amazons S3 service as an example where you have to use the S3 key and then some rules to turn something into a URI.

If instead Amazon provided a URI in the response it would both help the client and Amazon provide more discoverability.

"Leonard Richardson":

> Connections between resources are a form of data, and they should be described in the documents with the rest of the data. Let your clients focus on looking at that document and making decisions about what to do next. Not on internalizing your particular rules about where on the web you put your data.

"Leonard Richardson":

> There are two kinds of hypermedia: links and forms. Links tell you where the resources are, and forms tell you what you can do to the resources, and the media type tells you, among other things, how to extract the links and forms from the representation.

Web technologies are extensible including:

* HTTP such as [Atom Pub](https://en.wikipedia.org/wiki/Atom_(Web_standard))
* XML
* HTML [Microformats](https://en.wikipedia.org/wiki/Microformat)
* URI 
  * `http://`, `irc://`, `ftp://`, ...
  
#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [Rest Architecture Constraints](./rest-constraints.md) | [HTTP Concepts](./http-concepts.md) →
