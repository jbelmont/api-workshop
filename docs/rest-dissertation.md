API Workshop - Rest Dissertation

## Sections:

* [Bread Crumb Navigation](#bread-crumb-navigation)

#### Roy Fielding Abstract Quotes

"Roy Fielding":

> My  work  is motivated  by  the  desire  to  understand  and  evaluate  the  architectural  design  of  network-based  application  software  through  principled  use  of  architectural  constraints,  thereby obtaining the functional, performance, and social properties desired of an architecture. An architectural style is a named, coordinated set of architectural constraints.

"Roy Fielding":

> This  dissertation  defines  a  framework  for  understanding  software  architecture  viaarchitectural  styles  and  demonstrates  how  styles  can  be  used  to  guide  the  architecturaldesign  of  network-based  application  software.

"Roy Fielding":

> A  good  architecture is  not  created  in  a  vacuum. All design decisions at  the architectural  level should  be made  within the context of  the  functional,  behavioral,  and social requirements of the system being designed, which is a principle that applies equally to  both  software  architecture and the  traditional  field  of  building  architecture.

"Roy Fielding":

> The hyperbole of The Architects Sketch may seem ridiculous, but consider how oftenwe see software projects begin with adoption of the latest fad in architectural design, and only  later  discover  whether  or  not  the  system  requirements  call  for  such  an  architecture.

"Roy Fielding":

> This  dissertation  explores  a  junction  on  the  frontiers  of  two  research  disciplines  in computer  science:  software  and  networking. 

"Roy Fielding":

> The Web is intended to be an Internet-scale distributed hypermedia system, which means considerably more than just geographical dispersion. The Internet is about  interconnecting  information  networks  across  organizational  boundaries.  Suppliers of information services must be able to cope with the demands of anarchic scalability and the independent deployment of software components. Distributed hypermedia provides a uniform means of accessing services through the embedding of action controls within the presentation of information retrieved from remote sites. An architecture for the Web must therefore  be  designed  with  the  context  of  communicating  large-grain  data  objects  across high-latency networks and multiple trust boundaries.

"Roy Fielding":

> This work  was  done  in  conjunction  with  my  authoring  of  the  Internet  standards  for  the Hypertext  Transfer  Protocol  (HTTP)  and  Uniform  Resource  Identifiers  (URI),  the  two specifications that define the generic interface used by all component interactions on the Web.

#### Summary of Dissertation via Roy Fielding

**In  summary,  this  dissertation  makes  the  following  contributions  to  software  research within the field of Information and Computer Science:**

* a framework for understanding software architecture through architectural styles, including a consistent set of terminology for describing software architecture

* a classification of architectural styles for network-based application software by the architectural properties they would induce when applied to the architecture for a distributed hypermedia system

* REST, a novel architectural style for distributed hypermedia systems

* application and evaluation of the REST architectural style in the design and deployment of the architecture for the modern World Wide Web.

#### Chapter 1 - Software Architecture

"Roy Fielding":

> A software architecture is an abstraction of the run-time elements of a software system during some phase of its operation. A system may be composed of many levels of abstraction and many phases of operation, each with its own software architecture.

This chapter goes into detail on what Software Architecture is. Discussion about software architecture patterns ensues.

Fielding goes into detail about the Gang of Four Design Patterns and expounds upon some of the choices that Object-Oriented Design patterns want to tackle.

"Roy Fielding":

> In  many  ways,  Alexander’s  patterns  have  more  in  common  with  software  architectural styles than the design patterns of OOPL research. An architectural style, as a coordinated set  of  constraints,  is  applied  to  a  design  space  in  order  to  induce  the  architectural properties that are desired of the system. By applying a style, an architect is differentiatingthe software design space in the hope that the result will better match the forces inherent in the  application,  thus  leading  to  system  behavior  that  enhances  the  natural  pattern  rather than conflicting with it.

Here Roy is talking about [Christopher Alexander](https://en.wikipedia.org/wiki/Christopher_Alexander) who is a well known building architect. 

This chapter provides terminology that set the stage for later chapters.

#### CHAPTER 2 - Network-based Application Architectures

This chapter went over the framework of network-based application architectures.

Fielding explains network latency concepts and how the style of a network application architecture influences the performance of such architectures

There is great discussions on the following topics:

* Scalability
* Simplicity
* Modifiability
* Visibility
* Portability
* Reliability

**The overall premise of this chapter seems to point that architectural style influences architectural design**

#### CHAPTER 3 - Network-based Architectural Styles

This chapter discusses many different types of architectural styles in network-based architectures.

"Roy Fielding":

> I  have  intentionally  excluded  styles  that  do  not  enhance  the  communication  or interaction properties when combined with one of the surveyed styles to form a network-based  application.  

Pipe and Filter styles is mentioned quite often in this chapter.

A Pipe and Filter is a simple, yet powerful architecture. A pipe and filter consists of any number of components (filters) that transform or filter data, before passing it on via connectors (pipes) to other components. The filters are all working at the same time. The architecture is often used as a simple sequence, but it may also be used for very complex structures.

The most well known pipe and filter system is a Unix Program or in compilers.

Fielding discusses many different Hierarchical Network Styles such as Client Server, Layered Client Server and more.

#### CHAPTER 4 - Designing the Web Architecture: Problems and Insights

This chapter describes in detail the early thoughts on how the World Wide Web (WWW) would be.
As its intent was to share documents with University Faculty and Government Employees in Physics Laboratories.

Fielding explains how no one predicted the explosive growth of the web and the impending interest with businesses.

The Internet Engineering Task Force (IETF) set out to form standardization committees for URI, HTTP, and HTML.

Fielding formulates 3 hypothesis:

* The design rationale behind the WWW architecture can be described by an architectural style consisting of the set of constraints applied to the elements within the Web architecture.

* Constraints can be added to the WWW architectural style to derive a new hybrid style that better reflects the desired properties of a modern Web architecture.

* Proposals to modify the Web architecture can be compared to the updated WWW architectural style and analyzed for conflicts prior to deployment.

Fielding explains how over the past 6 years as a PHD student he worked with many individiduals and committees to help formulate standards.

Roy Fielding helped write much of HTTP 1.1 Specification

[HTTP 1.1 Specification](https://www.w3.org/Protocols/HTTP/1.1/draft-ietf-http-v11-spec-01)

#### CHAPTER 5 - Representational State Transfer (REST)

Fielding states the constraints of REST in this Chapter:

* Client - Server
  * The first constraints added to our hybrid style are those of the client-server architectural style.
  * Separation  of  concerns  is  the  principlebehind  the  client-server  constraints.  
  * By  separating  the  user  interface  concerns  from  the data  storage  concerns,  we  improve  the  portability  of  the  user  interface  across  multiple platforms  and  improve  scalability  by  simplifying  the  server  components.
  
* Stateless
  * Communication must be stateless in  nature,  as  in  the  client-stateless-server  (CSS)  style, such that  each  request  from  client  to  server  must  contain  all  of  the  information  necessary  to understand  the  request,  and  cannot  take  advantage  of  any  stored  context  on  the  server.
  * Session state is therefore kept entirely on the client.
  
 * Cache 
  * In order to improve network efficiency, we add cache constraints to form the client-cache-stateless-server style.
  * Cache constraints require that the data within  a  response  to  a  request  be  implicitly  or  explicitly  labeled  as  cacheable  or  non-cacheable.  
  * If  a  response  is  cacheable,  then  a  client  cache  is  given  the  right  to  reuse  that response data for later, equivalent requests.
  
*  Uniform Interface
  * The  central  feature  that  distinguishes  the  REST  architectural  style  from  other  network-based styles is its emphasis on a uniform interface between components.
  * By applying the software engineering principle of generality to the component interface, the overall  system  architecture  is  simplified  and  the  visibility  of  interactions  is  improved.
  
* Layered System
  * The  layered  system  style allows  an  architecture  to  be  composed  of  hierarchical  layers  by  constraining  component behavior such that each component cannot "see" beyond the immediate layer with which they are interacting. 
  * By restricting knowledge of the system to a single layer, we place a bound on the overall system complexity and promote substrate independence.
  
* Code-On-Demand
  * REST   allows   client   functionality   to   be   extended   by downloading and executing code in the form of applets or scripts. 
  * This simplifies clientsby reducing the number of features required to be pre-implemented. 
  * Allowing features to be downloaded after deployment improves system extensibility. 
  * However, it also reduces visibility, and thus is only an optional constraint within REST

"Roy Fielding":

> All  REST  interactions  are  stateless.  That  is,  each  request  contains  all  of  theinformation  necessary  for  a  connector  to  understand  the  request,  independent  of  anyrequests  that  may  have  preceded  it.  This  restriction  accomplishes  four  functions:  1)  it removes  any  need  for  the  connectors  to  retain  application  state  between  requests,  thus reducing  consumption  of  physical  resources  and  improving  scalability;  2)  it  allows interactions  to  be  processed  in  parallel  without  requiring  that  the  processing  mechanism understand the interaction semantics; 3) it allows an intermediary to view and understand a request in isolation, which may be necessary when services are dynamically rearranged; and,  4)  it  forces  all  of  the  information  that  might  factor  into  the  reusability  of  a  cachedresponse to be present in each request.

