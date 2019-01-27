API Workshop - OAuth
## Sections:

* [Difference between OAuth and OAuth2](#difference-between-oauth-and-oauth2)
* [OAuth Roles](#oauth-roles)
* [Protocol Flow](#protocol-flow)
* [Authorization Grant Types](#authorization-grant-types)
* [Access Tokens](#access-tokens)
* [Client Registration](#client-registration)
* [Client Types](#client-types)
* [Protocol Endpoints](#protocol-endpoints)
* [Obtaining Authorization](#obtaining-authorization)
* [Reading through RFC 6749](#reading-through-rfc-6749)
* [Bread Crumb Navigation](#bread-crumb-navigation)

#### Difference between OAuth and OAuth2

content

#### OAuth Roles

[OAuth Roles](https://tools.ietf.org/html/rfc6749#section-1.1)

OAuth defines four roles:

* resource owner
  * An entity capable of granting access to a protected resource.
  *  When the resource owner is a person, it is referred to as an end-user.

* resource server
  * The server hosting the protected resources, capable of accepting and responding to protected resource requests using access tokens.

* client
  * An application making protected resource requests on behalf of the
  resource owner and with its authorization.
  * The term "client" does not imply any particular implementation characteristics (e.g.,
  whether the application executes on a server, a desktop, or other
  devices).

* authorization server
  * The server issuing access tokens to the client after successfully authenticating the resource owner and obtaining authorization.

#### Protocol Flow

*This is not my Diagram but the diagram in RFC 6749*

![Protocol Flow](./images/protocol-flow.png)

[Protocol Flows](https://tools.ietf.org/html/rfc6749#section-1.2)

The abstract OAuth 2.0 flow illustrated in Figure 1 describes the
   interaction between the four roles and includes the following steps:

(A)  The client requests authorization from the resource owner.  The
    authorization request can be made directly to the resource owner
    (as shown), or preferably indirectly via the authorization
    server as an intermediary.

(B)  The client receives an authorization grant, which is a
    credential representing the resource owner's authorization,
    expressed using one of four grant types defined in this
    specification or using an extension grant type.  The
    authorization grant type depends on the method used by the
    client to request authorization and the types supported by the
    authorization server.

(C)  The client requests an access token by authenticating with the
    authorization server and presenting the authorization grant.

(D)  The authorization server authenticates the client and validates
    the authorization grant, and if valid, issues an access token.
(E)  The client requests the protected resource from the resource
    server and authenticates by presenting the access token.

(F)  The resource server validates the access token, and if valid,
    serves the request.

#### Authorization Grant Types


> An authorization grant is a credential representing the resource owner's authorization (to access its protected resources) used by the client to obtain an access token.

[RFC 6749](https://tools.ietf.org/html/rfc6749) defines 4 Protocol flows that are called grant types:

1. **Authorization Code**: The authorization code is obtained by using an authorization server as an intermediary between the client and resource owner.

  1. Instead of requesting authorization directly from the resource owner, the client directs the resource owner to an authorization server (via its user-agent as defined in [RFC2616](https://tools.ietf.org/html/rfc2616)), which in turn directs the resource owner back to the client with the authorization code.
  2. This is also used by mobile apps, using the Proof Key for Code Exchange (PKCE) technique.
  3. [Authorization Code](https://tools.ietf.org/html/rfc6749#section-1.3.1)

2. **Implicit Grant**: The implicit grant is a simplified authorization code flow optimized for clients implemented in a browser using a scripting language such as JavaScript.

  1. In the implicit flow, instead of issuing the client an authorization code, the client is issued an access token directly (as the result of the resource owner authorization).
  2. The grant typeis implicit, as no intermediate credentials (such as an authorization code) are issued (and later used to obtain an access token).
  3. [Implicit Grant](https://tools.ietf.org/html/rfc6749#section-1.3.2)

3. **Resource Owner Password Credentials**: The resource owner password credentials (i.e., username and password) can be used directly as an authorization grant to obtain an access
   token.

   1. The credentials should only be used when there is a high degree of trust between the resource owner and the client (e.g., the client is part of the device operating system or a highly privileged application), and when other authorization grant types are not available (such as an authorization code).
   2. Even though this grant type requires direct client access to the resource owner credentials, the resource owner credentials are used for a single request and are exchanged for an access token.
   3. This grant type can eliminate the need for the client to store the resource owner credentials for future use, by exchanging the credentials with a long-lived access token or refresh token.
   4. [Resource Owner Password Credentials](https://tools.ietf.org/html/rfc6749#section-1.3.3)

4. **Client Credentials**: The client credentials (or other forms of client authentication) can be used as an authorization grant when the authorization scope is limited to the protected resources under the control of the client, or to protected resources previously arranged with the authorization server.

  1. Client credentials are used as an authorization grant typically when the client is acting on its own behalf (the client is also the resource owner) or is requesting access to protected
   resources based on an authorization previously arranged with the authorization server.
  2. [Client Credentials](https://tools.ietf.org/html/rfc6749#section-1.3.4)

#### Access Tokens

[Access Tokens](https://tools.ietf.org/html/rfc6749#section-1.4)

> Access tokens are credentials used to access protected resources.  An access token is a string representing an authorization issued to the client.  The string is usually opaque to the client.

> Tokens represent specific scopes and durations of access, granted by the resource owner, and enforced by the resource server and authorization server.

#### Refresh Tokens

[Refresh Tokens](https://tools.ietf.org/html/rfc6749#section-1.5)

* Refresh tokens are credentials used to obtain access tokens.

* Refresh tokens are issued to the client by the authorization server and are used to obtain a new access token when the current access token becomes invalid or expires, or to obtain additional access tokens with identical or narrower scope (access tokens may have a shorter lifetime and fewer permissions than authorized by the resource owner).

* Issuing a refresh token is optional at the discretion of the authorization server.  If the authorization server issues a refresh token, it is included when issuing an access token.

#### Client Registration

[Client Registration](https://tools.ietf.org/html/rfc6749#section-2)

> Before initiating the protocol, the client registers with the authorization server.

> The means through which the client registers with the authorization server are beyond the scope of this specification but typically involve end-user interaction with an HTML registration form.

#### Client Types

[Client Types](https://tools.ietf.org/html/rfc6749#section-2.1)

* OAuth defines two client types, based on their ability to authenticate securely with the authorization server (i.e., ability to maintain the confidentiality of their client credentials).

1. *confidential*: Clients capable of maintaining the confidentiality of their credentials (e.g., client implemented on a secure server with restricted access to the client credentials), or capable of secure client authentication using other means.

2. *public*: Clients incapable of maintaining the confidentiality of their credentials (e.g., clients executing on the device used by the resource owner, such as an installed native application or a web browser-based application), and incapable of secure client authentication via any other means.

#### Protocol Endpoints

[Protocol Endpoints](https://tools.ietf.org/html/rfc6749#section-3)

*There are 3 Protocol Endpoints defined by the specification*

The authorization process utilizes two authorization server endpoints (HTTP resources):

* Authorization endpoint - used by the client to obtain authorization from the resource owner via user-agent redirection.

* Token endpoint - used by the client to exchange an authorization grant for an access token, typically with client authentication.

* Redirection endpoint - used by the authorization server to return responses containing authorization credentials to the client via the resource owner user-agent.

###### Authorization Endpoint

[Authorization Endpoint](https://tools.ietf.org/html/rfc6749#section-3.1)

> The authorization endpoint is used to interact with the resource owner and obtain an authorization grant.

> The authorization server MUST first verify the identity of the resource owner. The way in which the authorization server authenticates the resource owner (e.g., username and password login, session cookies) is beyond the scope of this specification.

*The authorization endpoint is used by the **authorization code grant type** and **implicit grant type flows**.

The client informs the authorization server of the desired grant type using the *response_type* parameter.

* response_type
  * The value MUST be one of "code" for requesting an authorization code, "token" for requesting an access token (implicit grant).

[Read RFC 6749 Sections 3.1.2 to 3.1.2.5 for more information on the Authorization Endpoint](https://tools.ietf.org/html/rfc6749#section-3.1.2)

Sample request to AUTH0 authorize endpoint:

```bash
GET https://random.auth0.com/authorize?
  audience=API_IDENTIFIER&
  scope=SCOPE&
  response_type=code&
  client_id=SomeClientId&
  redirect_uri=undefined&
  code_challenge=CODE_CHALLENGE&
  code_challenge_method=S111
```

*Sample Response*:

```http
HTTP/1.1 302 Found
Location: undefined?code=AUTHORIZATION_CODE
```

###### Token Endpoint

[Token Endpoint](https://tools.ietf.org/html/rfc6749#section-3.2)

> The token endpoint is used by the client to obtain an access token by presenting its authorization grant or refresh token.

> The token endpoint is used with every authorization grant except for the implicit grant type (since an access token is issued directly).

Sample curl request to AUTH0 token endpoint:

```bash
curl --request POST \
  --url 'https://YOUR_AUTH0_DOMAIN/oauth/token' \
  --header 'content-type: application/json' \
  --data '{"grant_type":"authorization_code","client_id": "YOUR_CLIENT_ID","client_secret": "YOUR_CLIENT_SECRET","code": "AUTHORIZATION_CODE","redirect_uri": "https://YOUR_APP/callback"'
```

*Sample Response*:

```http
HTTP/1.1 200 OK
Content-Type: application/json;charset=UTF-8
Cache-Control: no-store
Pragma: no-cache

{
  "access_token":"2YotnFZFEjr1zCsicMWpAA",
  "token_type":"example",
  "expires_in":3600,
  "refresh_token":"tGzv3JOkF0XG5Qx2TlKWIA",
  "example_parameter":"example_value"
}
```

#### Obtaining Authorization

[Obtaining Authorization](https://tools.ietf.org/html/rfc6749#section-4)

* To request an access token, the client obtains authorization from the resource owner.

* The authorization is expressed in the form of an authorization grant, which the client uses to request the access token.

[Read RFC 6749 Sections 4.1.1 4.4.3 for more information on Authorization Code Grant, Implicit Grant, Resource Owner Password Credentials Grant, and the Client Credentials Grant](https://tools.ietf.org/html/rfc6749#section-4.1.1)

#### Issuing an Access Token

[Issuing an Access Token](https://tools.ietf.org/html/rfc6749#section-5)

*If the access token request is valid and authorized, the authorization server issues an access token and optional refresh token.*

*If the request failed client authentication or is invalid, the authorization server returns an error response.

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [Authorization](./authorization.md) | [Federation](./federation.md) →
