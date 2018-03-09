API Workshop - API Design

## Sections:

* [Bread Crumb Navigation](#bread-crumb-navigation)

#### Restful URL

Use nouns not verbs for your resources

Let us make student resources

* `GET /students` - Retrieves a list of students

* `GET /students/:id` - Retrieves a specific student

* `POST /students` - Creates a new student

* `PUT /students/:id` - Updates student #5

* `PATCH /students/:id` - Partially updates student #5

* `DELETE /students/:id` - Deletes student #5

The `/students` endpoint is for a collection

The `/students/:id` endpoint is for a specific element in the collection like for example `/students/5`

Doing this will help you get HTTP verbs/actions out of your base URLs.

##### Why are verbs bad in URLs?

Let us try to think about all the particular types of actions we can do with students:

* `/getAllStudents`
* `/checkStudentGrades`
* `/retrieveStudent`
* `/updateStudents`
* `/updateSpecificStudent`
* `/deleteAllStudents`
* `/deleteSpecificStudent`

Mapping urls to actions can quickly get out of hand and is not intuitive for app developers.

**Keep verbs out of your base URLs**

**Use HTTP verbs to operate on the collections and elements.**

##### CRUD

For our student resources, we have two base URLs that use nouns as labels, and we can operate on them with HTTP verbs. 

Our HTTP verbs are `POST`, `GET`, `PUT`, `PATCH`, and `DELETE`. 

We can map them to the acronym CRUD: `CREATE READ UPDATE DELETE`

##### Resource Table

| Resource | POST | GET | PUT | PATCH | DELETE |
| --- | --- | --- | --- | --- | --- |
| `/students` | Create a New Student | List all students | Bulk Update of all Students | Partial Update | Delete all Students |
| `/students/:id` | Error | Show Rocky Balboa | If Rocky exists then update him else Error | Delete Rocky Balboa |

As you can see With the two resources (`/students` and `/students/:id`) and the 5 HTTP verbs, we have a rich set of capabilities that is intuitive to use.

**Tip:**

* Use two base URLs per resource.

* Keep verbs out of your base URLs.  

* Use HTTP verbs to operate on the collections and elements.

#### Plural and Concrete Names

Try to favor plural names for a resource but above all else be consistent, don't use a plural name for one resource and then a singular name for another, this can quickly get confusing for an API consumer.

For example we used `/students` over `/student` for our student resources

Choose concrete names over abstract names

For example what does an `/items` resource convey to you, not much right

But what does `/toys` resource convey to you

The level of abstraction can depend on scenarios but concrete names are more intuitive for API consumers.

#### Resource Associations

Resources tend to have relationships with other resources.

Let us go back to our students resource.

Now let us say that each student will have an associated teachers

If we wanted to get all the students belonging to a specific teacher we could do the following:

`GET /teachers/:id/students` so for example an http request like this `GET /teachers/754/students` would get all students belonging to teacher number 754

If we wanted to add a new student to a particular teacher we could do the following:

`POST /teachers/:id/students` so for example an http request like this `POST /teachers/754/student` would create a new student under teacher number 754

Relationships can get very complex however and it is not uncommon to see APIs nest 5 to 6 levels deep.

Once you have the primary key of a relationship then you shouldn't need to include relationship levels above the specific noun/object.

Meaning there should not be many cases where you have to make particular URL  more than 3 levels deep `/resource/identifier/resource`.

#### Limit complexity by using query strings

Many APIs have many different states that a resource can be in.

Simplify complexities by utilizing query string parameters.

*Any optional state that a particular resource can be in, can be conveyed by a particular query string.*

To get all students living in dorms perhaps you could use the following query:

`GET /students?dorms=true`

Keep the API intuitive by using query strings where appropriate.

#### Error Handling

To most API consumers your API is a blackbox so having good error reporting in place provides proper context and visibility to an API.

The Extreme Programming Model teaches developers to learn to write through errors.

So for example if an API consumer uses a TDD approach when building their application then having good error reporting will help the app developers as they build their application.

App Developers also depend on well-designed errors when they are troubleshooting and when resolving issues as they are building their application with your API.

##### Twilio Error Reporting

Let us look at a twilio API error:

We will make the following restful call:

```curl
curl -X POST https://notify.twilio.com/v1/Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Notifications \
    -d 'Identity=0000005' \
    -d 'Title=Generic loooooooong title for all Bindings' \
    -d 'Body=This is the body for all Bindings' \
    -d 'Data={"custom_key1":"custom value 1","custom_key2":"custom value 2"}' \
    -d 'Fcm={"notification":{"title":"New alert","body":"Hello Bob!"}}' \
    -d 'Apn={"aps":{"alert":{"title":"New alert","body":"Hello Bob!"}}}' \
    -u 'MyAuthToken12345:your_auth_token'
```

Here we put an improper identity for the notifications API.

We get the following JSON response:

```json
{
	"code": 20404,
	"message": "The requested resource /Services/ISXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Notifications was not found",
	"more_info": "https://www.twilio.com/docs/errors/20404",
	"status": 404
}
```

Notice here that you get an internal error code for twilio that maps to an application specific error code.

You get a detailed message explaining what occurred and you get a link to see how to fix the issue.

This is the kind of error response that is useful to an application developer where you give a descriptive message and provide a link to resolve the issue.

Try to be verbose and use plain language descriptions in your error messages. 

Add as many hints as you can possibly think of to the error itself to help the app developer figure out the problem.

I recommend that you add a link in your description for more information, like Twilio provides. 

##### Using HTTP Status Codes

Use HTTP status codes and try to map status codes cleanly to relevant standard-based codes.

There are over 70 HTTP status codes and not every developer has all of the status codes memorized.

Try to pick a decent subset of status code for your API.

Keep in mind that Google GData API only uses 10 status codes. 

Netflix uses about 9 status codes.

* 200, 201, 304, 400, 401, 403, 404, 409, 410, 500

Essentially you can reduce an interaction between an app and an API to 3 outcomes:

* Everything worked (success) 2xx-level http response code
* The application did something incorrect (client error) 4xx-level http response code
* The API did something incorrect (server error) 5xx-level http response code

Obviously this is completely true as there is need for 3xx-level (redirects) at times.

The main point I am trying to make is that start with a basic subset for your API and add more as you need them.

#### API Versioning

*Try not to release an API without a version and make the version mandatory.*

API Providers use different techniques to version their APIs.

A good approach to use is to specify the version with a 'v' prefix. 

Move the verions or `v` all the way to the left in the URL so that it has the highest scope (e.g. /v1/students).

Use a simple ordinal number like `v1` and avoid using the dot notation like v1.2 because it implies a level of granularity for versioning that doesn't work well with APIs. 

**Stick with v1, v2, v3, and so on**

##### API versions to maintain

How many API versions should you try to maintain? Try to maintain at least one version behind.

You want to give developers enough time to update their application code, but it all depends on the application developers' platform.

Some might need 6 months while others might need a year.

##### Puting API version information in the header

Although putting api version information in the HTTP header is more aligned with Roy Fielding's vision of rest not a lot API Providers do this.

Here is an example where using HTTP headers is cleaner:

These for example, all represent the same resource:

```http
GET students/1 
Content-Type: application/json
```

```http
GET students/1 
Content-Type: application/xml
```

In this example the header appears to be more correct and is still a very strong API design.

You will see some apis do the following though `GET students/1.json` and `GET students/1.xml` and that is fine as well.

#### Partial Response

We discussed using the convention of a fields query string in [Limit Payload section of the Query String Section](https://github.com/jbelmont/api-workshop/blob/736b679630a2e4d9dfa4b8445073ff9af3a125ea/docs/query-string.md#limit-payload-information-that-resources-return)

The field query parameters are also used in google and linkedin apis

It was actually google that pioneered the concept of a partial response, Read more about api performance tips at [Google Developers Performance Site](https://developers.google.com/discovery/v1/performance)

As we stated earlier in the query string section you can use a comma delimited list to the fields parameter.

Here is an example again with the students api:

```http
GET /students?fields=id,name,grade
```

#### Pagination

Stick to using the concept of a limit and offset as they are common with apis and developers are used what they mean

So let us go back to the students API.

```http
GET /students?limit=25&offset=50
```

This will get students 25 to 50

A lot of APIs set a default limit of 10 and an offset of 10

The pagination defaults are dependent on your data size. 

If your resources are large, you probably want to limit it to fewer than 10; if resources are small, then consider a larger limit.

#### API Responses that do not involve resources

There are times that an api response will not involve an actual resource.

For instance a financial institution might have an endpoint that does some type of calculation and which never hits a database.

An action with a name like this indicates that you aren't dealing with a "resource" type of response:

* Calculate
* Translate
* Convert

It is appropriate in this cases to use a verb instead of a noun.

Let us go back to the students api example:

`/calculateFinalGrade?grades=90,83,99,85,91,92&weightedAvg=5`

This endpoint although trivial is doing a calculation so the name `calculateFinalGrade` is appropriate in this instance.

You should make it clear in the API documentation that this is a `non-resource` scenario.

If you can separate endpoints like this into a separate part of your API so they get clearly delineated from resource type endpoints.

**Remember resources deal with nouns, while actions deal with verbs**

#### Accepting different types of responses

It is helpful to application developers can use multiple types of formats when working with a particular format.

**On the other hand to make things uniform only return one format.**

[Google Blogger API](https://developers.google.com/blogger/docs/3.0/getting_started)

Notice here that the api uses a header of `Content-Type: application/json`

Some other apis might do something like this: `GET /events.json` right in the url

We could do something similar in the students api like this:

`GET /students/12345/summary.json`

Most web apis will have a default format of JSON and it is okay to assume a default format as long as you highlight this in the API documentation.

#### Naming Attributes

Let us say you are returning metadata information back from a particular resource.

For instance if were to create new student in the students api:

`POST /students` we could return attributes like this `createdAt`, and `updatedAt`.

Some APIs will name an attribute like this `created_at` and `updated_at` but try to avoid this.

An attribute name like `created_at` does not follow JavaScript notation.

For instance once you parse the response:

```
let student = JSON.parse(response);

let createdAt = student.created_at;
```

Notice here that the `created_at` attribute breaks JavaScript convention.

Here are some recommendations:

* Pick JSON as default type

* Follow JavaScript conventions for naming attributes
  * camelCase (studentName) is best but depending on the attribute you could do Pascal Case (FinalGrade)
  * Helps JavaScript developers work with names that go with typical JavaScript/JSON notation
  
#### Searching

Let us like how google does searching:

I searched api workshops in Google with the following query: `api workshop github`

The query that google sent out looks like this:

`search?q=api+workshop+github&oq=api+workshop+github` I removed some specific query parameters that google added

Notice here Google uses `search` and uses the conventions of `q=api+workshop+github` to do the search.

*Remember that spaces are not allowed in query string so Google used the `+` instead*

If we wanted to add scope to our search we could do the following in the students api

`GET /teachers/1234/students?q=math+science` Here we do a search for all students that have math and science with teacher `1234`

If we wanted to format the result we could do the following `GET /search.xml?q=carolina+students` which would return in xml format.

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [Query String](./query-string.md) | [API Testing](./api-testing.md) →