"Roy Fielding":

> REST  provides  a  set  of  architectural  constraints  that, when applied as a whole, emphasizes scalability of component interactions, generality of interfaces,  independent  deployment  of  components,  and  intermediary  components  to reduce  interaction  latency,  enforce  security,  and  encapsulate  legacy  systems.  I  described the software engineering principles guiding REST and the interaction constraints chosen to retain those principles, while contrasting them to the constraints of other architectural styles

#### CHAPTER 6 - Experience and Evaluation

Fielding spends a lot of the chapter discussing different experiences and impacts of REST.

"Fielding":

> The early Web architecture defined URI as document identifiers. Authors were instructed to  define  identifiers  in  terms  of  a  document’s  location  on  the  network.  Web  protocols could  then  be  used  to  retrieve  that  document.

"Fielding":

> Information hiding is one of the key software engineering principles that motivates the uniform  interface  of  REST.  Because  a  client  is  restricted  to  the  manipulation  of representations  rather  than  directly  accessing  the  implementation  of  a  resource,  the implementation  can  be  constructed  in  whatever  form  is  desired  by  the  naming  authority without  impacting  the  clients  that  may  use  its  representations.

"Fielding":

> The key problem areas in HTTP that were identified by REST included planning for the  deployment  of  new  protocol  versions,  separating  message  parsing  from  HTTP semantics and the underlying transport layer (TCP), distinguishing between authoritative and  non-authoritative  responses,  fine-grained  control  of  caching,  and  various  aspects  of the  protocol  that  failed  to  be  self-descriptive.

