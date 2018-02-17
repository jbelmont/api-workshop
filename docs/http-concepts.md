API Workshop - HTTP Concepts

## Sections:

* [Bread Crumb Navigation](#bread-crumb-navigation)

#### What is HTTP

The Hypertext Transfer Protocol (HTTP) is an application protocol for distributed, collaborative, and hypermedia information systems.

#### HTTP Overview

[HTTP](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Technical_overview)

> HTTP functions as a request–response protocol in the client–server computing model. A web browser, for example, may be the client and an application running on a computer hosting a website may be the server. The client submits an HTTP request message to the server. The server, which provides resources such as HTML files and other content, or performs other functions on behalf of the client, returns a response message to the client. The response contains completion status information about the request and may also contain requested content in its message body.

> HTTP is an application layer protocol designed within the framework of the Internet protocol suite. Its definition presumes an underlying and reliable transport layer protocol, and Transmission Control Protocol (TCP) is commonly used. However HTTP can be adapted to use unreliable protocols such as the User Datagram Protocol (UDP), for example in HTTPU and Simple Service Discovery Protocol (SSDP).

> HTTP resources are identified and located on the network by Uniform Resource Locators (URLs), using the Uniform Resource Identifiers (URI's) schemes http and https. URIs and hyperlinks in HTML documents form inter-linked hypertext documents.

#### HTTP Verbs

> HTTP defines methods (sometimes referred to as verbs) to indicate the desired action to be performed on the identified resource. What this resource represents, whether pre-existing data or data that is generated dynamically, depends on the implementation of the server.

#### HTTP Status Codes

The HTTP response code designates the status of a particular operation for example a 200 for a successful GET request
The status code is the first line of the HTTP response and is called the status line and includes a numeric status code (such as "404") and a textual phrase (such as "Not Found").

#### HTTP Request Message

The request message consists of the following:

* A request line (e.g., GET /user/1 HTTP/1.1, which requests a resource called /user/1 from the server).

* Request header fields (e.g., Accept-Language: en).

* An empty line.

* An optional message body.

#### HTTP Response Message

The response message consists of the following:

* A status line which includes the status code and reason message (for example HTTP/1.1 200 OK, which indicates that the client's request succeeded).

* Response header fields (e.g., Content-Type: application/json).

* An empty line.

* An optional message body.

#### HTTP session state

> HTTP is a stateless protocol. A stateless protocol does not require the HTTP server to retain information or status about each user for the duration of multiple requests. However, some web applications implement states or server side sessions using for instance HTTP cookies or hidden variables within web forms.

#### HTTP encryption

The most popular way of establishing an encrypted HTTP connection is HTTP Secure or `HTTPS`

There are two other methods for establishing an encrypted HTTP connection: 

* Secure Hypertext Transfer Protocol

* HTTP/1.1 Upgrade header to specify an upgrade to TLS.

**There is limited browser support for these so beware**

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [Richardson Maturity Model](./maturity-model.md) | [HTTP Verbs](./http-verbs.md) →
