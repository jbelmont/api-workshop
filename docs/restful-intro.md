API Workshop - Restful API Introduction

## Sections:

* [Bread Crumb Navigation](#bread-crumb-navigation)

#### Soap 

Simple Object Access Protocol [SOAP](https://en.wikipedia.org/wiki/SOAP) is a Protocol

Soap uses (Extensible Markup Language (XML)) [XML](https://en.wikipedia.org/wiki/XML) as its message format

Hypertext Transfer Protocol [HTTP](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol) is the application protocol most used in SOAP services although Soap is not limited to HTTP.

SOAP has three major characteristics:

* extensibility (security and WS-Addressing are among the extensions under development)

* neutrality because it can operate over any protocol
  * [HTTP](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol)
  * Simple Mail Transfer [SMTP](https://en.wikipedia.org/wiki/Simple_Mail_Transfer_Protocol)
  * Transmission Control Protocol [TCP](https://en.wikipedia.org/wiki/Transmission_Control_Protocol)
  * User Datagram Protocol [UDP](https://en.wikipedia.org/wiki/User_Datagram_Protocol)

* independence (SOAP allows for any programming model)

The SOAP architecture consists of several layers of specifications for:

* message format
* Message Exchange Patterns (MEP)
* underlying transport protocol bindings
* message processing models
* protocol extensibility

#### Disadvantages of Soap

Standard implementations of SOAP utilize XML as the messaging format

Soap is a verbose protocol and it is slower to parse XML than [JSON](https://en.wikipedia.org/wiki/JSON)

#### Differences between Soap and REST

Soap is a protocol while Representational State Transfer (REST) is more of an architectural style

SOAP is a well-developed protocol used in the Web industry and is standardized by the World Wide Web Consortium [W3C](https://en.wikipedia.org/wiki/World_Wide_Web_Consortium).

Rest is the culmination [Roy Fieldings](https://en.wikipedia.org/wiki/Roy_Fielding) PHd Dissertation research

Rest does not a dictate a messaging format so you can in theory use xml, json, and text as message formats.

Modern Restful APIs use JSON as the message format

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [API Summary and Description](./api-summary.md) | [Roy Fieldings Dissertation on REST](./rest-dissertation.md) →