"Fielding":

> HTTP does not support write-back caching. An HTTP cache cannot assume that what gets written through it is the same as what would be retrievable from a subsequent request for that  resource,  and  thus  it  cannot  cache  a  PUT  request  body  and  reuse  it  for  a  later  GET response. There are two reasons for this rule: 1) metadata might be generated behind-the-scenes, and 2) access control on later GET requests cannot be determined from the PUT request. However, since write actions using the Web are extremely rare, the lack of write-back caching does not have a significant impact on performance.

Fielding also lays some disadvantages of HTTP cookie usage and frames (i.e. iframes) in terms of security concerns and how they violate REST constraints.

Fielding also explains how Multipurpose Internet Mail Extensions [MIME](https://en.wikipedia.org/wiki/MIME) was adopted by HTTP but how the message semantics don't line up well with the intentions of the WEB.

Fielding also explains the success of the Apache Project and how Apache became the first major implementor of the HTTP 1.1 Specification. Also notes its success as an open source project.

"Fielding":

> A network-based API is an on-the-wire syntax, with defined semantics, for application interactions. A network-based API does not place any restrictions on the application code aside from the need to read/write to the network, but does place restrictions on the set of semantics  that  can  be  effectively  communicated  across  the  interface.  On  the  plus  side, performance   is   only   bounded   by   the   protocol   design   and   not   by   any   particular implementation of that design.

Fielding also makes some distinctions with HTTP with CORBA.

Fielding also differentiates RPC from HTTP.

Lastly Fielding explains why JavaScript is better suited for the web than Java.

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [Restful API Introduction](./restful-intro.md) | [Rest Architecture Constraints](./rest-constraints.md) →
