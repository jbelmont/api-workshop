API Workshop - HTTP Request and Response

## Sections:

* [Bread Crumb Navigation](#bread-crumb-navigation)

#### HTTP Request

[RFC 2616 Section 5](https://tools.ietf.org/html/rfc2616#section-5)

> A request message from a client to a server includes, within the first line of that message, the method to be applied to the resource, the identifier of the resource, and the protocol version in use.

##### Request Line

[RFC 2616 Section 5.1](https://tools.ietf.org/html/rfc2616#section-5.1)

> The Request-Line begins with a method token, followed by the Request-URI and the protocol version, and ending with CRLF. The elements are separated by space characters. No Carriage Return (CR) or Line Feed (LF) is allowed except in the final CRLF sequence.

###### HTTP Request Methods

The request line has 3 parts:

* The HTTP Request Method
    * [GET](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
    * [POST](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
    * [HEAD](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
    * [OPTIONS](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
    * [PUT](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
    * [DELETE](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
    * [TRACE](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
    * [CONNECT](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods)
    
###### Request Uniform Resource Identifier (URI)

[Request URI](https://tools.ietf.org/html/rfc2616#section-5.1.2)

> URIs have been known by many names: WWW addresses, Universal Document Identifiers, Universal Resource Identifiers [3], and finally the combination of Uniform Resource Locators (URL) and Names (URN). As far as HTTP is concerned, Uniform Resource Identifiers are simply formatted strings which identify-via name, location, or any other characteristic--a resource.

> The Request-URI is a Uniform Resource Identifier and identifies the resource upon which to apply the request.

[Uniform Resource Identifier](https://en.wikipedia.org/wiki/Uniform_Resource_Identifier)

There are 4 options for a Request URI:

* `*` 
    * The asterisk "*" means that the request does not apply to a particular resource, but to the server itself, and is only allowed when the method used does not necessarily apply to a resource
    * Example: `OPTIONS * HTTP/1.1`
    
* AbsoluteURI
    * The absoluteURI form is REQUIRED when the request is being made to a proxy. The proxy is requested to forward the request or service it from a valid cache, and return the response. Note that the proxy MAY forward the request on to another proxy or directly to the server specified by the absoluteURI.
    * Examples:
      * https://example.org/absolute/URI/with/absolute/path/to/resource.txt
      * https://example.org/absolute/URI/with/absolute/path/to/resource
      * ftp://example.org/resource.txt
      * urn:ISSN:1535-3613
      
* AbsolutePath
    * The most common form of Request-URI is that used to identify a resource on an origin server or gateway. In this case the absolute path of the URI MUST be transmitted (see section 3.2.1, abs_path) as the Request-URI, and the network location of the URI (authority) MUST be transmitted in a Host header field.
    * Example: `https://example.org/absolute/URI/with/absolute/path/to/resource.txt`
    
* Authority
    * An authority part, comprising an optional authentication section of a user name and password, separated by a colon, followed by an at symbol (@) a "host", consisting of either a registered name (including but not limited to a hostname), or an IP address. IPv4 addresses must be in dot-decimal notation, and IPv6 addresses must be enclosed in brackets ([ ]). An optional port number, separated from the hostname by a colon
    * Example: `http://user:password@example.org:5071`
    
#### General Header Fields

> There are a few header fields which have general applicability for both request and response messages, but which do not apply to the entity being transferred. These header fields apply only to the message being transmitted.

| Header Field Name |	Description |	Example	| RFC Reference |
| --- | --- | --- | --- |
| Cache-Control | The Cache-Control general-header field is used to specify directives that MUST be obeyed by all caching mechanisms along the request/response chain. The directives specify behavior intended to prevent caches from adversely interfering with the request or response. These directives typically override the default caching algorithms. Cache directives are unidirectional in that the presence of a directive in a request does not imply that the same directive is to be given in the response. | `Cache-Control: no-cache` | [RFC 2616 Section 14.9](https://tools.ietf.org/html/rfc2616#section-14.9) |
| Connection | The Connection general-header field allows the sender to specify options that are desired for that particular connection and MUST NOT be communicated by proxies over further connections. | `Connection: keep-alive` | [RFC 2616 Section 14.10](https://tools.ietf.org/html/rfc2616#section-14.10) |
| Date | The Date general-header field represents the date and time at which the message was originated, having the same semantics as orig-date in [RFC 822](https://www.ietf.org/rfc/rfc822.txt). The field value is an HTTP-date and MUST be sent in [RFC 1123 Section 5](https://tools.ietf.org/html/rfc1123#section-5) -date format. | `Date: Tue, 15 Nov 2015 08:12:31 GMT` | [RFC 2616 Section 14.18](https://tools.ietf.org/html/rfc2616#section-14.18) |
| Pragma | The Pragma general-header field is used to include implementation-specific directives that might apply to any recipient along the request/response chain. All pragma directives specify optional behavior from the viewpoint of the protocol; however, some systems MAY require that behavior be consistent with the directives. | `Pragma: no-cache` | [RFC 2616 Section 14.32](https://tools.ietf.org/html/rfc2616#section-14.32) |
| Trailer | The Trailer general field value indicates that the given set of header fields is present in the trailer of a message encoded with chunked transfer-coding. | `Trailer: header-names` | [RFC 2616 Section 14.40](https://tools.ietf.org/html/rfc2616#section-14.40) |
| Transfer-Encoding | The Transfer-Encoding general-header field indicates what (if any) type of transformation has been applied to the message body in order to safely transfer it between the sender and the recipient. This differs from the content-coding in that the transfer-coding is a property of the message, not of the entity. | `Transfer-Encoding: chunked` | [RFC 2616 Section 14.41](https://tools.ietf.org/html/rfc2616#section-14.41) |
| Upgrade | The Upgrade general-header allows the client to specify what additional communication protocols it supports and would like to use if the server finds it appropriate to switch protocols. The server MUST use the Upgrade header field within a 101 (Switching Protocols) response to indicate which protocol(s) are being switched. | `Upgrade: HTTP/2.0, SHTTP/1.3, IRC/6.9, RTA/x11` | [RFC 2616 Section 14.42](https://tools.ietf.org/html/rfc2616#section-14.42) |
| Via | The Via general-header field MUST be used by gateways and proxies to indicate the intermediate protocols and recipients between the user agent and the server on requests, and between the origin server and the client on responses. It is analogous to the "Received" field of [RFC 822](https://www.ietf.org/rfc/rfc822.txt) and is intended to be used for tracking message forwards, avoiding request loops, and identifying the protocol capabilities of all senders along the request/response chain. | `Via: HTTP/1.1 GWA` | [RFC 2616 Section 14.45](https://tools.ietf.org/html/rfc2616#section-14.45) |
| Warning | The Warning general-header field is used to carry additional information about the status or transformation of a message which might not be reflected in the message. This information is typically used to warn about a possible lack of semantic transparency from caching operations or transformations applied to the entity body of the message. | `Warning: 110 anderson/1.3.37 "Response is stale"` | [RFC 2616 Section 14.46](https://tools.ietf.org/html/rfc2616#section-14.46) |

##### Request Header Fields

> The request-header fields allow the client to pass additional information about the request, and about the client itself, to the server. These fields act as request modifiers, with semantics equivalent to the parameters on a programming language method invocation.

**Please read over [HTTP headers Section of this workshop](https://github.com/jbelmont/api-workshop/blob/master/docs/http-headers.md)**

##### Entity-Header Fields

> Entity-header fields define metainformation about the entity-body or, if no body is present, about the resource identified by the request. Some of this metainformation is OPTIONAL; some might be REQUIRED by portions of this specification.

| Header Field Name |	Description |	Example	| RFC Reference |
| --- | --- | --- | --- |
| Allow | The Allow entity-header field lists the set of methods supported by the resource identified by the Request-URI. The purpose of this field is strictly to inform the recipient of valid methods associated with the resource. An Allow header field MUST be present in a 405 (Method Not Allowed) response | `Allow: GET, POST, HEAD` | [RFC 2616 Section 14.7](https://tools.ietf.org/html/rfc2616#section-14.7) |
| Content-Encoding | The Content-Encoding entity-header field is used as a modifier to the media-type. When present, its value indicates what additional content codings have been applied to the entity-body, and thus what decoding mechanisms must be applied in order to obtain the media-type referenced by the Content-Type header field. Content-Encoding is primarily used to allow a document to be compressed without losing the identity of its underlying media type. | `Content-Encoding: gzip` | [RFC 2616 Section 14.11](https://tools.ietf.org/html/rfc2616#section-14.11) |
| Content-Language | The Content-Language entity-header field describes the natural language(s) of the intended audience for the enclosed entity. Note that this might not be equivalent to all the languages used within the entity-body. | `Content-Language: en-US` | [RFC 2616 Section 14.12](https://tools.ietf.org/html/rfc2616#section-14.12) |
| Content-Length | The Content-Length entity-header field indicates the size of the entity-body, in decimal number of OCTETs, sent to the recipient or, in the case of the HEAD method, the size of the entity-body that would have been sent had the request been a GET. | `Content-Length: 3495` | [RFC 2616 Section 14.13](https://tools.ietf.org/html/rfc2616#section-14.13) |
| Content-Location | The Content-Location entity-header field MAY be used to supply the resource location for the entity enclosed in the message when that entity is accessible from a location separate from the requested resource's URI. A server SHOULD provide a Content-Location for the variant corresponding to the response entity; especially in the case where a resource has multiple entities associated with it, and those entities actually have separate locations by which they might be individually accessed, the server SHOULD provide a Content-Location for the particular variant which is returned. | `Content-Location: /documents/foo.txt` | [RFC 2616 Section 14.14](https://tools.ietf.org/html/rfc2616#section-14.14) |
| Content-MD5| The Content-MD5 entity-header field, as defined in [RFC 1864](https://www.ietf.org/rfc/rfc1864.txt), is an MD5 digest of the entity-body for the purpose of providing an end-to-end message integrity check (MIC) of the entity-body. (Note: a MIC is good for detecting accidental modification of the entity-body in transit, but is not proof against malicious attacks.) | `Content-MD5: md5-digest` | [RFC 2616 Section 14.15](https://tools.ietf.org/html/rfc2616#section-14.15) |
| Content-Range | The Content-Range entity-header is sent with a partial entity-body to specify where in the full entity-body the partial body should be applied. Range units are defined in [RFC 2616 Section 3.12](https://tools.ietf.org/html/rfc2616#section-3.12) | `Content-Range: bytes 200-1000/67589` | [RFC 2616 Section 14.16](https://tools.ietf.org/html/rfc2616#section-14.16) |
| Content-Type | The Content-Type entity-header field indicates the media type of the entity-body sent to the recipient or, in the case of the HEAD method, the media type that would have been sent had the request been a GET. | `Content-Type: application/json` | [RFC 2616 Section 14.17](https://tools.ietf.org/html/rfc2616#section-14.17) |
| Expires | The Expires entity-header field gives the date/time after which the response is considered stale. A stale cache entry may not normally be returned by a cache (either a proxy cache or a user agent cache) unless it is first validated with the origin server (or with an intermediate cache that has a fresh copy of the entity). | `Expires: Wed, 21 Oct 2015 07:28:00 GMT` | [RFC 2616 Section 14.21](https://tools.ietf.org/html/rfc2616#section-14.21) |
| Last-Modified | The Last-Modified entity-header field indicates the date and time at which the origin server believes the variant was last modified. | `Last-Modified: Wed, 21 Oct 2015 07:28:00 GMT` | [RFC 2616 Section 14.29](https://tools.ietf.org/html/rfc2616#section-14.29) |

`extension-header = message-header`

> The extension-header mechanism allows additional entity-header fields to be defined without changing the protocol, but these fields cannot be assumed to be recognizable by the recipient. Unrecognized header fields SHOULD be ignored by the recipient and MUST be forwarded by transparent proxies.

##### Message Body

> The message-body (if any) of an HTTP message is used to carry the entity-body associated with the request or response. The message-body differs from the entity-body only when a transfer-coding has been applied, as indicated by the Transfer-Encoding header field.

* All responses to the HEAD request method MUST NOT include a message-body, even though the presence of entity-header fields might lead one to believe they do. 
   
* All 1xx (informational), 204 (no content), and 304 (not modified) responses MUST NOT include a message-body. 

* All other responses do include a message-body, although it MAY be of zero length.

##### Example Request Message Body

`application/json` format:

```json
{
  "admin": true,
  "name": "Jean-Paul Belmont"
}
```

#### HTTP Response

After receiving and interpreting a request message, a server responds with an HTTP response message.

##### Status Line

The first line of a Response message is the Status-Line, consisting of the protocol version followed by a numeric status code and its associated textual phrase, with each element separated by space characters. No Carriage Return (CR) or Line Feed (LF) is allowed except in the final CRLF sequence.

Example: `302 Found`

##### General Header Information

See the section above [General Header Fields](#general-header-information)

##### Response Header Fields

See the HTTP Header document part of the workshop and specifically read the [Response Header Fields](https://github.com/jbelmont/api-workshop/blob/master/docs/http-headers.md#response-header-fields)

##### Entity-Header Fields

Please read the section above on [Entity-Header Fields](#entity\-header-fields)

##### Entity Body

* The entity-body (if any) sent with an HTTP request or response is in a format and encoding defined by the entity-header fields.

`entity-body    = *OCTET`

* An entity-body is only present in a message when a message-body is present. 

* The entity-body is obtained from the message-body by decoding any Transfer-Encoding that might have been applied to ensure safe and proper transfer of the message.


#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [HTTP Headers](./http-headers.md) | [URI Design](./uri-design.md) →
