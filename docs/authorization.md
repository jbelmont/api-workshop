API Workshop - Authorization

## Sections:

* [Basic overview of Authorization](#basic-overview-of-authorization)
* [Bread Crumb Navigation](#bread-crumb-navigation)

#### Basic overview of Authorization

[Wikipedia](https://en.wikipedia.org/wiki/Authorization)

> Access control in computer systems and networks rely on access policies. The access control process can be divided into the following phases: policy definition phase where access is authorized, and policy enforcement phase where access requests are approved or disapproved. Authorization is the function of the policy definition phase which precedes the policy enforcement phase where access requests are approved or disapproved based on the previously defined authorizations.

> Most modern, multi-user operating systems include access control and thereby rely on authorization. Access control also uses authentication to verify the identity of consumers. When a consumer tries to access a resource, the access control process checks that the consumer has been authorized to use that resource. Authorization is the responsibility of an authority, such as a department manager, within the application domain, but is often delegated to a custodian such as a system administrator. Authorizations are expressed as access policies in some types of "policy definition application", e.g. in the form of an access control list or a capability, on the basis of the "principle of least privilege": consumers should only be authorized to access whatever they need to do their jobs. Older and single user operating systems often had weak or non-existent authentication and access control systems.

> "Anonymous consumers" or "guests", are consumers that have not been required to authenticate. They often have limited authorization. On a distributed system, it is often desirable to grant access without requiring a unique identity. Familiar examples of access tokens include keys and tickets: they grant access without proving identity.

> Trusted consumers are often authorized for unrestricted access to resources on a system, but must be authenticated so that the access control system can make the access approval decision. "Partially trusted" and guests will often have restricted authorization in order to protect resources against improper access and usage. The access policy in some operating systems, by default, grant all consumers full access to all resources. Others do the opposite, insisting that the administrator explicitly authorizes a consumer to use each resource.

> Even when access is controlled through a combination of authentication and access control lists, the problems of maintaining the authorization data is not trivial, and often represents as much administrative burden as managing authentication credentials. It is often necessary to change or remove a user's authorization: this is done by changing or deleting the corresponding access rules on the system. Using atomic authorization is an alternative to per-system authorization management, where a trusted third party securely distributes authorization information.

Good information on [RFC 2196 - Site Security Handbook](https://tools.ietf.org/html/rfc2196)

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [OpenID Connect](./openid-connect.md) | [OAuth](./oauth.md) →
