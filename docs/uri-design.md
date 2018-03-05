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

| Character | Percent Encoding | Character | Percent Encoding | Character | Percent Encoding | Character | Percent Encoding |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Newline (\n) | %0A | Space (" ") | %20 | Double Quote (") | %22 | Percent (%) | %25 |
| Dash (-) | %2D | Period (.) | %2E | < | %3C | > | %3E |
| \ | %5C | Caret (^) | %5E | _ | %5F | \` | %60 |
| { | %7B | \| | %7C | } | %7D | ~ | %7E |

#### URI Syntactic Components

content

#### URI References

content

#### Relative URI References

content

#### URI Normalization and Equivalence

content

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [HTTP Request and Response](./request-response.md) | [Query String](./query-string.md) →
