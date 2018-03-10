API Workshop - Rest Constraints

## Sections:

* [Bread Crumb Navigation](#bread-crumb-navigation)

#### Rest Constraints overview

There are 5 Constraints that the Rest style advocates:

* [Client-Server](https://en.wikipedia.org/wiki/Client%E2%80%93server_model)
* [Stateless](https://en.wikipedia.org/wiki/State_%28computer_science%29)
* [Cache](https://en.wikipedia.org/wiki/Cache_%28computing%29)
* [Uniform Interface](https://en.wikipedia.org/wiki/Representational_state_transfer#Uniform_interface)
* [Layered System](https://en.wikipedia.org/wiki/Layered_system)
* [Code-On-Demand](https://en.wikipedia.org/wiki/Code_on_demand)

#### Client-Server

The server component provides a function or service to one or more clients.

The client can initiate requests for such services.

Servers are classified by the services they provide.

For example, a web server serves web pages and a file server serves files.

Separation of concerns is  the main principle behind the client-server model.

By separating the UI (the user interfaces) from the data storage we in turn gain portability of the UI across different platforms and improve scalability in the server components.

This separation of concerns helps the components evolve independently and thus helps support internet scale issues.

#### Stateless

Communication must be stateless in nature.

* Each  request  from the client to the server  must  contain  all  of  the  information  necessary  to understand the  request

* The request cannot  take  advantage  of  any  stored  context  on  the  server.

* Session state is therefore kept entirely on the client.

The stateless constraint emphasizes visibility, reliability, and scalability.

Visibility is improved because any monitoring sytem will not have to look beyond a single requet to determine the nature of the request.

"Roy Fielding":

> Reliability is improved because it eases the task of recovering from partial failures. Scalability is improved because not  having  to  store  state  between  requests  allows  the  server  component  to  quickly  free resources,  and  further  simplifies  implementation  because  the  server  doesn’t  have  to manage resource usage across requests.

"Roy Fielding":

> Like most architectural choices, the stateless constraint reflects a design trade-off. The disadvantage is that it may decrease network performance by increasing the repetitive data(per-interaction overhead) sent in a series of requests, since that data cannot be left on the server  in  a  shared  context.  In  addition,  placing  the  application  state  on  the  client-side reduces  the  server’s  control  over  consistent  application  behavior,  since  the  application becomes  dependent  on  the  correct  implementation  of  semantics  across  multiple  clientversions.

#### Cache

Network Efficiency is improved by adding cache constraints.

"Roy Fielding":

> Cache constraints require that the data within  a  response  to  a  request  be  implicitly  or  explicitly  labeled  as  cacheable  or  non-cacheable.  If  a  response  is  cacheable,  then  a  client  cache  is  given  the  right  to  reuse  that response data for later, equivalent requests.

"Roy Fielding":

> The advantage of adding cache constraints is that they have the potential to partially orcompletely  eliminate  some  interactions,  improving  efficiency,  scalability,  and  user-perceived  performance  by  reducing  the  average  latency  of  a  series  of  interactions.  The trade-off,  however,  is  that  a  cache  can  decrease  reliability  if  stale  data  within  the  cache differs significantly from the data that would have been obtained had the request been sent directly to the server.

"Roy Fielding":

> The  early  Web  architecture was defined by the client-cache-stateless-server set of constraints. That is, the design rationale presented  for  the  Web  architecture  prior  to  1994  focused  on  stateless  client-server interaction  for  the  exchange  of  static  documents  over  the  Internet.  The  protocols  for communicating  interactions  had  rudimentary  support  for  non-shared  caches,  but  did  not constrain the interface to a consistent set of semantics for all resources. Instead, the Web relied  on  the  use  of  a  common  client-server  implementation  library  (CERN  libwww)  to maintain consistency across Web applications. Developers  of  Web  implementations  had  already  exceeded  the  early  design.  In addition to static documents, requests could identify services that dynamically generated responses,  such  as  image-maps and  server-side  scripts.

#### Uniform Interface

"Roy Fielding":

> The  central  feature  that  distinguishes  the  REST  architectural  style  from  other  network-based styles is its emphasis on a uniform interface between components. By applying the software engineering principle of generality to the component interface, the overall  system  architecture  is  simplified  and  the  visibility  of  interactions  is  improved. Implementations  are  decoupled  from  the  services  they  provide,  which  encourages independent  evolvability.

The four constraints for a uniform interface are:

* Identification of Resources
  * Individual resources are identified in requests, for example using URIs in Web-based REST systems.

* Manipulation  of  resources  through  representations
  * When a client holds a representation of a resource, including any metadata attached, it has enough information to modify or delete the resource.

* Self-descriptive messages
  * Each message includes enough information to describe how to process the message.

* Hypermedia as the engine of application state (HATEOAS)
  * A REST client should then be able to use server-provided links dynamically to discover all the available actions and resources it needs.


#### Layered System

"Roy Fielding":

> A layered  system style allows  an  architecture  to  be  composed  of  hierarchical  layers  by  constraining  component behavior such that each component cannot "see" beyond the immediate layer with which they are interacting. By restricting knowledge of the system to a single layer, we place abound on the overall system complexity and promote substrate independence.

A client cannot ordinarily tell whether it is connected directly to the end server, or to an intermediary along the way.

Intermediary servers may improve system scalability by enabling load balancing and by providing shared caches.

Layered System can also enforce security policies.

#### Code-On-Demand

"Roy Fielding":

> REST allows client functionality to be extended by downloading and executing code in the form of applets or scripts. This simplifies clients by reducing the number of features required to be pre-implemented. Allowing features to be downloaded after deployment improves system extensibility. However, it also reduces visibility, and thus is only an optional constraint within REST.

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [Roy Fieldings Dissertation on REST](./rest-dissertation.md) | [Richardson Maturity Model](./maturity-model.md) →
