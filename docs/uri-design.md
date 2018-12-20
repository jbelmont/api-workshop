API Workshop - Uri Design

## Sections:

* [Bread Crumb Navigation](#bread-crumb-navigation)

#### Uniform Resource Identifiers (URI)

[RFC 2396](https://tools.ietf.org/html/rfc2396)

> A Uniform Resource Identifier (URI) is a compact string of characters for identifying an abstract or physical resource.

##### Overview of URI

###### Uniform

[RFC 2396 Section 1.1](https://tools.ietf.org/html/rfc2396#section-1.1)

Uniformity provides several benefits:

* Uniformity allows different types of resource identifiers to be used in the same context, even when the mechanisms used to access those resources may differ

* Uniformity allows uniform semantic interpretation of common syntactic conventions across different types of resource identifiers;

* Uniformity allows introduction of new types of resource identifiers without interfering with the way that existing identifiers are used

* Uniformity allows the identifiers to be reused in many different contexts, thus permitting new applications or protocols to leverage a pre-existing, large, and widely-used set of resource identifiers.

###### Resource

A resource can be anything that has identity.

* Familiar examples include an electronic document, an image, a service (e.g., "today's weather report for Los Angeles"), and a collection of other resources.

* Not all resources are network "retrievable"; e.g., human beings, corporations, and bound books in a library can also be considered resources.

* The resource is the conceptual mapping to an entity or set of entities, not necessarily the entity which corresponds to that mapping at any particular instance in time.

* Thus, a resource can remain constant even when its content---the entities to which it currently corresponds---changes over time, provided that the conceptual mapping is not changed in the process.

###### Identifier

* An identifier is an object that can act as a reference to something that has identity.

* In the case of URI, the object is a sequence of characters with a restricted syntax.

##### Example URI

* `ftp://ftp.is.co.za/rfc/rfc1808.txt`

* `mailto:marcelbelmont@gmail.com`

* `news:comp.infosystems.www.servers.unix`

* `https://developer.mozilla.org/en-US/`

#### URI Characters and Escape Sequences

[RFC 2396 Section 2](https://tools.ietf.org/html/rfc2396#section-2)

> URI consist of a restricted set of characters, primarily chosen to aid transcribability and usability both in computer systems and in non-computer communications. Characters used conventionally as delimiters around URI were excluded.  The restricted set of characters consists of digits, letters, and a few graphic symbols were chosen from those common to most of the character encodings and input facilities available to Internet users.

##### Reserved Characters

[RFC 2396 Section 2.2](https://tools.ietf.org/html/rfc2396#section-2.2)

Many URI include components consisting of or delimited by, certain special characters.

These characters are called "reserved", since their usage within the URI component is limited to their reserved purpose.

If the data for a URI component would conflict with the reserved purpose, then the conflicting data must be escaped before forming the URI.

Reserved characters:

* `;`
* `/`
* `?`
* `:`
* `@`
* `&`
* `=`
* `+`
* `$`
* `,`

* Characters in the "reserved" set are not reserved in all contexts.

* The set of characters actually reserved within any given URI component is defined by that component.

* In general, a character is reserved if the semantics of the URI changes if the character is replaced with its escaped US-ASCII encoding.

When a character from the reserved set (a "reserved character") has special meaning (a "reserved purpose") in a certain context, and a URI scheme says that it is necessary to use that character for some other purpose, then the character must be percent-encoded.

Percent-encoding a reserved character involves converting the character to its corresponding byte value in ASCII and then representing that value as a pair of hexadecimal digits.

##### Unreserved Characters

[RFC 2396 Section 2.3](https://tools.ietf.org/html/rfc2396#section-2.3)

Data characters that are allowed in a URI but do not have a reserved purpose are called unreserved.

These include upper and lower case letters, decimal digits, and a limited set of punctuation marks and symbols.

Unreserved characters:

* `-`
* `_`
* `.`
* `!`
* `~`
* `*`
*  `'`
* `(`
* `)`

Unreserved characters can be escaped without changing the semantics of the URI, but this should not be done unless the URI is being used in a context that does not allow the unescaped character to appear.

**Characters from the unreserved set never need to be percent-encoded.**

##### Escape Sequences

[RFC 2396 Section 2.4](https://tools.ietf.org/html/rfc2396#section-2.4)

* Data must be escaped if it does not have a representation using an unreserved character

* this includes data that does not correspond to a printable character of the US-ASCII coded character set, or that corresponds to any US-ASCII character that is disallowed, as explained below

An escaped octet is encoded as a character triplet, consisting of the percent character "%" followed by the two hexadecimal digits representing the octet code.

For example, "%20" is the escaped encoding for the US-ASCII space character.

A URI is always in an "escaped" form, since escaping or unescaping a completed URI might change its semantics.

Normally, the only time escape encodings can safely be made is when the URI is being created from its component parts

Each component may have its own set of characters that are reserved, so only the mechanism responsible for generating or interpreting that component can determine whether or not escaping a character will change its semantics.

Likewise, a URI must be separated into its components before the escaped characters within those components can be safely decoded.

##### Common Characters that are Percent Encoded

| Character | Percent Encoding (PE) | Character | PE | Character | PE | Character | PE |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Newline | %0A | &#32; | %20 | &#34; | %22 | % | %25 |
| &#8211; | %2D | &#46; | %2E | < | %3C | > | %3E |
| \ | %5C | &#94; | %5E | _ | %5F | &#96; | %60 |
| \{ | %7B | &#124; | %7C | \} | %7D | \~ | %7E |

#### URI Syntactic Components

[RFC 2396 Section 3](https://tools.ietf.org/html/rfc2396#section-3)

In general, absolute URI are written as follows:

`<scheme>:<scheme-specific-part>`

This "generic URI" syntax consists of a sequence of four main components:

`<scheme>://<authority><path>?<query>`

Scheme names consist of a sequence of characters beginning with a lower case letter and followed by any combination of lower case letters, digits, plus ("+"), period ("."), or hyphen ("-").

Scheme ==> `alpha *( alpha | digit | "+" | "-" | "." )`

##### Authority Component

[RFC 2396 Section 3.2](https://tools.ietf.org/html/rfc2396#section-3.2)

An authority part, comprising:

* An optional authentication section of a user name and password, separated by a colon, followed by an at symbol (@)

* A "host", consisting of either a registered name (including but not limited to a hostname), or an IP address.
  * IPv4 addresses must be in dot-decimal notation, and IPv6 addresses must be enclosed in brackets ([ ]).
  * An optional port number, separated from the hostname by a colon

Some URL schemes use the format "user:password" in the userinfo field.

This practice is NOT RECOMMENDED, because the passing of authentication information in clear text (such as URI) has proven to be a security risk in almost every case where it has been used.

###### Authority Component Example

`http://john:secret@rambo.net:3000/path/data`

*The part `john:secret@rambo.net` is the authority component*

##### Path Component

[RFC 2396 Section 3.3](https://tools.ietf.org/html/rfc2396#section-3.3)

The path component contains data, specific to the authority (or the scheme if there is no authority component), identifying the resource within the scope of that scheme and authority.

* The path may consist of a sequence of path segments separated by a single slash "/" character.  Within a path segment, the characters "/", ";", "=", and "?" are reserved.

* Each path segment may include a sequence of parameters, indicated by the semicolon ";" character.

* The parameters are not significant to the parsing of relative references.

###### Path Component Example

`https://www.marcelbelmont.com/post/api-testing`

The part `/post/api-testing` is the path component

##### Query Component

[RFC 2396 Section 3.4](https://tools.ietf.org/html/rfc2396#section-3.4)

The query component is a string of information to be interpreted by the resource.

*Within a query component, the characters ";", "/", "?", ":", "@", "&", "=", "+", ",", and "$" are reserved.*

###### Query Component Example

`https://someuser.auth0.com/api/v2/authorize?audience=https://someuser.auth0.com/api/v2/&scope=openid%20profile%20 email&response_type=code&client_id=abcd13555&redirect_uri=https://task.io/callback`

The part `?audience=https://someuser.auth0.com/api/v2/&scope=openid%20profile%20 email&response_type=code&client_id=abcd13555&redirect_uri=https://task.io/callback` is the query component part.

#### URI References

[RFC 2396 Section 4](https://tools.ietf.org/html/rfc2396#section-4)

* The term "URI-reference" is used here to denote the common usage of a resource identifier.

* A URI reference may be absolute or relative, and may have additional information attached in the form of a fragment identifier.

* However, "the URI" that results from such a reference includes only the absolute URI after the fragment identifier (if any) is removed and after any relative URI is resolved to its absolute form.

##### Fragment Identifier

[RFC 2396 Section 4.1](https://tools.ietf.org/html/rfc2396#section-4.1)

When a URI reference is used to perform a retrieval action on the identified resource, the optional fragment identifier, separated from the URI by a crosshatch ("#") character, consists of additional reference information to be interpreted by the user agent after the retrieval action has been successfully completed.

As such, it is not part of a URI, but is often used in conjunction with a URI.

###### Fragment Example

`https://github.com/jbelmont/api-workshop#workshop-details`

The Fragment part is `#workshop-details`

##### Same-document References

[RFC 2396 Section 4.2](https://tools.ietf.org/html/rfc2396#section-4.2)

A URI reference that does not contain a URI is a reference to the current document.  In other words, an empty URI reference within a document is interpreted as a reference to the start of that document, and a reference containing only a fragment identifier is a reference to the identified fragment of that document.

Traversal of such a reference should not result in an additional retrieval action.

##### Parsing a URI Reference

[RFC 2396 Section 4.3](https://tools.ietf.org/html/rfc2396#section-4.3)

A URI reference is typically parsed according to the four main components and fragment identifier in order to determine what components are present and whether the reference is relative or absolute.

The individual components are then parsed for their subparts and, if not opaque, to verify their validity.

#### Relative URI References

[RFC 2396 Section 5](https://tools.ietf.org/html/rfc2396#section-5)

> It is often the case that a group or "tree" of documents has been constructed to serve a common purpose; the vast majority of URI in these documents point to resources within the tree rather than outside of it.  Similarly, documents located at a particular site are much more likely to refer to other resources at that site than to resources at remote sites.

Relative addressing of URI allows document trees to be partially independent of their location and access scheme.

For instance, it is possible for a single set of hypertext documents to be simultaneously accessible and traversable via each of the "file", "http", and "ftp" schemes if the documents refer to each other using relative URI.

##### Establishing a Base URI

* The term "relative URI" implies that there exists some absolute "base URI" against which the relative reference is applied.

* Indeed, the base URI is necessary to define the semantics of any relative URI reference; without it, a relative reference is meaningless.

* In order for relative URI to be usable within a document, the base URI of that document must be known to the parser.

Please read sections: [5.1.1](https://tools.ietf.org/html/rfc2396#section-5.1.1), [5.1.2](https://tools.ietf.org/html/rfc2396#section-5.1.2), [5.1.3](https://tools.ietf.org/html/rfc2396#section-5.1.3), [5.1.4](https://tools.ietf.org/html/rfc2396#section-5.1.4) for more details.

##### Resolving Relative References to Absolute Form

**This section of RFC discusses an example algorithm for resolving URI references that might be relative to a given base URI.**

#### URI Normalization and Equivalence

[RFC 2396 Section 6](https://tools.ietf.org/html/rfc2396#section-6)

In many cases, different URI strings may actually identify the identical resource.

For example, the host names used in URL are actually case insensitive, and the URL <http://www.XEROX.com> is equivalent to <http://www.xerox.com>.

In general, the rules for equivalence and definition of a normal form, if any, are scheme dependent.

When a scheme uses elements of the common syntax, it will also use the common syntax equivalence rules, namely that the scheme and hostname are case insensitive and a URL with an explicit ":port", where the port is the default for the scheme, is equivalent to one where the port is elided.

#### Security Considerations

[RFC 2396 Section 7](https://tools.ietf.org/html/rfc2396#section-7)

A URI does not in itself pose a security threat.

Users should beware that there is no general guarantee that a URL, which at one time located a given resource, will continue to do so.

Nor is there any guarantee that a URL will not locate a different resource at some later point in time, due to the lack of any constraint on how a given authority apportions its namespace.

Such a guarantee can only be obtained from the person(s) controlling that namespace and the resource in question.

A specific URI scheme may include additional semantics, such as name persistence, if those semantics are required of all naming authorities for that scheme.

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [HTTP Request and Response](./request-response.md) | [Query String](./query-string.md) →
