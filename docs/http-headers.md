API Workshop - HTTP Headers

## Sections:

* [Bread Crumb Navigation](#bread-crumb-navigation)

#### HTTP Header Fields

[HTTP Header Fields](https://en.wikipedia.org/wiki/List_of_HTTP_header_fields)

> HTTP header fields are components of the header section of request and response messages in the Hypertext Transfer Protocol (HTTP). They define the operating parameters of an HTTP transaction.

The format of HTTP Headers follow that which is laid out in [RFC 822](https://tools.ietf.org/html/rfc822)

Each header field consists of a name followed by a colon (":") and the field value. 

According to [RFC 2616 Section 4.2](https://tools.ietf.org/html/rfc2616#section-4.2) a field value MAY be preceded by any amount of Linear White Space (LWS), though a single space (SP) is preferred. 

Header fields can be extended over multiple lines by preceding each extra line with at least one SP or HT. 

```http
message-header = field-name ":" [ field-value ]

field-name     = token
       
field-value    = *( field-content | LWS )

field-content  = <the OCTETs making up the field-value and consisting of either *TEXT or 
combinations of token, separators, and quoted-string>
```

#### Request Header Fields

[RFC 2616 Section 5.3](https://tools.ietf.org/html/rfc2616#section-5.3)

The request-header fields allow the client to pass additional information about the request, and about the client itself, to the server. 

These fields act as request modifiers, with semantics equivalent to the parameters on a programming language method invocation.

> Request-header field names can be extended reliably only in combination with a change in the protocol version. However, new or experimental header fields MAY be given the semantics of request-header fields if all parties in the communication recognize them to be request-header fields. 

Unrecognized header fields are treated as entity-header fields.

An entity header is an HTTP header describing the content of the body of the message. 

Entity headers are used in both, HTTP requests and responses. 

Headers like Content-Length, Content-Language, Content-Encoding are entity headers.

#### List of Request Header Fields

| Header Field Name |	Description |	Example	| RFC Reference |
| --- | --- | --- | --- |
| Accept | The Accept request-header field can be used to specify certain media types which are acceptable for the response. Accept headers can be used to indicate that the request is specifically limited to a small set of desired types, as in the case of a request for an in-line image. | `Accept: text/plain` | [RFC 2616 Section 14.1](https://tools.ietf.org/html/rfc2616#section-14.1) |
| Accept-Charset | The Accept-Charset request-header field can be used to indicate what character sets are acceptable for the response. This field allows clients capable of understanding more comprehensive or special-purpose character sets to signal that capability to a server which is capable of representing documents in those character sets. | `Accept-Charset: utf-8` | [RFC 2616 Section](https://tools.ietf.org/html/rfc2616#section-14.2) |
| Accept-Encoding | The Accept-Encoding request-header field is similar to Accept, but restricts the content-codings that are acceptable in the response. | `Accept-Encoding: gzip, deflate` | [RFC 2616 Section](https://tools.ietf.org/html/rfc2616#section-14.3) |
| Accept-Language | The Accept-Language request-header field is similar to Accept, but restricts the set of natural languages that are preferred as a response to the request. | `Accept-Language: en-US` | [RFC 2616 Section 14.4](https://tools.ietf.org/html/rfc2616#section-14.4) |
| Authorization | A user agent that wishes to authenticate itself with a server--usually, but not necessarily, after receiving a 401 response--does so by including an Authorization request-header field with the request. The Authorization field value consists of credentials containing the authentication information of the user agent for the realm of the resource being requested. | `Authorization: Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==` | [RFC 2616 Section 14.8](https://tools.ietf.org/html/rfc2616#section-14.8) |
| Expect | The Expect request-header field is used to indicate that particular server behaviors are required by the client. | `Expect: 100-continue` | [RFC 2616 14.20](https://tools.ietf.org/html/rfc2616#section-14.20)
| From | The From request-header field, if given, SHOULD contain an Internete-mail address for the human user who controls the requesting user agent. The address SHOULD be machine-usable, as defined by "mailbox" in [RFC 822](https://tools.ietf.org/html/rfc822) as updated by [RFC 1123](https://tools.ietf.org/html/rfc1123.html) | `From: user@example.com` | [RFC 2622 Section 14.22](https://tools.ietf.org/html/rfc2616#section-14.22) |
| Host | The Host request-header field specifies the Internet host and port number of the resource being requested, as obtained from the original URI given by the user or referring resource. The Host field value MUST represent the naming authority of the origin server or gateway given by the original URL. This allows the origin server or gateway to differentiate between internally-ambiguous URLs, such as the root "/" URL of a server for multiple host names on a single IP address. | `Host: www.owasp.org` | [RFC 2616 Section 14.23](https://tools.ietf.org/html/rfc2616#section-14.23) |
| If-Match | The If-Match request-header field is used with a method to make it conditional. A client that has one or more entities previously obtained from the resource can verify that one of those entities is current by including a list of their associated entity tags in the If-Match header field. | `If-Match: "737060cd8c284d8af7ad3082f209582d"` | [RFC 2616 Section 14.24](https://tools.ietf.org/html/rfc2616#section-14.24) |
| If-Modified-Since | The If-Modified-Since request-header field is used with a method to make it conditional: if the requested variant has not been modified since the time specified in this field, an entity will not be returned from the server; instead, a 304 (not modified) response will be returned without any message-body. | `If-Modified-Since: Sat, 29 Oct 1994 19:43:31 GMT` | [RFC 2616 Section 14.25](https://tools.ietf.org/html/rfc2616#section-14.25) |
| If-None-Match | The If-None-Match request-header field is used with a method to make it conditional. A client that has one or more entities previously obtained from the resource can verify that none of those entities is current by including a list of their associated entity tags in the If-None-Match header field. The purpose of this feature is to allow efficient updates of cached information with a minimum amount of transaction overhead. It is also used to prevent a method (e.g. PUT) from inadvertently modifying an existing resource when the client believes that the resource does not exist. | `If-None-Match: "737060cd8c284d8af7ad3082f209582d"` | [RFC 2616 Section 14.26](https://tools.ietf.org/html/rfc2616#section-14.26) |
| If-Range | If a client has a partial copy of an entity in its cache, and wishes to have an up-to-date copy of the entire entity in its cache, it could use the Range request-header with a conditional GET (using either or both of If-Unmodified-Since and If-Match.) However, if the condition fails because the entity has been modified, the client would then have to make a second request to obtain the entire current entity-body. | `If-Range: "737060cd8c284d8af7ad3082f209582d"` | [RFC 2616 Section 14.27](https://tools.ietf.org/html/rfc2616#section-14.27) |
| If-Unmodified-Since | The If-Unmodified-Since request-header field is used with a method to make it conditional. If the requested resource has not been modified since the time specified in this field, the server SHOULD perform the requested operation as if the If-Unmodified-Since header were not present. | `If-Unmodified-Since: Sat, 29 Oct 1994 19:43:31 GMT` | [RFC 2616 Section 14.28](https://tools.ietf.org/html/rfc2616#section-14.28) |
| Max-Forwards | The Max-Forwards request-header field provides a mechanism with the TRACE and OPTIONS  methods to limit the number of proxies or gateways that can forward the request to the next inbound server. This can be useful when the client is attempting to trace a request chain which appears to be failing or looping in mid-chain. | `Max-Forwards: 10` | [RFC 2616 Section 14.31](https://tools.ietf.org/html/rfc2616#section-14.31) |
| Proxy-Authorization | The Proxy-Authorization request-header field allows the client to identify itself (or its user) to a proxy which requires authentication. The Proxy-Authorization field value consists of credentials containing the authentication information of the user agent for the proxy and/or realm of the resource being requested. | `Proxy-Authorization: Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==` | [RFC 2616 Section 14.34](https://tools.ietf.org/html/rfc2616#section-14.34) |
| Range | Request only part of an entity. Bytes are numbered from 0 | `Range: bytes=500-999` | [RFC 2616 Section 14.35](https://tools.ietf.org/html/rfc2616#section-14.35) |
| Referer | The Referer[sic] request-header field allows the client to specify, for the server's benefit, the address (URI) of the resource from which the Request-URI was obtained (the "referrer", although the header field is misspelled.) The Referer request-header allows a server to generate lists of back-links to resources for interest, logging, optimized caching, etc. It also allows obsolete or mistyped links to be traced for maintenance. The Referer field MUST NOT be sent if the Request-URI was obtained from a source that does not have its own URI, such as input from the user keyboard. | `Referer: https://www.youtube.com/embed/videoseries?list=PLpr-xdpM8wG9g8l21uzfgyx4xAJ9aY7e7&` | [RFC 2616 Section 14.36](https://tools.ietf.org/html/rfc2616#section-14.36) |
| TE | The TE request-header field indicates what extension transfer-codings it is willing to accept in the response and whether or not it is willing to accept trailer fields in a chunked transfer-coding. Its value may consist of the keyword "trailers" and/or a comma-separated list of extension transfer-coding names with optional accept parameters | `TE: trailers, deflate` | [RFC 2616 Section 14.39](https://tools.ietf.org/html/rfc2616#section-14.39) |
| User-Agent | The User-Agent request-header field contains information about the user agent originating the request. This is for statistical purposes, the tracing of protocol violations, and automated recognition of user agents for the sake of tailoring responses to avoid particular user agent limitations. User agents SHOULD include this field with requests. The field can contain multiple product tokens (section 3.8) and comments identifying the agent and any subproducts which form a significant part of the user agent. By convention, the product tokens are listed in order of their significance for identifying the application. | `User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:12.0) Gecko/20100101 Firefox/12.0` | [RFC 2616 Section 14.43](https://tools.ietf.org/html/rfc2616#section-14.43) |


Reference [IANA Site for full list of Headers](https://www.iana.org/assignments/message-headers/message-headers.xml#perm-headers)

#### Response Header Fields

[RFC 2616 Section 6.2](https://tools.ietf.org/html/rfc2616#section-6.2)

The response-header fields allow the server to pass additional information about the response which cannot be placed in the Status-Line. 

Response header fields give information about the server and about further access to the resource identified by the Request-URI.

> Response-header field names can be extended reliably only in combination with a change in the protocol version. However, new or experimental header fields MAY be given the semantics of response-header fields if all parties in the communication recognize them to be response-header fields. Unrecognized header fields are treated as entity-header fields.

#### List of Response Header Fields

| Header Field Name |	Description |	Example	| RFC Reference |
| --- | --- | --- | --- |
| Accept-Ranges | The Accept-Ranges response-header field allows the server to indicate its acceptance of range requests for a resource | `Accept-Ranges: bytes` | [RFC 2616 Section 14.5](https://tools.ietf.org/html/rfc2616#section-14.5) |
| Age | The Age response-header field conveys the sender's estimate of the amount of time since the response (or its revalidation) was generated at the origin server. A cached response is "fresh" if its age does not exceed its freshness lifetime. Age values are calculated as specified in [RFC 2616 Section 13.2.3](https://tools.ietf.org/html/rfc2616#section-13.2.3) | `Age: 24` | [RFC 2616 Section 14.6](https://tools.ietf.org/html/rfc2616#section-14.6) |
| ETag | The ETag response-header field provides the current value of the entity tag for the requested variant. The headers used with entity tags are described in sections 14.24, 14.26 and 14.44. The entity tag MAY be used for comparison with other entities from the same resource [RFC 2616 Section 13.3.3](https://tools.ietf.org/html/rfc2616#section-13.3.3) | `ETag: "33a64df551425fcc55e4d42a148795d9f25f89d4"` | [RFC 2616 Section 14.19](https://tools.ietf.org/html/rfc2616#section-14.19) |
| Location | The Location response-header field is used to redirect the recipient to a location other than the Request-URI for completion of the request or identification of a new resource. For 201 (Created) responses, the Location is that of the new resource which was created by the request. For 3xx responses, the location SHOULD indicate the server's preferred URI for automatic redirection to the resource. The field value consists of a single absolute URI. | `Location: /index.html` | [RFC 2616 Section 14.30](https://tools.ietf.org/html/rfc2616#section-14.30) |
| Proxy-Authenticate | The Proxy-Authenticate response-header field MUST be included as part of a 407 (Proxy Authentication Required) response. The field value consists of a challenge that indicates the authentication scheme and parameters applicable to the proxy for this Request-URI. | `Proxy-Authorization: Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ==` | [RFC 2616 Section 14.33](https://tools.ietf.org/html/rfc2616#section-14.33) |
| Retry-After | The Retry-After response-header field can be used with a 503 (Service Unavailable) response to indicate how long the service is expected to be unavailable to the requesting client. This field MAY also be used with any 3xx (Redirection) response to indicate the minimum time the user-agent is asked wait before issuing the redirected request. The value of this field can be either an HTTP-date or an integer number of seconds (in decimal) after the time of the response. | `Retry-After: Wed, 21 Oct 2015 07:28:00 GMT` | [RFC 2616 Section 14.37](https://tools.ietf.org/html/rfc2616#section-14.37) |
| Server | The Server response-header field contains information about the software used by the origin server to handle the request. The field can contain multiple product tokens [RFC 2616 Section 3.8](https://tools.ietf.org/html/rfc2616#section-3.8) and comments identifying the server and any significant subproducts. The product tokens are listed in order of their significance for identifying the application. | `Server: Apache/2.4.1 (Unix)` | [RFC 2616 Section 14.38](https://tools.ietf.org/html/rfc2616#section-14.38) |
| Vary | The Vary field value indicates the set of request-header fields that fully determines, while the response is fresh, whether a cache is permitted to use the response to reply to a subsequent request without revalidation. For uncacheable or stale responses, the Vary field value advises the user agent about the criteria that were used to select the representation. A Vary field value of "*" implies that a cache cannot determine from the request headers of a subsequent request whether this response is the appropriate representation. | `Vary: User-Agent` | [RFC 2616 Section 14.44](https://tools.ietf.org/html/rfc2616#section-14.44) |
| WWW-Authenticate | The WWW-Authenticate response-header field MUST be included in 401 (Unauthorized) response messages. The field value consists of at least one challenge that indicates the authentication scheme(s) and parameters applicable to the Request-URI. | `WWW-Authenticate: Basic realm="Access to the staging site", charset="UTF-8"` | [RFC 2616 Section 14.47](https://tools.ietf.org/html/rfc2616#section-14.47) |

Reference [IANA Site for full list of Headers](https://www.iana.org/assignments/message-headers/message-headers.xml#perm-headers)

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [HTTP Status Codes](./status-codes.md) | [HTTP Request and Response](./request-response.md) →
