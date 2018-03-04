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
