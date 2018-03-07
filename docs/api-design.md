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

#### Bread Crumb Navigation
_________________________

Previous | Next
:------- | ---:
← [Query String](./query-string.md) | [API Testing](./api-testing.md) →
