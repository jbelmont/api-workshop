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
| --- | --- | --- | --- | --- | --- | --- |
| GET | [RFC 2616](https://tools.ietf.org/html/rfc2616#section-9) | Optional | Yes | Yes | Yes | Yes |
| HEAD | [RFC 2616](https://tools.ietf.org/html/rfc2616#section-9) | No | No | Yes | Yes | Yes |
| POST | [RFC 2616](https://tools.ietf.org/html/rfc2616#section-9) | Yes | Yes | No | No | Yes |
| PUT | [RFC 2616](https://tools.ietf.org/html/rfc2616#section-9) | Yes | Yes | No | Yes | No |
| DELETE | [RFC 2616](https://tools.ietf.org/html/rfc2616#section-9) | No | Yes | No | Yes | No |
| CONNECT | [RFC 2616](https://tools.ietf.org/html/rfc2616#section-9) | Yes | Yes | No | No | No |
| OPTIONS | [RFC 2616](https://tools.ietf.org/html/rfc2616#section-9) | Optional | Yes | Yes | Yes | No |
| TRACE | [RFC 2616](https://tools.ietf.org/html/rfc2616#section-9) | No | Yes | Yes | Yes | No |
| PATCH | [RFC 2616](https://tools.ietf.org/html/rfc2616#section-9) | Yes | Yes | No | No | No |

#### OPTIONS HTTP Method

[OPTIONS HTTP Method](https://tools.ietf.org/html/rfc2616#section-9.2)

* The OPTIONS method represents a request for information about the communication options available on the request/response chain identified by the Request-URI. 

* This method allows the client to determine the options and/or requirements associated with a resource, or the capabilities of a server, without implying a resource action or initiating a resource retrieval.

> The OPTIONS method returns the HTTP methods that the server supports for the specified URL. This can be used to check the functionality of a web server by requesting '*' instead of a specific resource.

###### HTTP Response Codes and HTTP Response Payloads

content

#### GET HTTP Method

[GET HTTP Method](https://tools.ietf.org/html/rfc2616#section-9.3)

> The GET method means retrieve whatever information (in the form of an entity) is identified by the Request-URI. If the Request-URI refers to a data-producing process, it is the produced data which shall be returned as the entity in the response and not the source text of the process, unless that text happens to be the output of the process. 

> The GET method requests a representation of the specified resource. Requests using GET should only retrieve data and should have no other effect.

###### HTTP Response Codes

[RFC 7231 Section 6.3.1](https://tools.ietf.org/html/rfc7231#section-6.3.1)

Typically a 200 (OK) status code is used with a `GET` request.

* The 200 (OK) status code indicates that the request has succeeded.

###### HTTP Response Payloads
   
For a `GET` request the payload will be a representation of the target resource

###### Example GET Request

Here is a `GET` request from github that lists all users, in the order that they signed up on GitHub. This list includes personal user accounts and organization accounts.

```http
GET /users

Status: 200 OK
Link: <https://api.github.com/users?since=135>; rel="next"

[
  {
    "login": "octocat",
    "id": 1,
    "avatar_url": "https://github.com/images/error/octocat_happy.gif",
    "gravatar_id": "",
    "url": "https://api.github.com/users/octocat",
    "html_url": "https://github.com/octocat",
    "followers_url": "https://api.github.com/users/octocat/followers",
    "following_url": "https://api.github.com/users/octocat/following{/other_user}",
    "gists_url": "https://api.github.com/users/octocat/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/octocat/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/octocat/subscriptions",
    "organizations_url": "https://api.github.com/users/octocat/orgs",
    "repos_url": "https://api.github.com/users/octocat/repos",
    "events_url": "https://api.github.com/users/octocat/events{/privacy}",
    "received_events_url": "https://api.github.com/users/octocat/received_events",
    "type": "User",
    "site_admin": false
  }
]
```

#### HEAD HTTP Method

[HEAD HTTP Method](https://tools.ietf.org/html/rfc2616#section-9.4)

> The HEAD method is identical to GET except that the server MUST NOT return a message-body in the response. The metainformation contained in the HTTP headers in response to a HEAD request SHOULD be identical to the information sent in response to a GET request. This method can be used for obtaining metainformation about the entity implied by the request without transferring the entity-body itself. This method is often used for testing hypertext links for validity, accessibility, and recent modification.

#### POST HTTP Method

[POST HTTP Method](https://tools.ietf.org/html/rfc2616#section-9.5)

> The POST method is used to request that the origin server accept the entity enclosed in the request as a new subordinate of the resource identified by the Request-URI in the Request-Line.

POST is designed to allow a uniform method to cover the following functions:

* Annotation of existing resources

* Posting a message to a bulletin board, newsgroup, mailing list, or similar group of articles

* Providing a block of data, such as the result of submitting a form, to a data-handling process

* Extending a database through an append operation.

#### PUT HTTP Method

[PUT HTTP Method](https://tools.ietf.org/html/rfc2616#section-9.6)

* The PUT method requests that the enclosed entity be stored under the supplied Request-URI. 

* If the Request-URI refers to an already existing resource, the enclosed entity SHOULD be considered as a modified version of the one residing on the origin server. 

* If the Request-URI does not point to an existing resource, and that URI is capable of being defined as a new resource by the requesting user agent, the origin server can create the resource with that URI. 

* If a new resource is created, the origin server MUST inform the user agent via the 201 (Created) response. 

* If an existing resource is modified, either the 200 (OK) or 204 (No Content) response codes SHOULD be sent to indicate successful completion of the request. 

* If the resource could not be created or modified with the Request-URI, an appropriate error response SHOULD be given that reflects the nature of the problem.

* The recipient of the entity MUST NOT ignore any Content-* (e.g. Content-Range) headers that it does not understand or implement and MUST return a 501 (Not Implemented) response in such cases.

###### Difference between POST and PUT HTTP Methods

* The fundamental difference between the POST and PUT requests is
   reflected in the different meaning of the Request-URI. 
   
* The URI in a POST request identifies the resource that will handle the enclosed
   entity. 
   
* That resource might be a data-accepting process, a gateway to some other protocol, or a separate entity that accepts annotations.

* In contrast, the URI in a PUT request identifies the entity enclosed with the request -- the user agent knows what URI is intended and the server MUST NOT attempt to apply the request to some other resource.

* If the server desires that the request be applied to a different URI, it MUST send a 301 (Moved Permanently) response; the user agent MAY then make its own decision regarding whether or not to redirect the request.

#### DELETE HTTP Method

[DELETE HTTP Method](https://tools.ietf.org/html/rfc2616#section-9.7)

* The DELETE method requests that the origin server delete the resource identified by the Request-URI. 

* This method MAY be overridden by human intervention (or other means) on the origin server. 

* The client cannot be guaranteed that the operation has been carried out, even if the status code returned from the origin server indicates that the action has been completed successfully. 

* However, the server SHOULD NOT indicate success unless, at the time the response is given, it intends to delete the resource or move it to an inaccessible location.

A successful response SHOULD be 200 (OK) if the response includes an entity describing the status, 202 (Accepted) if the action has not yet been enacted, or 204 (No Content) if the action has been enacted but the response does not include an entity.

#### TRACE HTTP Method

[TRACE HTTP Method](https://tools.ietf.org/html/rfc2616#section-9.8)

* The TRACE method is used to invoke a remote, application-layer loop-vback of the request message. 

* The final recipient of the request SHOULD reflect the message received back to the client as the entity-body of a 200 (OK) response. 

* The final recipient is either the origin server or the first proxy or gateway to receive a Max-Forwards value of zero (0) in the request (see section 14.31). 

* A TRACE request MUST NOT include an entity.

#### CONNECT HTTP Method

[CONNECT HTTP Method](https://tools.ietf.org/html/rfc2616#section-9.9)

This specification reserves the method name CONNECT for use with a proxy that can dynamically switch to being a tunnel (e.g. SSL tunneling [44])

> The CONNECT method converts the request connection to a transparent TCP/IP tunnel, usually to facilitate SSL-encrypted communication (HTTPS) through an unencrypted HTTP proxy

Read more about this in [HTTP Connect Tunneling in Wikipedia](https://en.wikipedia.org/wiki/HTTP_tunnel#HTTP_CONNECT_tunneling)

#### PATCH HTTP Method

[PATCH HTTP Method](https://tools.ietf.org/html/rfc5789#section-1)

* The PATCH method requests that a set of changes described in the request entity be applied to the resource identified by the Request-URI.  

* The set of changes is represented in a format called a "patch document" identified by a media type.  

* If the Request-URI does not point to an existing resource, the server MAY create a new resource, depending on the patch document type (whether it can logically modify a null resource) and permissions, etc.

* The difference between the PUT and PATCH requests is reflected in the way the server processes the enclosed entity to modify the resource identified by the Request-URI.  

* In a PUT request, the enclosed entity is considered to be a modified version of the resource stored on the origin server, and the client is requesting that the stored version be replaced.  

* With PATCH, however, the enclosed entity contains a set of instructions describing how a resource currently residing on the origin server should be modified to produce a new version.  

* The PATCH method affects the resource identified by the Request-URI, and italso MAY have side effects on other resources; i.e., new resources may be created, or existing ones modified, by the application of a PATCH.


#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [HTTP Concepts](./http-concepts.md) | [HTTP Status Codes](./status-codes.md) →
