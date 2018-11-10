API Workshop - Delegation

## Sections:

* [Types of Delegation](#types-of-delegation)
* [Bread Crumb Navigation](#bread-crumb-navigation)

#### Types of Delegation

[Wikipedia](https://en.wikipedia.org/wiki/Delegation_(computer_security))

2 classes of delegation:

* Delegation at Authentication/Identity Level
* Delegation at Authorization/Access Control Level

###### Delegation at Authentication/Identity Level

> It is defined as follows: If an authentication mechanism provides an effective identity different from the validated identity of the user then it is called identity delegation at the authentication level, provided the owner of the effective identity has previously authorized the owner of the validated identity to use his identity.

> The existing techniques of identity delegation using sudo or su commands of UNIX are very popular.[citation needed] To use the sudo command, a person first has to start his session with his own original identity. It requires the delegated account password or explicit authorizations granted by the system administrator. The user login delegation described in the patent of Mercredi and Frey is also an identity delegation.

###### Delegation at Authorization/Access Control Level

> The most common way of ensuring computer security is access control mechanisms provided by operating systems such as UNIX, Linux, Windows, Mac OS, etc.

> If the delegation is for very specific rights, also known as fine-grained, such as with Role-based access control (RBAC) delegation, then there is always a risk of under-delegation, i.e., the delegator does not delegate all the necessary permissions to perform a delegated job. This may cause the denial of service, which is very undesirable in some environments, such as in safety critical systems or in health care. In RBAC-based delegation, one option to achieve delegation is by reassigning a set of permissions to the role of a delegatee; however, finding the relevant permissions for a particular job is not an easy task for large and complex systems. Moreover, by assigning these permissions to a delegatee role, all other users who are associated with that particular role get the delegated rights.

> If the delegation is achieved by assigning the roles of a delegator to a delegatee then it would not only be a case of over-delegation but also the problem that the delegator has to figure out what roles, in the complex hierarchy of RBAC, are necessary to perform a particular job. These types of problems are not present in identity delegation mechanisms and normally the user interface is simpler.

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [Federation](./federation.md) | [API Description Languages](./api-description-languages.md) →
