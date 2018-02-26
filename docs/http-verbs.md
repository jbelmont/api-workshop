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

###### HTTP Response Codes

[RFC 7231 Section 6.3.1](https://tools.ietf.org/html/rfc7231#section-6.3.1)

Typically a 200 (OK) status code is used with a `OPTIONS` request.

* The 200 (OK) status code indicates that the request has succeeded.

###### HTTP Response Payloads

Typically an OPTIONS is a representation of the communications options like the HTTP verbs used for the particular resource

###### Example OPTIONS request

```curl
curl -X OPTIONS http://example.org -i

HTTP/1.1 200 OK
Allow: OPTIONS, GET, HEAD, POST
Cache-Control: max-age=604800
Date: Sun, 25 Feb 2018 14:56:31 GMT
Expires: Sun, 04 Mar 2018 14:56:31 GMT
Server: EOS (vny006/0452)
Content-Length: 0
```

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

* The HEAD method is identical to GET except that the server MUST NOT return a message-body in the response. 

* The metainformation contained in the HTTP headers in response to a HEAD request SHOULD be identical to the information sent in response to a GET request. 

* This method can be used for obtaining metainformation about the entity implied by the request without transferring the entity-body itself. 

* This method is often used for testing hypertext links for validity, accessibility, and recent modification.

###### HTTP Response Codes

[RFC 7231 Section 6.3.1](https://tools.ietf.org/html/rfc7231#section-6.3.1)

Typically a 200 (OK) status code is used with a `GET` request.

* The 200 (OK) status code indicates that the request has succeeded.

###### HTTP Response Payloads
   
For a `HEAD` request you get the same representation as GET, but without the representation data.

This means you cannot send a message body with `HEAD` request, you only get the HTTP Response Headers sent back.

###### Example HEAD Request

```curl
curl -X HEAD http://example.org/index.html -I

HTTP/1.1 200 OK
Content-Encoding: gzip
Accept-Ranges: bytes
Cache-Control: max-age=604800
Content-Type: text/html
Date: Sun, 25 Feb 2018 15:06:38 GMT
Etag: "1541025663+gzip"
Expires: Sun, 04 Mar 2018 15:06:38 GMT
Last-Modified: Fri, 09 Aug 2013 23:54:35 GMT
Server: ECS (dca/53DB)
X-Cache: HIT
Content-Length: 606
```

#### POST HTTP Method

[POST HTTP Method](https://tools.ietf.org/html/rfc2616#section-9.5)

> The POST method is used to request that the origin server accept the entity enclosed in the request as a new subordinate of the resource identified by the Request-URI in the Request-Line.

POST is designed to allow a uniform method to cover the following functions:

* Annotation of existing resources

* Posting a message to a bulletin board, newsgroup, mailing list, or similar group of articles

* Providing a block of data, such as the result of submitting a form, to a data-handling process

* Extending a database through an append operation.

###### HTTP Response Codes

[RFC 7231 Section 6.3.1](https://tools.ietf.org/html/rfc7231#section-6.3.1)

A 200 (OK) status code is used with a `POST` request to denote a successful operation.

* The 200 (OK) status code indicates that the request has succeeded.

[RFC 7231 Section 6.3.2](https://tools.ietf.org/html/rfc7231#section-6.3.2)

A 201 (CREATED) status code indicates that a new resource has been created

###### HTTP Response Payloads
   
A 200 response code you will get a message body

A 201 response code you will get a message body

More response codes are applicable please consult [RFC 7231 Section 6.1](https://tools.ietf.org/html/rfc7231#section-6.1)

###### Example POST Request

```http
POST /v1/payments/payment HTTP/1.1
Host: api.sandbox.paypal.com
Content-Type: application/json
Authorization: Bearer AFakeBearerToken12345
Cache-Control: no-cache

{
  "intent": "sale",
  "payer": {
  "payment_method": "paypal"
  },
  "transactions": [
  {
    "amount": {
    "total": "30.11",
    "currency": "USD",
    "details": {
      "subtotal": "30.00",
      "tax": "0.07",
      "shipping": "0.03",
      "handling_fee": "1.00",
      "shipping_discount": "-1.00",
      "insurance": "0.01"
    }
    },
    "description": "The payment transaction description.",
    "custom": "EBAY_EMS_90048630024435",
    "invoice_number": "48787589673",
    "payment_options": {
    "allowed_payment_method": "INSTANT_FUNDING_SOURCE"
    },
    "soft_descriptor": "ECHI5786786",
    "item_list": {
    "items": [
      {
      "name": "hat",
      "description": "Brown hat.",
      "quantity": "5",
      "price": "3",
      "tax": "0.01",
      "sku": "1",
      "currency": "USD"
      },
      {
      "name": "handbag",
      "description": "Black handbag.",
      "quantity": "1",
      "price": "15",
      "tax": "0.02",
      "sku": "product34",
      "currency": "USD"
      }
    ],
    "shipping_address": {
      "recipient_name": "Brian Robinson",
      "line1": "4th Floor",
      "line2": "Unit #34",
      "city": "San Jose",
      "country_code": "US",
      "postal_code": "95131",
      "phone": "011862212345678",
      "state": "CA"
    }
    }
  }
  ],
  "note_to_payer": "Contact us for any questions on your order.",
  "redirect_urls": {
  "return_url": "https://www.example.com/return",
  "cancel_url": "https://www.example.com/cancel"
  }
}
```

HTTP POST Response with `201` Response Code

```json
{
    "id": "PAY-60M79518TW7602539LKJOIWA",
    "intent": "sale",
    "state": "created",
    "payer": {
        "payment_method": "paypal"
    },
    "transactions": [
        {
            "amount": {
                "total": "30.11",
                "currency": "USD",
                "details": {
                    "subtotal": "30.00",
                    "tax": "0.07",
                    "shipping": "0.03",
                    "insurance": "0.01",
                    "handling_fee": "1.00",
                    "shipping_discount": "-1.00"
                }
            },
            "description": "The payment transaction description.",
            "custom": "EBAY_EMS_90048630024435",
            "invoice_number": "48787589673",
            "soft_descriptor": "ECHI5786786",
            "payment_options": {
                "allowed_payment_method": "INSTANT_FUNDING_SOURCE",
                "recurring_flag": false,
                "skip_fmf": false
            },
            "item_list": {
                "items": [
                    {
                        "name": "hat",
                        "sku": "1",
                        "description": "Brown hat.",
                        "price": "3.00",
                        "currency": "USD",
                        "tax": "0.01",
                        "quantity": 5
                    },
                    {
                        "name": "handbag",
                        "sku": "product34",
                        "description": "Black handbag.",
                        "price": "15.00",
                        "currency": "USD",
                        "tax": "0.02",
                        "quantity": 1
                    }
                ],
                "shipping_address": {
                    "recipient_name": "Brian Robinson",
                    "line1": "4th Floor",
                    "line2": "Unit #34",
                    "city": "San Jose",
                    "state": "CA",
                    "postal_code": "95131",
                    "country_code": "US",
                    "phone": "011862212345678"
                }
            },
            "related_resources": []
        }
    ],
    "note_to_payer": "Contact us for any questions on your order.",
    "create_time": "2018-02-25T16:29:12Z",
    "links": [
        {
            "href": "https://api.sandbox.paypal.com/v1/payments/payment/PAY-60M79518TW7602539LKJOIWA",
            "rel": "self",
            "method": "GET"
        },
        {
            "href": "https://www.sandbox.paypal.com/cgi-bin/webscr?cmd=_express-checkout&token=EC-4BG64645KB941924E",
            "rel": "approval_url",
            "method": "REDIRECT"
        },
        {
            "href": "https://api.sandbox.paypal.com/v1/payments/payment/PAY-60M79518TW7602539LKJOIWA/execute",
            "rel": "execute",
            "method": "POST"
        }
    ]
}
```


###### Create a test issue in Github 

```curl
curl -X post \
  https://api.github.com/repos/jbelmont/api-workshop/issues \
  -H 'Authorization: token ATestToken12345' \
  -H 'Content-Type: application/json' \
  -d '{
  "title": "Test Issue for Github",
  "body": "Please close me.",
  "assignees": [
    "jbelmont"
  ],
  "labels": [
    "bug"
  ]
}'
```

This returns the following JSON payload and a 201 status code:

```json
{
    "url": "https://api.github.com/repos/jbelmont/api-workshop/issues/4",
    "repository_url": "https://api.github.com/repos/jbelmont/api-workshop",
    "labels_url": "https://api.github.com/repos/jbelmont/api-workshop/issues/4/labels{/name}",
    "comments_url": "https://api.github.com/repos/jbelmont/api-workshop/issues/4/comments",
    "events_url": "https://api.github.com/repos/jbelmont/api-workshop/issues/4/events",
    "html_url": "https://github.com/jbelmont/api-workshop/issues/4",
    "id": 300071585,
    "number": 4,
    "title": "Test Issue for Github",
    "user": {
        "login": "jbelmont",
        "id": 2974744,
        "avatar_url": "https://avatars3.githubusercontent.com/u/2974744?v=4",
        "gravatar_id": "",
        "url": "https://api.github.com/users/jbelmont",
        "html_url": "https://github.com/jbelmont",
        "followers_url": "https://api.github.com/users/jbelmont/followers",
        "following_url": "https://api.github.com/users/jbelmont/following{/other_user}",
        "gists_url": "https://api.github.com/users/jbelmont/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/jbelmont/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/jbelmont/subscriptions",
        "organizations_url": "https://api.github.com/users/jbelmont/orgs",
        "repos_url": "https://api.github.com/users/jbelmont/repos",
        "events_url": "https://api.github.com/users/jbelmont/events{/privacy}",
        "received_events_url": "https://api.github.com/users/jbelmont/received_events",
        "type": "User",
        "site_admin": false
    },
    "labels": [
        {
            "id": 812066649,
            "url": "https://api.github.com/repos/jbelmont/api-workshop/labels/bug",
            "name": "bug",
            "color": "d73a4a",
            "default": true
        }
    ],
    "state": "open",
    "locked": false,
    "assignee": {
        "login": "jbelmont",
        "id": 2974744,
        "avatar_url": "https://avatars3.githubusercontent.com/u/2974744?v=4",
        "gravatar_id": "",
        "url": "https://api.github.com/users/jbelmont",
        "html_url": "https://github.com/jbelmont",
        "followers_url": "https://api.github.com/users/jbelmont/followers",
        "following_url": "https://api.github.com/users/jbelmont/following{/other_user}",
        "gists_url": "https://api.github.com/users/jbelmont/gists{/gist_id}",
        "starred_url": "https://api.github.com/users/jbelmont/starred{/owner}{/repo}",
        "subscriptions_url": "https://api.github.com/users/jbelmont/subscriptions",
        "organizations_url": "https://api.github.com/users/jbelmont/orgs",
        "repos_url": "https://api.github.com/users/jbelmont/repos",
        "events_url": "https://api.github.com/users/jbelmont/events{/privacy}",
        "received_events_url": "https://api.github.com/users/jbelmont/received_events",
        "type": "User",
        "site_admin": false
    },
    "assignees": [
        {
            "login": "jbelmont",
            "id": 2974744,
            "avatar_url": "https://avatars3.githubusercontent.com/u/2974744?v=4",
            "gravatar_id": "",
            "url": "https://api.github.com/users/jbelmont",
            "html_url": "https://github.com/jbelmont",
            "followers_url": "https://api.github.com/users/jbelmont/followers",
            "following_url": "https://api.github.com/users/jbelmont/following{/other_user}",
            "gists_url": "https://api.github.com/users/jbelmont/gists{/gist_id}",
            "starred_url": "https://api.github.com/users/jbelmont/starred{/owner}{/repo}",
            "subscriptions_url": "https://api.github.com/users/jbelmont/subscriptions",
            "organizations_url": "https://api.github.com/users/jbelmont/orgs",
            "repos_url": "https://api.github.com/users/jbelmont/repos",
            "events_url": "https://api.github.com/users/jbelmont/events{/privacy}",
            "received_events_url": "https://api.github.com/users/jbelmont/received_events",
            "type": "User",
            "site_admin": false
        }
    ],
    "milestone": null,
    "comments": 0,
    "created_at": "2018-02-25T23:14:45Z",
    "updated_at": "2018-02-25T23:14:45Z",
    "closed_at": null,
    "author_association": "OWNER",
    "body": "Please close me.",
    "closed_by": null
}
```

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

###### HTTP Response Codes

[RFC 7231 Section 6.3.1](https://tools.ietf.org/html/rfc7231#section-6.3.1)

A 200 (OK) status code is used with a `PUT` request to denote a successful operation.

[RFC 7231 Section 6.3.2](https://tools.ietf.org/html/rfc7231#section-6.3.2)

A 201 (CREATED) status code indicates that a new resource has been created

###### HTTP Response Payloads
   
A 200 response code you will get a message body

A 201 response code you will get a message body

More response codes are applicable please consult [RFC 7231 Section 6.1](https://tools.ietf.org/html/rfc7231#section-6.1)

###### Example PUT Request

```curl
curl -X put \
  https://api.github.com/repos/jbelmont/api-workshop/issues/4/labels \
  -H 'Accept: application/json' \
  -H 'Authorization: token ATestToken12345' \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '[
  "enhancement",
  "duplicate"
]'
```

This returns the following JSON payload and a 200 status code:

```json
[
    {
        "id": 812066652,
        "url": "https://api.github.com/repos/jbelmont/api-workshop/labels/duplicate",
        "name": "duplicate",
        "color": "cfd3d7",
        "default": true
    },
    {
        "id": 812066655,
        "url": "https://api.github.com/repos/jbelmont/api-workshop/labels/enhancement",
        "name": "enhancement",
        "color": "a2eeef",
        "default": true
    }
]
```

#### DELETE HTTP Method

[DELETE HTTP Method](https://tools.ietf.org/html/rfc2616#section-9.7)

* The DELETE method requests that the origin server delete the resource identified by the Request-URI. 

* This method MAY be overridden by human intervention (or other means) on the origin server. 

* The client cannot be guaranteed that the operation has been carried out, even if the status code returned from the origin server indicates that the action has been completed successfully. 

* However, the server SHOULD NOT indicate success unless, at the time the response is given, it intends to delete the resource or move it to an inaccessible location.

A successful response SHOULD be 200 (OK) if the response includes an entity describing the status, 202 (Accepted) if the action has not yet been enacted, or 204 (No Content) if the action has been enacted but the response does not include an entity.

###### HTTP Response Codes

[RFC 7231 Section 6.3.1](https://tools.ietf.org/html/rfc7231#section-6.3.1)

A 200 (OK) status code is used with a `PUT` request to denote a successful operation.

[RFC 7231 Section 6.3.5](https://tools.ietf.org/html/rfc7231#section-6.3.5)

A 204 (NO CONTENT) status code indicates indicates that the server has successfully fulfilled the request and that there is no additional content to send in the response payload body.

###### HTTP Response Payloads
   
A 200 response code you will get a message body

A 204 response code you will *not* get a message body

###### Example DELETE Request

```curl
curl -X delete \
  https://api.github.com/repos/jbelmont/api-workshop/issues/4/labels/duplicate \
  -H 'Authorization: token 9342bd73e336e5cf2fe7d1a84d77f206f65c66c8' \
  -H 'Content-Type: application/json'
```

This will return a 200 status code and the following JSON payload:

```json
[
    {
        "id": 812066655,
        "url": "https://api.github.com/repos/jbelmont/api-workshop/labels/enhancement",
        "name": "enhancement",
        "color": "a2eeef",
        "default": true
    }
]
```

#### TRACE HTTP Method

[TRACE HTTP Method](https://tools.ietf.org/html/rfc2616#section-9.8)

* The TRACE method is used to invoke a remote, application-layer loop-vback of the request message. 

* The final recipient of the request SHOULD reflect the message received back to the client as the entity-body of a 200 (OK) response. 

* The final recipient is either the origin server or the first proxy or gateway to receive a Max-Forwards value of zero (0) in the request (see section 14.31). 

* A TRACE request MUST NOT include an entity.

###### HTTP Response Codes

[RFC 7231 Section 6.3.1](https://tools.ietf.org/html/rfc7231#section-6.3.1)

A 200 (OK) status code is used with a `PUT` request to denote a successful operation.

[RFC 7231 Section 6.5.5](https://tools.ietf.org/html/rfc7231#section-6.5.5)

The 405 (Method Not Allowed) status code indicates that the method received in the request-line is known by the origin server but not supported by the target resource. 

The origin server MUST generate an Allow header field in a 405 response containing a list of the target resource's currently supported methods.

###### Example TRACE request

```curl
curl --insecure -v -X TRACE https://www.google.com/ -i
```

This will return a 405 status code (NOT ALLOWED)

#### CONNECT HTTP Method

[CONNECT HTTP Method](https://tools.ietf.org/html/rfc2616#section-9.9)

This specification reserves the method name CONNECT for use with a proxy that can dynamically switch to being a tunnel (e.g. SSL tunneling [44])

> The CONNECT method converts the request connection to a transparent TCP/IP tunnel, usually to facilitate SSL-encrypted communication (HTTPS) through an unencrypted HTTP proxy

Read more about this in [HTTP Connect Tunneling in Wikipedia](https://en.wikipedia.org/wiki/HTTP_tunnel#HTTP_CONNECT_tunneling)

###### Example Connect Request

```curl
curl --insecure -X CONNECT curl.haxx.se -v -i
```

Has the following output:

```http
* Rebuilt URL to: curl.haxx.se/
*   Trying 2a04:4e42:2f::561...
* TCP_NODELAY set
* Connected to curl.haxx.se (2a04:4e42:2f::561) port 80 (#0)
> CONNECT / HTTP/1.1
> Host: curl.haxx.se
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 301 Moved Permanently
HTTP/1.1 301 Moved Permanently
< Server: Varnish
Server: Varnish
< Retry-After: 0
Retry-After: 0
< Location: https://curl.haxx.se/
Location: https://curl.haxx.se/
< Content-Length: 0
Content-Length: 0
< Accept-Ranges: bytes
Accept-Ranges: bytes
< Date: Mon, 26 Feb 2018 00:07:45 GMT
Date: Mon, 26 Feb 2018 00:07:45 GMT
< Via: 1.1 varnish
Via: 1.1 varnish
< Connection: close
Connection: close
< X-Served-By: cache-dca17743-DCA
X-Served-By: cache-dca17743-DCA
< X-Cache: HIT
X-Cache: HIT
< X-Cache-Hits: 0
X-Cache-Hits: 0
< X-Timer: S1519603666.815007,VS0,VE0
X-Timer: S1519603666.815007,VS0,VE0

<
* Closing connection 0
```

#### PATCH HTTP Method

[PATCH HTTP Method](https://tools.ietf.org/html/rfc5789#section-1)

* The PATCH method requests that a set of changes described in the request entity be applied to the resource identified by the Request-URI.  

* The set of changes is represented in a format called a "patch document" identified by a media type.  

* If the Request-URI does not point to an existing resource, the server MAY create a new resource, depending on the patch document type (whether it can logically modify a null resource) and permissions, etc.

* The difference between the PUT and PATCH requests is reflected in the way the server processes the enclosed entity to modify the resource identified by the Request-URI.  

* In a PUT request, the enclosed entity is considered to be a modified version of the resource stored on the origin server, and the client is requesting that the stored version be replaced.  

* With PATCH, however, the enclosed entity contains a set of instructions describing how a resource currently residing on the origin server should be modified to produce a new version.  

* The PATCH method affects the resource identified by the Request-URI, and italso MAY have side effects on other resources; i.e., new resources may be created, or existing ones modified, by the application of a PATCH.

###### HTTP Response Codes

[RFC 7231 Section 6.3.1](https://tools.ietf.org/html/rfc7231#section-6.3.1)

A 200 (OK) status code is used with a `PUT` request to denote a successful operation.

[RFC 7231 Section 6.3.5](https://tools.ietf.org/html/rfc7231#section-6.3.5)

A 204 (NO CONTENT) status code indicates indicates that the server has successfully fulfilled the request and that there is no additional content to send in the response payload body.



###### HTTP Response Payloads
   
A 200 response code you will get a message body

A 204 response code you will *not* get a message body

More response codes are applicable please consult [RFC 7231 Section 6.1](https://tools.ietf.org/html/rfc7231#section-6.1)

###### Example PATCH Request

```curl
curl -X patch \
  https://api.github.com/repos/jbelmont/api-workshop/labels/enhancement \
  -H 'Accept: application/json' \
  -H 'Authorization: token 9342bd73e336e5cf2fe7d1a84d77f206f65c66c8' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "bug :bug:",
  "description": "Small bug fix required",
  "color": "b01f26"
}'
```

This returns a 200 status code with the following JSON payload:

```json
{
    "id": 812066655,
    "url": "https://api.github.com/repos/jbelmont/api-workshop/labels/bug%20:bug:",
    "name": "bug :bug:",
    "color": "b01f26",
    "default": false
}
```

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [HTTP Concepts](./http-concepts.md) | [HTTP Status Codes](./status-codes.md) →
