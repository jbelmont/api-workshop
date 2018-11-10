API Workshop - API Security
## Sections:

* [The Pillars of API Security](#the-pillars-of-api-security)
* [Bread Crumb Navigation](#bread-crumb-navigation)

#### The Pillars of API Security

There are 4 pillars of API Security that we will go over in this section:

* Authentication
* Authorization
* Federation
* Delegation

###### Authentication

According to [Wikipedia](https://en.wikipedia.org/wiki/Authentication):

> Authentication is the act of confirming the truth of an attribute of a single piece of data claimed true by an entity. In contrast with identification, which refers to the act of stating or otherwise indicating a claim purportedly attesting to a person or thing's identity, authentication is the process of actually confirming that identity. It might involve confirming the identity of a person by validating their identity documents, verifying the authenticity of a website with a digital certificate, determining the age of an artifact by carbon dating, or ensuring that a product is what its packaging and labeling claim to be. In other words, authentication often involves verifying the validity of at least one form of identification.

Authentication is the process of verifying who you are.
When you log on to a PC with a user name and password you are authenticating.

**Authentication is about who somebody is.**

###### Authorization

According to [Wikipedia](https://en.wikipedia.org/wiki/Authorization):

> Authorization is the function of specifying access rights/privileges to resources related to information security and computer security in general and to access control in particular. More formally, "to authorize" is to define an access policy. For example, human resources staff are normally authorized to access employee records and this policy is usually formalized as access control rules in a computer system. During operation, the system uses the access control rules to decide whether access requests from (authentication) shall be approved (granted) or disapproved (rejected) [2]. Resources include individual files or an item's data, computer programs, computer devices and functionality provided by computer applications. Examples of consumers are computer users, computer Software and other Hardware on the computer.

Authorization is the process of verifying that you have access to something.
Gaining access to a resource because the permissions configured on it allow you access is authorization.

**Authorisation is about what someone is allowed to do.**

###### Federation

According to [Wikipedia](https://en.wikipedia.org/wiki/Federated_identity)

> A federated identity in information technology is the means of linking a person's electronic identity and attributes, stored across multiple distinct identity management systems.

> Federated identity is related to single sign-on (SSO), in which a user's single authentication ticket, or token, is trusted across multiple IT systems or even organizations. SSO is a subset of federated identity management, as it relates only to authentication and is understood on the level of technical interoperability and it would not be possible without some sort of federation.

###### Delegation

Delegation is another process by which access and rights can be given to authorized users while maintaining a relatively limited amount of access. Whereas federation can work by giving a user a token to use on multiple domains, delegation works by authorizing a user to function partially as if they were another user.

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [API Testing Tools](./api-testing-tools.md) | [Authentication](./authentication.md) →
