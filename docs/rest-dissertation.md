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

content

#### CHAPTER 4 - Designing the Web Architecture: Problems and Insights

content

#### CHAPTER 5 - Representational State Transfer (REST)

content

#### CHAPTER 6 - Experience and Evaluation

content


#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [Restful API Introduction](./restful-intro.md) | [Rest Architecture Constraints](./rest-constraints.md) →
