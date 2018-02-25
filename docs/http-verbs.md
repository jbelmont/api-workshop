API Workshop - HTTP Verbs

## Sections:

* [Bread Crumb Navigation](#bread-crumb-navigation)

#### What are HTTP methods?

[HTTP Methods](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)

HTTP defines methods which are sometimes called HTTP verbs which indicate an action that is to be performed on a identified resource.

> What this resource represents, whether pre-existing data or data that is generated dynamically, depends on the implementation of the server.

The resource can correspond to a specific file on a server or can even be the output of an executable on a server.

The HTTP 1.0 specification specified the following HTTP methods:

* [GET](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
* [POST](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
* [HEAD](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)

You can read more about this in [RFC 1945 Section 8.1](https://tools.ietf.org/html/rfc1945#section-8.1)

The HTTP 1.1 specification added 5 new methods:

* [OPTIONS](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
* [PUT](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
* [DELETE](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
* [TRACE](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
* [CONNECT](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)

You can read more about this in [RFC 2616 Section 9](https://tools.ietf.org/html/rfc2616#section-9)

Since these methods are specified in [RFC 1945](https://tools.ietf.org/html/rfc1945) and [RFC 2616](https://tools.ietf.org/html/rfc2616) the semantics of these HTTP methods are well known and can be depended on. 

*Any client can use any method and the server can be configured to support any combination of methods.*

**If a method is an unknown to an intermediate server then it will treated as an unsafe and non-idempotent method**

There is no limit to the number of methods that can be defined and this allows for future methods to be specified without breaking existing infrastructure.

For instance [RFC 5789](https://tools.ietf.org/html/rfc5789) specified the PATCH method.

#### Safe HTTP Method

The following HTTP Methods are considerd `safe`:

* [GET](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
* [HEAD](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
* [OPTIONS](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
* [TRACE](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)

**An HTTP method is considered `safe` if it is only intended for information retrieval and does not change the state of the server.**

#### Idempotent Method

[Idempotence](https://en.wikipedia.org/wiki/Idempotence)

> Idempotence is the property of certain operations in mathematics and computer science that they can be applied multiple times without changing the result beyond the initial application

HTTP Methods have the property of "idempotence" in that (aside from error or expiration issues) the side-effects of N > 0 identical requests is the same as for a single request.

[Idempotence in HTTP Methods](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Idempotent_methods_and_web_applications)

> (note that idempotence refers to the state of the system after the request has completed, so while the action the server takes (e.g. deleting a record) or the response code it returns may be different on subsequent requests, the system state will be the same every time

In this sense `PUT` and `DELETE` HTTP methods are idempotent even though multiple calls can return a different HTTP response code. 

For example calling a resource such as `DELETE `/users/:id` could return `204` status code on first operation and then a `404` on the second and subsequent requests.

The following HTTP methods are idempotent:

* [GET](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
* [HEAD](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
* [PUT](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
* [DELETE](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
* [OPTIONS](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
* [TRACE](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)

#### HTTP Method Summary Table

| HTTP Method | RFC | Has Request Body | Has Response Body | Safe | Idempotent | Cacheable |
| GET | [RFC 2616](https://tools.ietf.org/html/rfc2616#section-9) | Optional | Yes | Yes | Yes | Yes |
| HEAD | [RFC 2616](https://tools.ietf.org/html/rfc2616#section-9) | No | No | Yes | Yes | Yes |
| POST | [RFC 2616](https://tools.ietf.org/html/rfc2616#section-9) | Yes | Yes | No | No | Yes |
| PUT | [RFC 2616](https://tools.ietf.org/html/rfc2616#section-9) | Yes | Yes | No | Yes | No |
| DELETE | [RFC 2616](https://tools.ietf.org/html/rfc2616#section-9) | No | Yes | No | Yes | No |
| CONNECT | [RFC 2616](https://tools.ietf.org/html/rfc2616#section-9) | Yes | Yes | No | No | No |
| OPTIONS | [RFC 2616](https://tools.ietf.org/html/rfc2616#section-9) | Optional | Yes | yes | Yes | No |
| TRACE | [RFC 2616](https://tools.ietf.org/html/rfc2616#section-9) | No | Yes | Yes | Yes | No |
| PATCH | [RFC 2616](https://tools.ietf.org/html/rfc2616#section-9) | Yes | Yes | No | No | No |

#### GET HTTP Method

content

#### HEAD HTTP Method

content

#### POST HTTP Method

content

#### PUT HTTP Method

content

#### DELETE HTTP Method

content

#### OPTIONS HTTP Method

content

#### TRACE HTTP Method

content

#### CONNECT HTTP Method

content

#### PATCH HTTP Method

content

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [HTTP Concepts](./http-concepts.md) | [HTTP Status Codes](./status-codes.md) →
