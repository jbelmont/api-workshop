API Workshop - Url Structure

## Sections:

* [Bread Crumb Navigation](#bread-crumb-navigation)

#### Query Component of URL

[RFC 3986 Section 3.4](https://tools.ietf.org/html/rfc3986#section-3.4)

The query component contains non-hierarchical data that, along with data in the path component, serves to identify a resource within the scope of the URI's scheme and naming authority (if any).  

The query component is indicated by the first question mark ("?") character and terminated by a number sign ("#") character or by the end of the URI.


The characters slash ("/") and question mark ("?") may represent data within the query component.

#### Query Component Example

`http://api.walmartlabs.com/v1/search?query=chromebook&format=json&apiKey=AKey12345`

The query string here start after the `?` character

So in this query you have the following pairs that are separated by the delimiter `&`:

1. `query=chromebook`
2. `format=json`
3. `apiKey=AKey12345`

#### Result filtering, sorting & searching

content

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [URI Design](./uri-design.md) | [Endpoint Design](./endpoint-design.md) →
